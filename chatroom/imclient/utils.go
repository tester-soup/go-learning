package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-learning/chatroom/common/message"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Printf("conn.Read fail err=%v\n", err)

	}

	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Printf("json.Unmarshal err=%v\n", err)
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	// 发送包长度

	// 定义消息长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	// 将消息长度在转成字节
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	// 发送消息长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Printf("conn.Write err=%v\n", err)
		return
	}

	// 发送data本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Printf("conn.Write err=%v\n", err)
		return
	}
	return
}
