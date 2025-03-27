package game

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

type MsgType byte

const (
	INIT    MsgType = 0
	CONTROL MsgType = 1
	UPDATE  MsgType = 2
)

type ControlMsgType byte

const (
	START_SOLVING ControlMsgType = 0
)

type WsHandle struct {
	conn *websocket.Conn
}

func NewWsHandle(conn *websocket.Conn) *WsHandle {
	h := new(WsHandle)
	h.conn = conn
	return h
}

func (handle *WsHandle) parseControlMessage(msgType ControlMsgType, data []byte) bool {
	return msgType == ControlMsgType(START_SOLVING)

}

func (handle *WsHandle) ParseMessage(data []byte) bool {
	msgType := MsgType(data[0])
	switch msgType {
	case MsgType(CONTROL):
		return handle.parseControlMessage(ControlMsgType(data[1]), data[2:])

	}
	return false
}

func (handle *WsHandle) sendMessage(data []byte, msgType MsgType) error {

	conn := handle.conn
	if conn == nil {
		return nil
	}
	err := conn.WriteMessage(websocket.BinaryMessage, append([]byte{byte(msgType)}, data...))

	if err != nil {
		fmt.Println("Error while writing message:", err)
		return err
	}
	return nil
}

func (handle *WsHandle) SendToWs(game *Game) error {
	conn := handle.conn
	if conn == nil {
		return nil
	}
	fmt.Println("Sending game to ws")
	var data [81]byte
	for i := range 81 {
		data[i] = byte(game._data[i].num)
	}
	defer time.Sleep(30 * time.Millisecond)
	return handle.sendMessage(data[:], MsgType(UPDATE))
}
