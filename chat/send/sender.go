package send

import (
	"context"
	"encoding/json"
	"log"
	"github.com/jpsiyu/ethtut-cli/chat/common"

	"github.com/ethereum/go-ethereum/whisper/shhclient"
	"github.com/ethereum/go-ethereum/whisper/whisperv6"
)

type Sender struct {
	client *shhclient.Client
	keyID  string
	pubKey []byte
	user   *common.User
}

func NewSender(user *common.User, client *shhclient.Client, keyID string) *Sender {
	pub, err := client.PublicKey(context.Background(), keyID)
	if err != nil {
		log.Fatal((err))
	}
	return &Sender{
		client: client,
		keyID:  keyID,
		pubKey: pub,
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
		PublicKey: sender.pubKey,
		Payload:   bytes,
		TTL:       60,
		PowTime:   2,
		PowTarget: 2.5,
	}

	_, err := sender.client.Post(context.Background(), message)
	if err != nil {
		log.Fatal((err))
	}
}
