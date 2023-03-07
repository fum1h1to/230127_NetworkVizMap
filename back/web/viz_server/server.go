package viz_server

import (
	"fmt"
    "os"
    "io"
	"log"
    "time"
	"net/http"
    "path/filepath"

    "NetworkVizMap/packet_capture/frompcap"
	"NetworkVizMap/util/ip2LatLng"
	"NetworkVizMap/config"
)

type WebServer struct {
    ip2LatLngExchanger *ip2LatLng.Ip2LatLngExchanger
}

func CreateWebServer() *WebServer {
    return &WebServer{
        ip2LatLngExchanger: ip2LatLng.CreateIp2LatLngExchanger(),
    }
}

func (s *WebServer) StartServer() {
    mux := http.NewServeMux()
    mux.HandleFunc("/pcap2json", s.pcap2jsonHandleFunc)

    log.Println("server started at 8080")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal(err)
    }
}

/*
    pcapを受け取りjsonにして返す。
*/
func (s *WebServer) pcap2jsonHandleFunc(w http.ResponseWriter, r *http.Request) {
    log.Println("file uploaded")
    uploadFilepath := s.fileUploadHandler(w, r)

    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    w.Write(s.makeReturnData(uploadFilepath))
    log.Println("file uploaded end")
}

/*
    pcapファイルを受け取り、
    config.FROMPCAP_UPLOAD_DIR配下に保存する
    そして、保存する際に付けたファイル名を返す
*/
func (s *WebServer) fileUploadHandler(w http.ResponseWriter, r *http.Request) (uploadFilepath string){
    
    file, fileHeader, err := r.FormFile("pcap")
    if err != nil {
        log.Println(err)
        return ""
    }
    defer file.Close()

    err = os.MkdirAll(config.FROMPCAP_UPLOAD_DIR, os.ModePerm)
    if err != nil {
        log.Println(err)
        return ""
    }
    
    dstpath := fmt.Sprintf("%s/%d%s", config.FROMPCAP_UPLOAD_DIR, time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))
    
    dst, err := os.Create(dstpath)
    if err != nil {
        log.Println(err)
        return ""
    }
    defer dst.Close()

    _, err = io.Copy(dst, file)
    if err != nil {
        log.Println(err)
        return ""
    }

    
    return dst.Name()
}

/*
    保存されたpcapファイルのパスを受け取り、
    jsonに変換して返す。
    inputFilePathがなかったの場合は、errorの旨のjsonを返す。
*/
func (s *WebServer) makeReturnData(inputFilePath string) (json_data []byte) {
    pcapAnalyzer := frompcap.CreatePcapAnalyzer(inputFilePath, s.ip2LatLngExchanger)
    
    return pcapAnalyzer.AnalyzeStartAndGetResult()
}