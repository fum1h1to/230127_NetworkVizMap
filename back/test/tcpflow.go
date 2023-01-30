package tcpflow

import (
	"fmt"
	"io"
	"os/exec"
	"time"
	"log"
)

func GetTime() (str string) {
	tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}
	t := time.Now().In(tokyo)

	return t.Format("2006-01-02_15-04-05")
}

func Tcpflow() {
	
	outputDirName := GetTime()
	outputDirName = "tcpflow-data/" + outputDirName

	cmd := exec.Command("tcpflow", "-o", outputDirName)
	stdout, _ := cmd.StdoutPipe()

	buff := make([]byte, 1024)

	err := cmd.Start()

	n, err := stdout.Read(buff)

    for err == nil || err != io.EOF {
		if n > 0 {
			fmt.Printf(string(buff[:n]))
		}
		n, err = stdout.Read(buff)
	}
}