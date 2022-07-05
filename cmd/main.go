package main

import (
	"log"

	"github.com/rivo/tview"

	"p2p-messenger/internal/network"
	"p2p-messenger/internal/ui"
)

func main() {
	if err := runUI(); err != nil {
		log.Fatal(err)
	}
}

func runUI() error {
	networkManager := network.NewManager()
	networkManager.Start()

	app := ui.NewApp()
	if err := tview.NewApplication().SetRoot(app.View, true).SetFocus(app.Sidebar.View).Run(); err != nil {
		return err
	}
	return nil
}
