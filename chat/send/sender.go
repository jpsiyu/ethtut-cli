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
	client   *shhclient.Client
	symKeyID string
	user     *common.User
	keypair  string
}

func NewSender(user *common.User, client *shhclient.Client, symKeyID string) *Sender {
	keypair, _ := common.RandomKeyPair()
	return &Sender{
		client:   client,
		symKeyID: symKeyID,
		user:     user,
		keypair:  keypair,
	}
}

func (sender *Sender) Say(msg string) {
	userMsg := common.UserMsg{
		User: *sender.user,
		Msg:  msg,
	}
	bytes, _ := json.Marshal(&userMsg)

	message := whisperv6.NewMessage{
		SymKeyID:  sender.symKeyID,
		Payload:   bytes,
		TTL:       60,
		PowTime:   2,
		PowTarget: 2.5,
		Topic:     whisperv6.BytesToTopic(common.Topic()),
		Sig:       sender.keypair,
	}

	_, err := sender.client.Post(context.Background(), message)
	if err != nil {
		log.Fatal((err))
	}
}
