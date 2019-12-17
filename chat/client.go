package chat

import (
	"bufio"
	"fmt"
	"github.com/jpsiyu/ethtut-cli/chat/common"
	"log"
	"os"
	"strings"

	"github.com/jpsiyu/ethtut-cli/chat/receive"
	"github.com/jpsiyu/ethtut-cli/chat/send"
	"github.com/jpsiyu/ethtut-cli/conf"

	tm "github.com/buger/goterm"
	"github.com/ethereum/go-ethereum/whisper/shhclient"
)

func clear() {
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Flush()
}

func Run() {
	clear()
	keyID := common.GenKey()
	client, err := shhclient.Dial(conf.ShhUrl)
	if err != nil {
		log.Fatal((err))
	}

	user := common.RandomUser()

	receiver := receive.NewReceiver(&user, client, keyID)
	go receiver.Run()

	sender := send.NewSender(&user, client, keyID)
	fmt.Println("Enter your message:")
	for {
		reader := bufio.NewReader(os.Stdin)
		msg, _ := reader.ReadString('\n')
		msg = strings.Replace(msg, "\n", "", -1)
		sender.Say(msg)
	}
}
