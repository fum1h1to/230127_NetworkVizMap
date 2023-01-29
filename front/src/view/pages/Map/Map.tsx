import React, { useState } from "react";
import "leaflet/dist/leaflet.css";
import { MapContainer, TileLayer, Marker, Popup, Polyline } from "react-leaflet";
// add icons
import Leaflet from "leaflet";
import icon from "leaflet/dist/images/marker-icon.png";
import iconShadow from "leaflet/dist/images/marker-shadow.png";
import { GeodesicLine } from "view/components/leaflet-local/GeodesicLine/GeodesicLine";

// marker setting
let DefaultIcon = Leaflet.icon({
  iconUrl: icon,
  shadowUrl: iconShadow,
  iconAnchor: [12.5, 41],
  popupAnchor: [0, -41],
});
Leaflet.Marker.prototype.options.icon = DefaultIcon;

const vizMaker = [
  {
    from: {
      lat: 40.7449041,
      lng: -73.9886504,
    },
    to: {
      lat: 35.6812405,
      lng: 139.7649361,
    }
  },
  {
    from: {
      lat: 35.14730734477858,
      lng: 129.03368930249715,
    },
    to: {
      lat: 35.6812405,
      lng: 139.7649361,
    }
  },
  {
    from: {
      lat: 47.925458609308635,
      lng: 106.89603328136243,
    },
    to: {
      lat: 35.6812405,
      lng: 139.7649361,
    }
  },
  {
    from: {
      lat: -12.034662151727536,
      lng: -77.04546569604334,
    },
    to: {
      lat: 35.6812405,
      lng: 139.7649361,
    }
  },
  {
    from: {
      lat: 14.708033187842627,
      lng: -17.444882846679775,
    },
    to: {
      lat: 35.6812405,
      lng: 139.7649361,
    }
  },
]

const OtherMapTest = () => {
  const [zoom, setZoom] = useState(2.3);
  const [position, setPosition] = useState({
    lat: 35.6812405,
    lng: 139.7649361,
  });

  return (
    <div>
      <MapContainer center={position} zoom={zoom} style={{ width: "100%", height: "100vh" }}>
        <TileLayer
          attribution='&amp;copy <a href="http://osm.org/copyright";>OpenStreetMap</a> contributors'
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        />
        { vizMaker.map((viz, index) => {
          let from = viz.from;
          if (viz.from.lng + 180 < viz.to.lng ){
            from.lng += 360;
          }
          return (
            <>
            <Marker position={from}>
              <Popup>
                A pretty CSS3 popup. <br /> Easily customizable.
              </Popup>
            </Marker>
            <Polyline
              key={index}
              positions={[from, viz.to]}
              pathOptions={{opacity: 0}}
            />
            <GeodesicLine
              key={index}
              positions={[from, viz.to]}
              pathOptions={{color: "#000"}}
            />
            </>
          )
        })}
      </MapContainer>
    </div>
  );
};

export default OtherMapTest;