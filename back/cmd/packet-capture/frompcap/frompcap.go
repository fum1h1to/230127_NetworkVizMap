package frompcap

import (
	"NetworkVizMap/config"
	"NetworkVizMap/cmd/packet-capture"
	"os/exec"
	// "time"
	"log"
	"encoding/xml"
	"io/ioutil"
)

var outputDirName string
var cmd *exec.Cmd

func AnalyzeStart(analyzeFilePath string) (){
	
	outputDirName = config.OUTPUT_ROOT_DIR + "/frompcap"

	cmd = exec.Command("tcpflow", "-o", outputDirName, "-r", analyzeFilePath, "-FX")

	err := cmd.Run()

    if err != nil {
		log.Fatal(err)
	}
}

func ReadXML() (parsedData *types.TcpflowXML) {
	readData, err := ioutil.ReadFile(outputDirName + "/" + config.REPORT_FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}

	xmlData := new(types.TcpflowXML)
	if err := xml.Unmarshal(readData, xmlData); err != nil {
		log.Fatal(err)
	}

	return xmlData
}
