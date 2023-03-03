package config

//Database connection's config
var (
	Name         = "root"
	Password     = "sx221410"
	DNS          = "localhost"
	DatabasePort = 701
	DatabaseName = "minidouyin"
)

//redis config
var (
	RdbNetwork     = "tcp"
	RdbAddress     = "121.199.43.175:6379"
	RdbUseDatabase = 0 //select use which one database
	RdbPassword    = "sx221410"
)

// JwtKey JWT secret Key
var JwtKey = []byte("tyhngebvfpliyergfgdf")

//Qi Niu cloud config file
var (
	AccessKey     = "XuigBGSCJ7vpAtRtpu04NqLGLXpEROCaqgOxTZ0W"
	SecretKey     = "mhV_z93CyJCcDTmSfU2cSfx_LiejWCjujCCRMuqg"
	VideoBucket   = "minidouyin-video"
	PictureBucket = "minidouyin-picture"
	DomainVideo   = "http://rq9lt9dry.bkt.clouddn.com"
	DomainPicture = "http://rq9lfs4ld.bkt.clouddn.com"
)

//n:a CRUD process can get n video
const N = 10
