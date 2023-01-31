package cmd

import (
	"fmt"
    "log"
	"io"
	"os/exec"
)

func Tcpflow() {
	cmd := exec.Command("tcpflow")
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

func Say(){
    fmt.Println("hello!")
}