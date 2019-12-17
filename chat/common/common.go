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

func RandomKeyPair() (string, error) {
	client, err := shhclient.Dial(conf.ShhUrl)
	if err != nil {
		return "", err
	}

	keypair, err := client.NewKeyPair(context.Background())
	if err != nil {
		return "", err
	}
	return keypair, nil
}

func GenSymKey() (string, error) {
	client, err := shhclient.Dial(conf.ShhUrl)
	if err != nil {
		return "", err
	}
	keyID, err := client.GenerateSymmetricKeyFromPassword(context.Background(), "Helo")
	return keyID, err
}

func Topic() []byte {
	return []byte("1234")
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
