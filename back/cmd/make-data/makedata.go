package makedata

import (
	"NetworkVizMap/cmd/packet-capture"
	"NetworkVizMap/cmd/ip2LatLng"
	// "fmt"
	"strconv"
)

type MarkerSchema struct {
	From struct {
		Lat float64	`json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"from"`
	To struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"to"`
	Srcport int `json:"srcport"`
	Dstport int	`json:"dstport"`
}

func MakeData(tcpflowXML *types.TcpflowXML) (datas *[]MarkerSchema){
	markers := []MarkerSchema{}

	for _, v := range tcpflowXML.Configuration.Fileobject {
		src_ipn := v.Tcpflow.Src_ipn
		dst_ipn := v.Tcpflow.Dst_ipn
		srcport := v.Tcpflow.Srcport
		dstport := v.Tcpflow.Dstport

		marker := new(MarkerSchema)

		src_ipn_lat, src_ipn_lng := ip2LatLng.GetLatLng(src_ipn)
		dst_ipn_lat, dst_ipn_lng := ip2LatLng.GetLatLng(dst_ipn)

		marker.From.Lat = src_ipn_lat
		marker.From.Lng = src_ipn_lng
		marker.To.Lat = dst_ipn_lat
		marker.To.Lng = dst_ipn_lng
		marker.Srcport, _ = strconv.Atoi(srcport)
		marker.Dstport, _ = strconv.Atoi(dstport)

		markers = append(markers, *marker)
	}

	return &markers
}