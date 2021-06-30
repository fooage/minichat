package main

import (
	"fmt"

	"github.com/fooage/minichat/config"
	"github.com/fooage/minichat/console"
	"github.com/fooage/minichat/handle"

	"github.com/jroimartin/gocui"
)

// Default ip address is 127.0.0.1, default listening port is 10000.

func main() {
	conf := config.LoadConfig(".")
	if conf == nil {
		return
	}
	handler := handle.NewHandler()
	handler.LocalAddr = conf.Local
	handler.RemoteAddr = conf.Remote
	handler.AesKey = []byte(conf.Key)
	// show the console interface
	cui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		fmt.Println(err)
		return
	}
	console.SetHandler(handler)
	console.RunInterface(cui)
}
