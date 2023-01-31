package main

import (
	"NetworkVizMap/cmd/packet-capture/frompcap"
	"NetworkVizMap/cmd/make-data"
	"fmt"
	"encoding/json"
	// "time"
	"log"
)


func main() {
   	frompcap.AnalyzeStart("test.pcap.pcapng")
	tcpxml := frompcap.ReadXML()

	datas := makedata.MakeData(tcpxml)
	json_data, err := json.Marshal(datas)

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("string(sj)=%+v\n\n", string(json_data))
}