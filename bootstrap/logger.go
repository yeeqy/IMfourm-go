package bootstrap

import (
	"IMfourm-go/pkg/config"
	"IMfourm-go/pkg/logger"
)

//存放我们的初始化logger的逻辑

func SetupLogger(){
	logger.InitLogger(
		config.GetString("log.filename"),
		config.GetInt("log.max_size"),
		config.GetInt("log.max_backup"),
		config.GetInt("log.max_age"),
		config.GetBool("log.compress"),
		config.GetString("log.type"),
		config.GetString("log.level"),
		)
}