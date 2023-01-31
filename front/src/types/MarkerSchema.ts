type MarkerSchema = {
    from: { lat: number; lng: number; };
    to: { lat: number; lng: number; };
    srcip: String;
    dstip: String;
    srcport: Number;
    dstport: Number;
}

export default MarkerSchema;