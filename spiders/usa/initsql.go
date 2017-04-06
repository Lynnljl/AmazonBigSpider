package main

import (
	"github.com/hunterhug/AmazonBigSpider"
	"github.com/hunterhug/AmazonBigSpider/public/core"
)

func main() {
	if dudu.Local {
		core.InitConfig(dudu.Dir+"/config/usa_local_config.json", dudu.Dir+"/config/usa_log.json")
	} else {
		core.InitConfig(dudu.Dir+"/config/usa_config.json", dudu.Dir+"/config/usa_log.json")
	}
	core.InitDB()
}
