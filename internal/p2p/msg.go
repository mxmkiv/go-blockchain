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

type MsgManager struct {
	conn  net.Conn
	MsgCh chan Message
}

func NewManager() *MsgManager {
	return &MsgManager{}
}

func (m *MsgManager) Init(conn net.Conn, chSize int) {
	m.conn = conn
	m.MsgCh = make(chan Message, chSize)
	go m.ListenMessage()
}

func (m *MsgManager) SendMessage(msg Message) {

	data, err := json.Marshal(msg)
	if err != nil {
		log.Fatal("[error] failed to marshal message ", err)
	}

	length := int32(len(data))

	err = binary.Write(m.conn, binary.BigEndian, length)
	if err != nil {
		log.Fatal("[error] failed to send data length")
	}

	_, err = m.conn.Write(data)
	if err != nil {
		log.Fatal("[error] failed to send data")
	}
}

func (m *MsgManager) ListenMessage() {
	defer m.conn.Close()

	for {
		var length int32
		err := binary.Read(m.conn, binary.BigEndian, &length)
		if err != nil {
			if err == io.EOF {
				fmt.Println("connection lost")
			} else {
				fmt.Println("[error] failed getting packet length")
			}
			return
		}

		data := make([]byte, length)

		_, err = io.ReadFull(m.conn, data)
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

		m.MsgCh <- msg
	}
}
