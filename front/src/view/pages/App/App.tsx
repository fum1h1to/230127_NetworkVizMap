import { useCallback, useEffect, useRef, useState } from "react";
import VizMap from "../VizMap/VizMap";
import "./App.css";

function App() {
  const [message, setMessage] = useState("");
  const webSocketRef = useRef<WebSocket>();

  // useEffect(() => {
  //   const socket = new WebSocket("ws://localhost:8080/ws");
  //   webSocketRef.current = socket;

  //   socket.addEventListener("message", event => {
  //     setMessage(event.data);
  //   });

  //   return () => socket.close();
  // }, []);

  const [inputText, setInputText] = useState("");
  const [markers, setMarkers] = useState([
    {
      from: {
        lat: 40.7449041,
        lng: -73.9886504,
      },
      to: {
        lat: 35.6812405,
        lng: 139.7649361,
      },
      srcport: 80,
      dstport: 80,
    },
    {
      from: {
        lat: 35.14730734477858,
        lng: 129.03368930249715,
      },
      to: {
        lat: 35.6812405,
        lng: 139.7649361,
      },
      srcport: 80,
      dstport: 80,
    },
    {
      from: {
        lat: 47.925458609308635,
        lng: 106.89603328136243,
      },
      to: {
        lat: 35.6812405,
        lng: 139.7649361,
      },
      srcport: 80,
      dstport: 80,
    },
    {
      from: {
        lat: -12.034662151727536,
        lng: -77.04546569604334,
      },
      to: {
        lat: 35.6812405,
        lng: 139.7649361,
      },
      srcport: 80,
      dstport: 80,
    },
    {
      from: {
        lat: 14.708033187842627,
        lng: -17.444882846679775,
      },
      to: {
        lat: 35.6812405,
        lng: 139.7649361,
      },
      srcport: 80,
      dstport: 80,
    },
  ])
  // const submit: React.FormEventHandler = useCallback(
  //   event => {
  //     event.preventDefault();
  //     webSocketRef.current?.send(inputText);
  //   },
  //   [inputText],
  // );

  const click = (e: any) => {
    e.preventDefault();
    setMarkers([
      {
        from: {
          lat: 50.708033187842627,
          lng: -10.444882846679775,
        },
        to: {
          lat: 35.6812405,
          lng: 139.7649361,
        },
        srcport: 80,
        dstport: 80,
      },
    ])
  }

  return (
    <div className="App">
      <h1>{JSON.stringify(message)}</h1>
      <form>
        <input value={inputText} onChange={e => setInputText(e.target.value)} />
        <button onClick={click}>送信</button>
      </form>
      <div className="vizMapArea">
        <VizMap 
          markers={ markers }
        />
      </div>
    </div>
  );
}

export default App;