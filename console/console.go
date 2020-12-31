package console

import (
	"fmt"

	"github.com/fooage/minichat/handle"

	"github.com/jroimartin/gocui"
)

// Please notice that in Chinese and Japanese fonts windows system, the default
// encoding format of the terminal needs to be changed to UTF-8.

// Realize the link between the message in the interface layer and the service layer.
var handler *handle.Handler

func outputView(g *gocui.Gui, x0, y0, x1, y1 int) error {
	v, err := g.SetView("output", x0, y0, x1, y1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Autoscroll = true
		v.Overwrite = false
		v.Title = "Messages"
	}
	return nil
}

func inputView(g *gocui.Gui, x0, y0, x1, y1 int) error {
	if v, err := g.SetView("input", x0, y0, x1, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Editable = true
		v.Overwrite = false
		v.Title = "Talk"
		// Set the input window to focus.
		if _, err := g.SetCurrentView("input"); err != nil {
			return err
		}
	}
	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if err := outputView(g, 1, 0, maxX-2, maxY-2); err != nil {
		return err
	}
	if err := inputView(g, 1, maxY-3, maxX-2, maxY-1); err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func send(g *gocui.Gui, v *gocui.View) error {
	msg := v.Buffer()
	length := len([]byte(msg))
	v.MoveCursor(0-length, 0, true)
	v.Clear()
	// Post the message to the other chat client and add it to the buffer
	// channel use the handler which minichat used.
	handler.Send([]byte(msg))
	return nil
}

func recv(g *gocui.Gui) {
	// Start goroutine to listen for messages.
	go handler.Recv()
	msg := make([]byte, 1024)
	for {
		select {
		case msg = <-handler.Buf:
			// This function is very important. In order to make the operation
			// thread safe, need to write the operation into Update function.
			g.Update(func(g *gocui.Gui) error {
				v, err := g.View("output")
				if err != nil {
					fmt.Println(err)
					return err
				}
				v.Write(msg)
				return nil
			})
		}
	}
}

// SetHandler is function that set the handler.
func SetHandler(h *handle.Handler) {
	handler = h
}

// RunInterface function is the display function of the user interface.
func RunInterface(g *gocui.Gui) {
	defer func() {
		g.Close()
	}()
	g.SetManagerFunc(layout)
	if err := g.SetKeybinding("input", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		fmt.Println(err)
		return
	}
	if err := g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, send); err != nil {
		fmt.Println(err)
		return
	}
	// Start to read the channel and refresh the interface.
	go recv(g)
	if err := g.MainLoop(); err != nil {
		fmt.Println(err)
	}
}
