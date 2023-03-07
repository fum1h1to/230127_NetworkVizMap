package frompcap

import (
	"encoding/xml"
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"strconv"

	"NetworkVizMap/config"
	"NetworkVizMap/models"
	"NetworkVizMap/util/ip2LatLng"
)

type PcapAnalyzer struct {
	analyzeFilePath string
	outputDirName string
	cmd *exec.Cmd
	xmlData *models.TcpflowXML
	ip2LatLngExchanger *ip2LatLng.Ip2LatLngExchanger
	packetDatas []models.PacketData
	returnJsonData []byte
}

func CreatePcapAnalyzer(anylyzeFilePath string, ip2LatLngExchanger *ip2LatLng.Ip2LatLngExchanger) *PcapAnalyzer {
	outputDirName := config.OUTPUT_ROOT_DIR + "/frompcap"
	cmd := exec.Command("tcpflow", "-o", outputDirName, "-r", anylyzeFilePath, "-FX")

	return &PcapAnalyzer{
		analyzeFilePath: anylyzeFilePath,
		outputDirName: outputDirName,
		cmd: cmd,
		ip2LatLngExchanger: ip2LatLngExchanger,
	}
}

func (p *PcapAnalyzer) AnalyzeStartAndGetResult() (return_data []byte) {
	if p.analyzeFilePath == "" {
		return p.makeErrorJson()
	}

	err := p.cmd.Run()

  if err != nil {
		log.Println(err)
		return p.makeErrorJson()
	}

	if err = p.readXML(); err != nil {
		log.Println(err)
		return p.makeErrorJson()
	}

	if err = p.getPacketData(); err != nil {
		log.Println(err)
		return p.makeErrorJson()
	}

	if err = p.packetDatas2json(); err != nil {
		log.Println(err)
		return p.makeErrorJson()
	}

	return p.returnJsonData
}

func (p *PcapAnalyzer) readXML() (err error) {
	readData, err := os.ReadFile(p.outputDirName + "/" + config.REPORT_FILE_NAME)
	if err != nil {
		return
	}

	p.xmlData = new(models.TcpflowXML)
	if err = xml.Unmarshal(readData, p.xmlData); err != nil {
		return
	}
	return
}

func (p *PcapAnalyzer) getPacketData() (err error){
	p.packetDatas = []models.PacketData{}

	for _, v := range p.xmlData.Configuration.Fileobject {
		src_ipn := v.Tcpflow.Src_ipn
		dst_ipn := v.Tcpflow.Dst_ipn
		srcport := v.Tcpflow.Srcport
		dstport := v.Tcpflow.Dstport

		packetData := new(models.PacketData)

		src_ipn_lat, src_ipn_lng, err_s := p.ip2LatLngExchanger.GetLatLng(src_ipn)
		if err_s != nil {
			return err_s
		}

		dst_ipn_lat, dst_ipn_lng, err_d := p.ip2LatLngExchanger.GetLatLng(dst_ipn)
		if err_d != nil {
			return err_d
		}

		packetData.From.Lat = src_ipn_lat
		packetData.From.Lng = src_ipn_lng
		packetData.To.Lat = dst_ipn_lat
		packetData.To.Lng = dst_ipn_lng
		packetData.Srcip = src_ipn
		packetData.Dstip = dst_ipn
		packetData.Srcport, _ = strconv.Atoi(srcport)
		packetData.Dstport, _ = strconv.Atoi(dstport)

		p.packetDatas = append(p.packetDatas, *packetData)
	}

	return nil
}

func (p *PcapAnalyzer) packetDatas2json() (err error){
	p.returnJsonData, err = json.Marshal(p.packetDatas)
	if err != nil {
		return
	}
	return

	// file, _ := json.MarshalIndent(p.packetDatas, "", " ")
	// _ = os.WriteFile(config.OUTPUT_ROOT_DIR + "/frompcap/report.json", file, 0644)
	// return
}

func (p *PcapAnalyzer) makeErrorJson() (err_json []byte){
	err_json, err := json.Marshal(models.ErrorJson{Error: "error"})
	if err != nil {
		log.Panicln(err)
	}
	return
}