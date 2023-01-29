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

const VizMap = (props: {
  markers: { from: { lat: number; lng: number; }; to: { lat: number; lng: number; }; }[],
}) => {
  const [zoom, setZoom] = useState(2.3);
  const [position, setPosition] = useState({
    lat: 35.6812405,
    lng: 139.7649361,
  });

  return (
    <MapContainer center={position} zoom={zoom} style={{ width: "100%", height: "100%" }}>
      <TileLayer
        attribution='&amp;copy <a href="http://osm.org/copyright";>OpenStreetMap</a> contributors'
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
      />
      { props.markers.map((viz, index) => {
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
  );
};

export default VizMap;