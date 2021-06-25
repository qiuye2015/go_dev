package config

const (
	ItemSaverPort = 1234
	WorkerPort0   = 9000

	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"
	ElasticIndex    = "dating_profile"

	// Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"

	Qps = 1
)
