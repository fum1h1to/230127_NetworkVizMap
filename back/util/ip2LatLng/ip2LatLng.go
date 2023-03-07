package ip2LatLng

import (
	"log"
	"net"

	"NetworkVizMap/config"
	"github.com/oschwald/geoip2-golang"
)

type Ip2LatLngExchanger struct {
	Db *geoip2.Reader
}

func CreateIp2LatLngExchanger() *Ip2LatLngExchanger {
	Db, err := geoip2.Open(config.GEOIP_DB_PATH)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("ip2LatLng init success")

	return &Ip2LatLngExchanger{
		Db: Db,
	}
}

func (e *Ip2LatLngExchanger) GetLatLng(ip string) (lat float64, lng float64, err error) {
	parsedIp := net.ParseIP(ip)
	record, err := e.Db.City(parsedIp)
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}
	
	return record.Location.Latitude, record.Location.Longitude, nil
}