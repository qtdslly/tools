package constant

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB

	RootPath        = "/home/ymminer/"
	MinerPath       = "/home/ymminer/miners/"
	MinerLog        = "/home/ymminer/logs/"
	CommandPath     = "/home/ymminer/bin/"
	CommandCorePath = "/home/ymminer/core/command/"
	TmpPath         = "/home/ymminer/tmp/"
	ConfigPath      = "/home/ymminer/etc/"

	KernelUrl       = "http://ymos-bucket.oss-cn-beijing.aliyuncs.com/v2/kernels/"
	DownloadBaseUrl = "http://ymos-bucket.oss-cn-beijing.aliyuncs.com/v2/"
	DBName          = "mysql"
	//DBSource = "mqtt:Qq61299400@tcp(rm-2ze8n5q1w2x07qv81zo.mysql.rds.aliyuncs.com:3306)/ymos_2?charset=utf8&parseTime=True&loc=Local"
	//
	//RedisAddr = "r-2zek2bnndze5dvhod1pd.redis.rds.aliyuncs.com:6379"
	//RedisPass = "QAZwsx!@#456"

	Success = 200
	Error   = 500

	EncryptKey = "0g.BSDYmMiner*(6"

	OsReleaseFile = "/etc/os-release"
)
