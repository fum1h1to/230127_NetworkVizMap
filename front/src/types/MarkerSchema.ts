type MarkerSchema = {
    from: { lat: number; lng: number; };
    to: { lat: number; lng: number; };
    srcport: Number;
    dstport: Number;
}

export default MarkerSchema;