package database

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
	"time"
)

type User struct {
	Name              string
	Password          string
	NeedPasswordReset bool
	Apps              []string
}

// generateUserPassword generate a random temporary password with specifics criteria (it has to be easy to use)
func generateUserPassword() string {
	// define the random seed
	rand.Seed(time.Now().UnixNano())

	// define random password settings
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	length := 8

	// generate a random password
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

//hashPassword take a password and return a hashed password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//checkPasswordHash returns true if the password matches the hash
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUnregisteredUser(user User) (User, error) {
	//create a new user in the db and create him a password
	user.Password = generateUserPassword()
	user.NeedPasswordReset = true

	//add the user to the db
	return CreateUser(user)
}

func CreateUser(user User) (User, error) {
	var err error

	if !user.NeedPasswordReset {
		//hash the password
		user.Password, err = hashPassword(user.Password)

		if err != nil {
			return user, err
		}
	}

	//add the user to the db
	_, err = db.Exec("INSERT INTO users (user_name, user_password, password_reset) VALUES ($1, $2, $3)", user.Name, user.Password, user.NeedPasswordReset)

	return user, err
}

// AuthUser compare the provided User with db records and return the user if the password is correct
// returns the user, an error and weather the password match or not
// if the user needs to reset the password, raise an error
func AuthUser(user User) (User, error, bool) {
	//get the user from the db
	row := db.QueryRow("SELECT user_password, password_reset FROM users WHERE user_name = $1", user.Name)

	//extract the values from the row
	var password string
	var needPasswordReset bool
	if err := row.Scan(&password, &needPasswordReset); err != nil {
		return user, err, false
	}

	//check if the user is allowed to login
	if needPasswordReset {
		err := errors.New("user needs to reset password")
		return user, err, false
	}

	//compare the password hash with the provided password
	if checkPasswordHash(user.Password, password) {
		return user, nil, false
	}

	// the password hash is matching, return the user
	return user, nil, true
}

//DeleteUser delete a user from the db and all their apps
func DeleteUser(user User) error {
	//delete the user from the db
	_, err := db.Exec("DELETE FROM users WHERE user_name = $1", user.Name)

	return err
}
