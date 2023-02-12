import { useEffect, useRef, useState } from "react";
import VizMap from "../VizMap/VizMap";
import "./App.css";

function App() {
  // const webSocketRef = useRef<WebSocket>();
  const [isClick, setIsClick] = useState(false);
  const [pcap, setPcap] = useState<File | null>(null)
  const [markers, setMarkers] = useState([])
  const [myipv4, setMyipv4] = useState<string | undefined>()
  const [myipv6, setMyipv6] = useState<string | undefined>()
  const [center, setCenter] = useState<{lat: number, lng: number}>({lat: 35, lng: 140})
  const [fromORto, setFromORto] = useState<"from" | "to" | "all">("all")

  // useEffect(() => {
  //   const socket = new WebSocket("ws://localhost:8080/ws");
  //   webSocketRef.current = socket;

  //   socket.addEventListener("message", event => {
  //     const convert_json = JSON.parse(event.data);
  //     setMarkers(convert_json)
  //   });

  //   return () => socket.close();
  // }, []);

  const onChangePcap = (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = e.target.files;
    if (files && files[0]) {
      setPcap(files[0])
    } else {
      setPcap(null)
    }
  }

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setIsClick(true);
    const formData = new FormData(e.currentTarget);
    
    if (pcap) {
      const send_formData = new FormData();
      send_formData.append("pcap", pcap)
      await fetch("http://localhost:8080/pcap2json", {
        method: "POST",
        credentials: 'same-origin',
        // headers: {
        //   'Content-Type': 'multipart/form-data',
        // }, // これはいらないっぽい
        body: send_formData
      })
      .then(res => {
        return res.json()
      })
      .then(res => {
        setMarkers(res)
      })
      .catch(err => {
        console.log(err)
      })
    }
    
    setMyipv4(String(formData.get("myipv4")))
    setMyipv6(String(formData.get("myipv6")))
    let centerLat = Number(formData.get("lat"))
    let centerLng = Number(formData.get("lng"))
    if (centerLat == 0 && centerLng == 0) {
      centerLat = 35
      centerLng = 140
    }
    setCenter({lat: centerLat, lng: centerLng})
    setFromORto(String(formData.get("fromORto")) as "from" | "to" | "all")
    setPcap(null)
    setIsClick(false);
  }

  return (
    <div className="App">
      <div className="controlBoxArea">
        <form onSubmit={handleSubmit}>
          <div className="controlPanel">
            <label>your IPv4</label>
            <input name="myipv4" placeholder="127.0.0.1" />
          </div>
          <div className="controlPanel">
            <label>your IPv6</label>
            <input name="myipv6" placeholder=":::1" />
          </div>
          <div className="controlPanel">
            <label>your Position</label>
            <input name="lat" placeholder="lat (default: 35)" />
            <input name="lng" placeholder="lng (default: 140)" />
          </div>
          <div className="controlPanel">
            <label>visible</label>
            <select name="fromORto">
              <option value="from" selected>other → your</option>
              <option value="to">your → other</option>
              <option value="all">all</option>
            </select>
          </div>
          <div className="controlPanel">
            <label>pcap file</label>
            <input type="file" name="pcap" onChange={onChangePcap} />
          </div>
          <div className="controlPanel">
            <button type="submit" disabled={isClick}>更新</button>
            { isClick ? <span>loading...</span> : <></>}
          </div>
        </form>
      </div>
      <div className="vizMapArea">
        <VizMap
          myipv4={myipv4}
          myipv6={myipv6}
          center={center}
          fromORto={fromORto}
          markers={ markers }
        />
      </div>
    </div>
  );
}

export default App;