package receive

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jpsiyu/ethtut-cli/chat/common"

	"github.com/ethereum/go-ethereum/whisper/shhclient"
	"github.com/ethereum/go-ethereum/whisper/whisperv6"
)

type Receiver struct {
	client  *shhclient.Client
	keyID   string
	user    *common.User
	handler func(string)
}

func NewReceiver(user *common.User, client *shhclient.Client, keyID string, handler func(string)) *Receiver {
	return &Receiver{
		client:  client,
		keyID:   keyID,
		user:    user,
		handler: handler,
	}
}

func (receiver *Receiver) Run() {
	messages := make(chan *whisperv6.Message)
	criteria := whisperv6.Criteria{
		PrivateKeyID: receiver.keyID,
	}

	sub, err := receiver.client.SubscribeMessages(context.Background(), criteria, messages)

	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case message := <-messages:
			var userMsg common.UserMsg
			json.Unmarshal(message.Payload, &userMsg)
			formatStr := fmt.Sprintf("(%d)%s: %s\n", userMsg.User.ID, userMsg.User.Name, userMsg.Msg)
			receiver.handler(formatStr)
		}
	}
}
