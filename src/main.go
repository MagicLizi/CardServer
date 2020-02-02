// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"./handler"
	"encoding/binary"
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8888", "http service address")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ws(w http.ResponseWriter, r *http.Request) {
	conn, connectErr := upgrader.Upgrade(w, r, nil)

	if connectErr != nil {
		log.Println(connectErr)
		return
	}

	defer conn.Close()

	for {
		messageType, p, readErr := conn.ReadMessage()
		if readErr != nil {
			log.Println(readErr)
			return
		}

		if messageType == websocket.BinaryMessage && len(p) > 2 {
			log.Println("receive buffers:",p)
			msgTypeBuf := p[0:2]

			protoData := p[2:len(p)]

			msgTyp := binary.BigEndian.Uint16(msgTypeBuf)

			log.Println("msgTypeBuf buffers:",msgTypeBuf,msgTyp)
			log.Println("protoData buffers:",protoData)

			curHandler := handler.GetHandlerByMsgId(msgTyp)
			if curHandler!= nil {
				curHandler(protoData, msgTypeBuf, conn)
			}
		} else {
			log.Println("not valid message type or data_buff < 2")
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/ws", ws)
	handler.InitHandlerConf()
	log.Fatal(http.ListenAndServe(*addr, nil))
}