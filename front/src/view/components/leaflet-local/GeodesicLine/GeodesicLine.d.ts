import { type PathProps } from '@react-leaflet/core';
import { type LatLngExpression, GeodesicLine as LeafletGeodesicLine, type GeodesicOptions } from 'leaflet.geodesic';
import { type PathOptions } from 'leaflet';
import type { ReactNode } from 'react';

export interface GeodesicLineProps extends PathProps {
    children?: ReactNode;
    positions: LatLngExpression[] | LatLngExpression[][];
}
export declare const GeodesicLine: import("react").ForwardRefExoticComponent<GeodesicOptions & import("react").RefAttributes<LeafletGeodesicLine<import("geojson").LineString | import("geojson").MultiLineString, any>>>;
