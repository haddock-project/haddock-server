package ws

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/gofiber/websocket/v2"
)

type Client struct {
	Id    uint
	Type  uint
	Token string
	Conn  *websocket.Conn
}

var Websocket struct {
	Clients      []*Client
	AvailableIds []uint
}

// Register a websocket connection to the cache
func Register(conn *websocket.Conn) (Client, error) {

	/*
		Generate an ID
	*/
	var clientId uint
	if len(Websocket.AvailableIds) == 0 {
		//add an ID
		clientId = uint(len(Websocket.Clients)) + 1
	} else {
		//Get a random available ID and remove it
		clientId, Websocket.AvailableIds = Websocket.AvailableIds[0], Websocket.AvailableIds[1:]
	}

	/*
		Generate a token for the client (to reference the websocket using the rest API)
	*/
	b := make([]byte, 24)
	_, err := rand.Read(b)
	token := hex.EncodeToString(b)

	/*
		Add client to known clients
	*/
	client := Client{Id: clientId, Conn: conn, Token: token}
	Websocket.Clients = append(Websocket.Clients, &client)

	//handle token generation error
	if err != nil {
		client.CloseConnection("failed to generate a token")
	}

	return client, nil
}

// RemoveClientId removes a client from the cache from ID
func RemoveClientId(id uint) {
	//search the client index
	wsIndex := 0
	for index, c := range Websocket.Clients {
		//check if the ID match
		if c.Id == id {
			wsIndex = index
			break
		}
	}
	//if there is a match
	if wsIndex != 0 {
		//remove the client from the cache using the index
		Websocket.Clients = append(Websocket.Clients[:wsIndex], Websocket.Clients[wsIndex+1:]...)
	}

	// add ID to available IDS
	Websocket.AvailableIds = append(Websocket.AvailableIds, id)
}

func GetClientByToken(token string) *Client {
	//search the client index
	for _, c := range Websocket.Clients {
		//check if the ID match
		if c.Token == token {
			return c
		}
	}
	return nil
}

func (client *Client) CloseConnection(reason string) {
	/*
		Generate the request's JSON
	*/
	reason = fmt.Sprintf(`{"name": "CLOSE", "args": {"reason":"%s"}}`, reason)

	/*
		Send the event
	*/
	client.Conn.WriteMessage(websocket.TextMessage, []byte(reason))

	/*
		Close the connection
	*/
	client.Conn.Close()

	/*
		Delete client from the cache
	*/
	client.Remove()
}

func (client *Client) Remove() {
	RemoveClientId(client.Id)
}
