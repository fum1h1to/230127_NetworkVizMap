package server

import (
	// "fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
    "NetworkVizMap/cmd/packet-capture/frompcap"
	"NetworkVizMap/cmd/tcpxml2Other"
	// "NetworkVizMap/config"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    // Add this lines
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func webSocketHandleFunc(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("upgrade:", err)
        return
    }
    defer ws.Close()

    json_data := makeReturnData("test.pcap")

    err = ws.WriteMessage(websocket.TextMessage, json_data)
    if err != nil {
        log.Println("WriteMessage:", err)
        return
    }
}

func makeReturnData(inputFilePath string) (json_data []byte) {
    frompcap.AnalyzeStart(inputFilePath)
	tcpxml := frompcap.ReadXML()

	datas := tcpxml2Other.GetMarkerStruct(tcpxml)
	json_data = tcpxml2Other.GetJsonFromMarkerStruct(datas)

    return
}

func StartServer() {
    http.HandleFunc("/ws", webSocketHandleFunc)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}