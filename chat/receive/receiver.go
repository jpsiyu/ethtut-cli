package receive

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jpsiyu/ethtut-cli/chat/common"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/whisper/shhclient"
	"github.com/ethereum/go-ethereum/whisper/whisperv6"
)

type Receiver struct {
	client *shhclient.Client
	keyID  string
	user   *common.User
}

func NewReceiver(user *common.User, client *shhclient.Client, keyID string) *Receiver {
	return &Receiver{
		client: client,
		keyID:  keyID,
		user:   user,
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
			reader := bufio.NewReader(nil)
			/*
				content, err := reader.ReadString('\n')
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("buffer data", content)
			*/
			reader.Reset(os.Stdin)
			if userMsg.User.ID != receiver.user.ID {
				fmt.Printf("(%d)%s: %s\n", userMsg.User.ID, userMsg.User.Name, userMsg.Msg)
			}
		}
	}
}
