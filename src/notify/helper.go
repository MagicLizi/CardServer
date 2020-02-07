package notify

import (
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
)

var notifyMap = make(map[string]uint16)

func InitNotifyConf() {
	notifyMap["RefreshCenterShop"] = 5
	notifyMap["PlayerRoomInfo"] = 6
}

func GetNotifyIdWithKey(key string) uint16 {
	return notifyMap[key]
}

func ToProtoNotify(protoStruct proto.Message, msgTypeBuf []byte, conn *websocket.Conn) {
	if data, err := proto.Marshal(protoStruct); err == nil {
		lastData := append(msgTypeBuf, data...)
		//log.Println("lastData", lastData)
		conn.WriteMessage(websocket.BinaryMessage, lastData)
	} else {
		log.Println("to proto error", err)
	}
}
