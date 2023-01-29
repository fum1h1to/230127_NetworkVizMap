import { createElementObject, createPathComponent, extendContext } from '@react-leaflet/core';
import { GeodesicLine as LeafletGeodesicLine } from 'leaflet.geodesic'
export const GeodesicLine = createPathComponent(function createGeodesicLine({ positions , ...options }, ctx) {
    const GeodesicLine = new LeafletGeodesicLine(positions, { wrap: false, ...options});
    return createElementObject(GeodesicLine, extendContext(ctx, {
        overlayContainer: GeodesicLine
    }));
}, function updateGeodesicLine(layer, props, prevProps) {
    if (props.positions !== prevProps.positions) {
        layer.setLatLngs(props.positions);
    }
});
