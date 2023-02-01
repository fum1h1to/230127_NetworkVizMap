package config

const (
	GEOIP_DB_PATH = "./GeoLite2-City_20230127/GeoLite2-City.mmdb"

	REPORT_FILE_NAME = "report.xml"
	OUTPUT_ROOT_DIR = "tcpflow-data"

	FROMPCAP_ROOT_DIR = OUTPUT_ROOT_DIR + "/frompcap"
	FROMPCAP_UPLOAD_DIR = FROMPCAP_ROOT_DIR + "/upload"

	REALTIME_ROOT_DIR = OUTPUT_ROOT_DIR + "/realtime"

)