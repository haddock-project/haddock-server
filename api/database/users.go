package database

import (
	"errors"
	"github.com/google/uuid"
	"go/token"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
	"time"
)

var (
	ErrUserNeedPasswordReset = errors.New("user needs to reset password")
	ErrPasswordTooShort      = errors.New("password is too short")
)

type User struct {
	UUID              uuid.UUID       `json:"uuid"`
	Name              string          `json:"name"`
	Icon              string          `json:"profile_picture"`
	Token             token.Token     `json:"token"`
	Permissions       map[string]bool `json:"permissions"`
	Password          string
	NeedPasswordReset bool
}

// GeneratePassword generate a random temporary password with specifics criteria (it has to be easy to use)
// WARNING: override the previous password
// NOTE: the password is not hashed
func (user *User) GeneratePassword() {
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
	user.Password = b.String()
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

// Set update an existing User or create a new one
func (user *User) Set() error {
	var err error

	//check password length
	if len(user.Password) < 8 {
		return ErrPasswordTooShort
	}

	//hash the password
	user.Password, err = hashPassword(user.Password)
	if err != nil {
		return err
	}

	//add the user to the db
	_, err = db.Exec(
		"INSERT INTO users (user_name, user_password, password_reset, user_icon, user_permissions) VALUES ($1, $2, $3, $4, $5)"+
			"ON CONFLICT(user_name) DO UPDATE SET (user_password, password_reset, user_icon, user_permissions) = ($2, $3, $4, $5)",
		user.Name, user.Password, user.NeedPasswordReset, user.Icon, user.Permissions)

	return err
}

// Get fills the user object from its Name
func (user *User) Get() error {
	var perms int
	err := db.QueryRow("SELECT user_password, password_reset, user_icon, user_permissions FROM users   WHERE user_name = ?", user.Name).Scan(&user.Password, &user.NeedPasswordReset, &user.Icon, perms)
	if err != nil {
		return err
	}
	//TODO: bitfields to map
	return nil
}

// Auth compare the provided User with db records and return the user if the password is correct
// returns the user, an error and weather the password match or not
// if the user needs to reset the password, raise an error
func (user User) Auth() (bool, error) {
	//get the user from the db
	row := db.QueryRow("SELECT user_password, password_reset FROM users WHERE user_name = $1", user.Name)

	//extract the values from the row
	var password string
	var needPasswordReset bool
	if err := row.Scan(&password, &needPasswordReset); err != nil {
		return false, err
	}

	//check if the user is allowed to log in
	if needPasswordReset {
		return false, ErrUserNeedPasswordReset
	}

	//compare the password hash with the provided password
	if checkPasswordHash(user.Password, password) {
		return false, nil
	}

	// the password hash is matching, return the user
	return true, nil
}

//Delete a user from the db and all their apps
func (user *User) Delete() error {
	//delete the user from the db
	_, err := db.Exec("DELETE FROM users WHERE user_name = $1", user.Name)

	return err
}
