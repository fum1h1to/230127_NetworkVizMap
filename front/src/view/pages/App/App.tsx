import { useEffect, useRef, useState } from "react";
import VizMap from "../VizMap/VizMap";
import "./App.css";

function App() {
  const webSocketRef = useRef<WebSocket>();
  const [markers, setMarkers] = useState([])
  const [myipv4, setMyipv4] = useState<string | undefined>()
  const [myipv6, setMyipv6] = useState<string | undefined>()
  const [center, setCenter] = useState<{lat: number, lng: number}>({lat: 1, lng: 1})
  const [fromORto, setFromORto] = useState<"from" | "to" | "all">("all")

  useEffect(() => {
    const socket = new WebSocket("ws://localhost:8080/ws");
    webSocketRef.current = socket;

    socket.addEventListener("message", event => {
      const convert_json = JSON.parse(event.data);
      setMarkers(convert_json)
    });

    return () => socket.close();
  }, []);

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    setMyipv4(String(formData.get("myipv4")))
    setMyipv6(String(formData.get("myipv6")))
    setCenter({lat: Number(formData.get("lat")), lng: Number(formData.get("lng"))})
    setFromORto(String(formData.get("fromORto")) as "from" | "to" | "all")
  }

  return (
    <div className="App">
      <div className="controlBoxArea">
        <form onSubmit={handleSubmit}>
          <div>
            <label>your IPv4</label>
            <input name="myipv4" placeholder="127.0.0.1" />
          </div>
          <div>
            <label>your IPv6</label>
            <input name="myipv6" placeholder=":::1" />
          </div>
          <div>
            <label>your Position</label>
            <input name="lat" placeholder="lat" />
            <input name="lng" placeholder="lng" />
          </div>
          <div>
            <label>visible</label>
            <select name="fromORto">
              <option value="from" selected>other → your</option>
              <option value="to">your → other</option>
              <option value="all">all</option>
            </select>
          </div>
          <div>
            <button type="submit">更新</button>
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