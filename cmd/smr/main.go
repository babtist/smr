package main

import (
	"github.com/sinchop/smpp"
)

func main() {

	s := server.NewServer("test", 3601,
		func(c server.Conn, submit *message.ShortMessage) (*message.ShortMessageResp, error) {
			return &message.ShortMessageResp{
				MessageID: "1234",
				Status:    message.Status_OK,
			}, nil
		})

}
