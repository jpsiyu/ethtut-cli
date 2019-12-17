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
	/*
		keyID, _ := RandomKey()
		fmt.Println(keyID)
		return keyID
	*/
	return "050376991e0d2a08b9e486b5989f032fdda0c92971accb24677dcec20c4f5509"
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
