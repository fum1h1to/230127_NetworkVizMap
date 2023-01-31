package realtime

import (
	"NetworkVizMap/config"
	"NetworkVizMap/cmd/packet-capture"
	"os/exec"
	"time"
	"log"
	"syscall"
	"encoding/xml"
	"io/ioutil"
)

var outputDirName string
var cmd *exec.Cmd

func GetTime(isTokyo bool) (str string) {
	tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}
	t := time.Now()
	if isTokyo {
		t = t.In(tokyo)
	}

	return t.Format("2006-01-02_15-04-05")
}

func Tcpflow() {
	
	outputDirName = GetTime(false)
	outputDirName = config.OUTPUT_ROOT_DIR + "/realtime/" + outputDirName

	cmd = exec.Command("tcpflow", "-o", outputDirName)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	err := cmd.Start()

    if err != nil {
		log.Fatal(err)
	}
}

func StopTcpflow() {
	err := syscall.Kill(-cmd.Process.Pid, syscall.SIGINT)

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
