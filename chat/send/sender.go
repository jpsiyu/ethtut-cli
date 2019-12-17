package send

import (
	"context"
	"encoding/json"
	"github.com/jpsiyu/ethtut-cli/chat/common"
	"log"

	"github.com/ethereum/go-ethereum/whisper/shhclient"
	"github.com/ethereum/go-ethereum/whisper/whisperv6"
)

type Sender struct {
	client *shhclient.Client
	keyID  string
	user   *common.User
}

func NewSender(user *common.User, client *shhclient.Client, keyID string) *Sender {
	return &Sender{
		client: client,
		keyID:  keyID,
		user:   user,
	}
}

func (sender *Sender) Say(msg string) {
	userMsg := common.UserMsg{
		User: *sender.user,
		Msg:  msg,
	}
	bytes, _ := json.Marshal(&userMsg)

	message := whisperv6.NewMessage{
		SymKeyID:  sender.keyID,
		Payload:   bytes,
		TTL:       60,
		PowTime:   2,
		PowTarget: 2.5,
		Topic:     whisperv6.BytesToTopic(common.Topic()),
	}

	_, err := sender.client.Post(context.Background(), message)
	if err != nil {
		log.Fatal((err))
	}
}
