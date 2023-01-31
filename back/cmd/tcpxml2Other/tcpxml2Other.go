package tcpxml2Other

import (
	"NetworkVizMap/cmd/packet-capture"
	"NetworkVizMap/cmd/ip2LatLng"
	"NetworkVizMap/config"
	"encoding/json"
	// "fmt"
	"strconv"
	"log"
    "io/ioutil"
)

type MarkerStruct struct {
	From struct {
		Lat float64	`json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"from"`
	Srcip string `json:"srcip"`
	To struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"to"`
	Dstip string `json:"dstip"`
	Srcport int `json:"srcport"`
	Dstport int	`json:"dstport"`
}

func GetMarkerStruct(tcpflowXML *types.TcpflowXML) (datas *[]MarkerStruct){
	markers := []MarkerStruct{}

	for _, v := range tcpflowXML.Configuration.Fileobject {
		src_ipn := v.Tcpflow.Src_ipn
		dst_ipn := v.Tcpflow.Dst_ipn
		srcport := v.Tcpflow.Srcport
		dstport := v.Tcpflow.Dstport

		marker := new(MarkerStruct)

		src_ipn_lat, src_ipn_lng := ip2LatLng.GetLatLng(src_ipn)
		dst_ipn_lat, dst_ipn_lng := ip2LatLng.GetLatLng(dst_ipn)

		marker.From.Lat = src_ipn_lat
		marker.From.Lng = src_ipn_lng
		marker.To.Lat = dst_ipn_lat
		marker.To.Lng = dst_ipn_lng
		marker.Srcip = src_ipn
		marker.Dstip = dst_ipn
		marker.Srcport, _ = strconv.Atoi(srcport)
		marker.Dstport, _ = strconv.Atoi(dstport)

		markers = append(markers, *marker)
	}

	return &markers
}

func GetJsonFromMarkerStruct(datas *[]MarkerStruct) (json_data []byte){
	json_data, err := json.Marshal(datas)
	if err != nil {
		log.Fatal(err)
	}

	file, _ := json.MarshalIndent(datas, "", " ")
	_ = ioutil.WriteFile(config.OUTPUT_ROOT_DIR + "/frompcap/report.json", file, 0644)
	return
}