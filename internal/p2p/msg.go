package p2p

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
)

type MessageType string

const MsgAddressList MessageType = "address_list"
const MsgEcho MessageType = "echo"
const MsgJoin MessageType = "join"

// const MsgNewBlock
// const MsgNewTransaction

type Message struct {
	Type    MessageType
	Payload []byte
}

func SendMessage(conn net.Conn, msg Message) {

	data, err := json.Marshal(msg)
	if err != nil {
		log.Fatal("[error] failed to marshal message ", err)
	}

	length := int32(len(data))

	err = binary.Write(conn, binary.BigEndian, length)
	if err != nil {
		log.Fatal("[error] failed to send data length")
	}

	_, err = conn.Write(data)
	if err != nil {
		log.Fatal("[error] failed to send data")
	}
}

func ListenMessage(conn net.Conn, MsgCh chan Message) {
	defer conn.Close()

	for {
		var length int32
		err := binary.Read(conn, binary.BigEndian, &length)
		if err != nil {
			if err == io.EOF {
				fmt.Println("connection lost")
			} else {
				fmt.Println("[error] failed getting packet length")
			}
			return
		}

		data := make([]byte, length)

		_, err = io.ReadFull(conn, data)
		if err != nil {
			log.Fatal("[error] failed to read package data")
			return
		}

		var msg Message
		err = json.Unmarshal(data, &msg)
		if err != nil {
			fmt.Println("[error] failed to unmarshal data")
			continue
		}

		MsgCh <- msg
	}
}

// func HandleConnection(c net.Conn) {
// 	c.Close()

// }

// func MsgManager(MsgCh chan Message) {
// 	for msg := range MsgCh {
// 		if msg.Type == "join" {
// 			n.Peers = append(n.Peers, conn)

// 			fmt.Printf("[P2P][New msg] Type: %s, Payload %s\n", msg.Type, string(msg.Payload))
// 			p2p.SendMessage(conn, p2p.Message{Type: "msg", Payload: []byte("successful connection")})
// 		}

// 		if msg.Type == "echo" {

// 			fmt.Printf("[P2P] Echo\n")
// 		}
// 	}
// }
