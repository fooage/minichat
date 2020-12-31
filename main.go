package main

import (
	"flag"
	"fmt"

	"github.com/fooage/minichat/console"
	"github.com/fooage/minichat/handle"

	"github.com/jroimartin/gocui"
)

var (
	// Default ip address is 127.0.0.1, default listening port is 10000.
	local  = flag.String("h", "127.0.0.1:10000", "local listen host and listen port")
	remote = flag.String("t", "127.0.0.1:11000", "target client's ip address and port")
	key    = flag.String("k", "8SMEE7ieNjSWVFqq", "key for data encryption and decryption")
)

func main() {
	// Create handler and let it start work.
	handler := handle.NewHandler()
	handler.LocalAddr = *local
	handler.RemoteAddr = *remote
	handler.AesKey = []byte(*key)
	// Show the console interface.
	cui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		fmt.Println(err)
		return
	}
	console.SetHandler(handler)
	console.RunInterface(cui)
}
