package chat

import (
	"log"

	"github.com/ethereum/go-ethereum/whisper/shhclient"
	"github.com/gdamore/tcell"
	"github.com/jpsiyu/ethtut-cli/chat/common"
	"github.com/jpsiyu/ethtut-cli/chat/receive"
	"github.com/jpsiyu/ethtut-cli/chat/send"
	"github.com/jpsiyu/ethtut-cli/conf"
	"github.com/rivo/tview"
)

func Run() {
	keyID := common.GenKey()
	client, err := shhclient.Dial(conf.ShhUrl)
	if err != nil {
		log.Fatal((err))
	}
	user := common.RandomUser()
	sender := send.NewSender(&user, client, keyID)

	input := tview.NewInputField()
	input.SetLabel("Input message: ")
	input.SetFieldWidth(300)
	input.SetFieldBackgroundColor(tcell.ColorBlack)
	input.SetFieldTextColor(tcell.ColorWhite)
	input.SetDoneFunc(func(key tcell.Key) {
		text := input.GetText()
		if text != "" {
			sender.Say(text)
			input.SetText("")
		}
	})

	table := tview.NewTable()
	var handler = func(text string) {
		input.SetText("add cell")
		count := table.GetRowCount()
		table.SetCell(count, 0, tview.NewTableCell(text))
	}

	grid := tview.NewGrid()
	grid.SetRows(0, 5)
	grid.SetBorders(true)
	grid.AddItem(table, 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(input, 1, 0, 1, 1, 0, 0, false)

	receiver := receive.NewReceiver(&user, client, keyID, handler)
	go receiver.Run()

	app := tview.NewApplication()
	app.SetRoot(grid, true)
	app.SetFocus(input)
	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
