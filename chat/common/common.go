package common

import (
	"context"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/whisper/shhclient"
	"github.com/jpsiyu/ethtut-cli/conf"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type UserMsg struct {
	User User   `json:"user"`
	Msg  string `json:"msg"`
}

func GenKey() string {
	return "cc6587d8c5a8925158abffedd30b8df663c1a1eec1b9b4f3157f743c5970e5dd"
}

func RandomKey() (string, error) {
	client, err := shhclient.Dial(conf.ShhUrl)
	if err != nil {
		return "", err
	}

	keyID, err := client.NewKeyPair(context.Background())
	if err != nil {
		return "", err
	}
	return keyID, nil
}

var names []string = []string{
	"Adom",
	"Bob",
	"Cindy",
	"David",
	"Emmy",
	"Fredy",
	"Gay",
	"Hillary",
	"Jack",
	"Kitty",
	"Lemon",
	"Mansole",
	"Net",
}

func RandomUser() User {
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(names))
	name := names[index]
	return User{
		ID:   rand.Int63n(100000000000000),
		Name: name,
	}
}
