package server

import (
	"fmt"
    "os"
    "io"
	"log"
    "time"
	"net/http"
    "path/filepath"
    "NetworkVizMap/cmd/packet-capture/frompcap"
	"NetworkVizMap/cmd/tcpxml2Other"
	"NetworkVizMap/config"
)

func pcap2jsonHandleFunc(w http.ResponseWriter, r *http.Request) {
    fmt.Println("file uploaded")
    uploadFilepath := fileUploadHandler(w, r)

    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    w.Write(makeReturnData(uploadFilepath))
    fmt.Println("file uploaded end")
}

func fileUploadHandler(w http.ResponseWriter, r *http.Request) (uploadFilepath string){
    
    file, fileHeader, err := r.FormFile("pcap")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    err = os.MkdirAll(config.FROMPCAP_UPLOAD_DIR, os.ModePerm)
    if err != nil {
        log.Fatal(err)
    }
    
    dstpath := fmt.Sprintf("%s/%d%s", config.FROMPCAP_UPLOAD_DIR, time.Now().UnixNano, filepath.Ext(fileHeader.Filename))
    
    dst, err := os.Create(dstpath)
    if err != nil {
        log.Fatal(err)
    }
    defer dst.Close()

    _, err = io.Copy(dst, file)
    if err != nil {
        log.Fatal(err)
    }

    
    return dst.Name()

}

func makeReturnData(inputFilePath string) (json_data []byte) {
    frompcap.AnalyzeStart(inputFilePath)
	tcpxml := frompcap.ReadXML()

	datas := tcpxml2Other.GetMarkerStruct(tcpxml)
	json_data = tcpxml2Other.GetJsonFromMarkerStruct(datas)

    return
}

func StartServer() {
    mux := http.NewServeMux()
    mux.HandleFunc("/pcap2json", pcap2jsonHandleFunc)

    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal(err)
    }
}