import { useCallback, useEffect, useRef, useState } from "react";
import VizMap from "../VizMap/VizMap";
import "./App.css";

function App() {
  const [message, setMessage] = useState("");
  const webSocketRef = useRef<WebSocket>();

  useEffect(() => {
    const socket = new WebSocket("ws://localhost:8080/ws");
    webSocketRef.current = socket;

    socket.addEventListener("message", event => {
      const convert_json = JSON.parse(event.data);
      setMarkers(convert_json)
    });

    return () => socket.close();
  }, []);

  const [inputText, setInputText] = useState("");
  const [markers, setMarkers] = useState([])

  return (
    <div className="App">
      {/* <h1>{JSON.stringify(message)}</h1>
      <form>
        <input value={inputText} onChange={e => setInputText(e.target.value)} />
        <button onClick={submit}>送信</button>
      </form> */}
      <div className="vizMapArea">
        <VizMap 
          markers={ markers }
        />
      </div>
    </div>
  );
}

export default App;