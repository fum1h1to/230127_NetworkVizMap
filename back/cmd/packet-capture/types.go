package types

import (
	"encoding/xml"
)

type TcpflowXML struct {
	XMLName xml.Name `xml:"dfxml"`
	Version string `xml:"xmloutputversion,attr"`
	Configuration Configuration `xml:"configuration"`
}

type Configuration struct {
	Fileobject []Fileobject `xml:"fileobject"`
}

type Fileobject struct {
	Tcpflow Tcpflow_xml `xml:"tcpflow"`
}

type Tcpflow_xml struct {
	Startime string `xml:"startime,attr"`
	Endtime string `xml:"endtime,attr"`
	Mac_daddr string `xml:"mac_daddr,attr"`
	Mac_saddr string `xml:"mac_saddr,attr"`
	Family string `xml:"family,attr"`
	Src_ipn string `xml:"src_ipn,attr"`
	Dst_ipn string `xml:"dst_ipn,attr"`
	Srcport string `xml:"srcport,attr"`
	Dstport string `xml:"dstport,attr"`
	Packets string `xml:"packets,attr"`
	Len string `xml:"len,attr"`
}