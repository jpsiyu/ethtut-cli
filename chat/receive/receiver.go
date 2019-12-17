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
	client   *shhclient.Client
	symKeyID string
	user     *common.User
	handler  func(string)
}

func NewReceiver(user *common.User, client *shhclient.Client, symKeyID string, handler func(string)) *Receiver {
	return &Receiver{
		client:   client,
		symKeyID: symKeyID,
		user:     user,
		handler:  handler,
	}
}

func (receiver *Receiver) Run() {
	messages := make(chan *whisperv6.Message)
	topic := whisperv6.BytesToTopic(common.Topic())
	criteria := whisperv6.Criteria{
		SymKeyID: receiver.symKeyID,
		Topics:   []whisperv6.TopicType{topic},
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
			formatStr := fmt.Sprintf("%s@%d: %s\n", userMsg.User.Name, userMsg.User.ID, userMsg.Msg)
			receiver.handler(formatStr)
		}
	}
}
