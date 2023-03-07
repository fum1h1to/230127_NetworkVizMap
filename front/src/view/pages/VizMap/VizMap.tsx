import React, { useState } from "react";
import "leaflet/dist/leaflet.css";
import { MapContainer, TileLayer, Marker, Popup, Polyline } from "react-leaflet";
// add icons
import Leaflet from "leaflet";
import icon from "leaflet/dist/images/marker-icon.png";
import iconShadow from "leaflet/dist/images/marker-shadow.png";
import { GeodesicLine } from "view/components/leaflet-local/GeodesicLine/GeodesicLine";
import MarkerSchema from "types/MarkerSchema";

// marker setting
let DefaultIcon = Leaflet.icon({
  iconUrl: icon,
  shadowUrl: iconShadow,
  iconAnchor: [12.5, 41],
  popupAnchor: [0, -41],
});
Leaflet.Marker.prototype.options.icon = DefaultIcon;

const VizMap = (props: {
  fromORto: "from" | "to" | "all",
  myipv4: string | undefined,
  myipv6: string | undefined,
  center: { lat: number, lng: number },
  markers: MarkerSchema[],
}) => {
  const [zoom, setZoom] = useState(2.3);
  const local_markers = JSON.parse(JSON.stringify(props.markers)) as MarkerSchema[];

  local_markers.forEach((viz) => {

    if (viz.srcip === props.myipv4 || viz.srcip === props.myipv6) {
      viz.from.lat = props.center.lat;
      viz.from.lng = props.center.lng;
    }
    if (viz.dstip === props.myipv4 || viz.dstip === props.myipv6) {
      viz.to.lat = props.center.lat;
      viz.to.lng = props.center.lng;
    }

    if (props.fromORto === "to") {
      if (viz.to.lng + 180 < viz.from.lng) {
        viz.to.lng += 360;
      }
    } else {
      if (viz.from.lng + 180 < viz.to.lng) {
        viz.from.lng += 360;
      }
    }

  })

  return (
    <MapContainer center={props.center} zoom={zoom} style={{ width: "100%", height: "100%" }}>
      <TileLayer
        attribution='&amp;copy <a href="http://osm.org/copyright";>OpenStreetMap</a> contributors'
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
      />
      {
        props.fromORto !== "all"
          ?
          <Marker position={props.center}>
            <Popup>
              me
            </Popup>
          </Marker>
          :
          <></>
      }
      {local_markers.filter((viz) => {
        if (props.fromORto === "from" && (props.myipv4 !== "" || props.myipv4 !== undefined || props.myipv6 !== "" || props.myipv6 !== undefined)) {
          if (viz.srcip === props.myipv4 || viz.srcip === props.myipv6) {
            return false
          }
        }
        else if (props.fromORto === "to" && (props.myipv4 !== "" || props.myipv4 !== undefined || props.myipv6 !== "" || props.myipv6 !== undefined)) {
          if (viz.dstip === props.myipv4 || viz.dstip === props.myipv6) {
            return false
          }
        }

        if ((viz.from.lng === 0 && viz.from.lat === 0) || (viz.to.lng === 0 && viz.to.lat === 0)) {
          return false
        }
        return true
      }).map((viz, index) => {
        return (
          <div key={index}>
            <Marker position={props.fromORto === "to" ? viz.to : viz.from}>
              <Popup>
                index: {String(index)}<br />
                lat,lng: {String(viz.from.lat)}, {String(viz.from.lng)}<br />
                srcip: {String(viz.srcip)}<br />
                dstip: {String(viz.dstip)}<br />
                srcport: {String(viz.srcport)}<br />
                dstport: {String(viz.dstport)}
              </Popup>
            </Marker>
            <GeodesicLine
              positions={props.fromORto === "to" ? [viz.to, viz.from] : [viz.from, viz.to]}
              pathOptions={{ color: "#000" }}
            />
          </div>
        )
      })}
    </MapContainer>
  );
};

export default VizMap;