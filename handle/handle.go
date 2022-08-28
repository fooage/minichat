package handle

import (
	"fmt"
	"net"
	"strings"
)

// Handler is a structure that accepts and sends messages and is connected to
// the user interface through the channel.
type Handler struct {
	// message receive channel
	Buf chan []byte

	// key for AES encryption
	AesKey []byte

	// listen address and remote address
	LocalAddr  string
	RemoteAddr string
}

// NewHandler is constructor of handler.
func NewHandler() *Handler {
	return &Handler{
		Buf: make(chan []byte, 1024),
	}
}

// Recv function is used as receiving message usually need a goroutine to run
// this function asynchronously. And it is necessary to ensure thread safety
// when combined with some user interface.
func (h *Handler) Recv() {
	addr, err := net.ResolveUDPAddr("udp", h.LocalAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// listening for messages and receive
	for {
		data := make([]byte, 1024)
		_, _, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		orig := decrypt(data, h.AesKey)
		h.Buf <- orig
	}
}

// Send the message to the other client, and the size of buffer is 1024 bytes.
func (h *Handler) Send(orig []byte) {
	remoteAddr, err := net.ResolveUDPAddr("udp", h.RemoteAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.DialUDP("udp", nil, remoteAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	ip := []byte("[" + strings.Split(h.LocalAddr, ":")[0] + "]" + ":")
	orig = append(ip, orig...)
	data := encrypt(orig, h.AesKey)
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	h.Buf <- orig
}
