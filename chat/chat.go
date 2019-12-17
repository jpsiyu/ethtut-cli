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
	// user
	user := common.RandomUser()

	// view
	app := tview.NewApplication()
	input := tview.NewInputField()
	input.SetLabel(user.Name + ": ")
	input.SetFieldWidth(300)
	input.SetFieldBackgroundColor(tcell.ColorBlack)
	input.SetFieldTextColor(tcell.ColorWhite)
	table := tview.NewTable()
	grid := tview.NewGrid()
	grid.SetRows(0, 5)
	grid.SetBorders(true)
	grid.AddItem(table, 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(input, 1, 0, 1, 1, 0, 0, false)
	app.SetRoot(grid, true)
	app.SetFocus(input)

	// shh client
	symKeyID, err := common.GenSymKey()
	if err != nil {
		log.Fatal((err))
	}

	client, err := shhclient.Dial(conf.ShhUrl)
	if err != nil {
		log.Fatal((err))
	}
	sender := send.NewSender(&user, client, symKeyID)

	var handler = func(text string) {
		count := table.GetRowCount()
		table.SetCell(count, 0, tview.NewTableCell(text))
		app.Draw()
	}
	receiver := receive.NewReceiver(&user, client, symKeyID, handler)
	go receiver.Run()

	// bind view and logic
	input.SetFinishedFunc(func(key tcell.Key) {
		text := input.GetText()
		sender.Say(text)
		input.SetText("")
	})

	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
