package models

import (
	
)

type PacketData struct {
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

type ErrorJson struct {
	Error string `json:"error"`
}