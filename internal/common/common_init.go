package common

func Init() {
	initCfg()
	initDB()
	initScript()
	initListen()
}

const (
	DefaultConfigPath = "./config/config.json"
)
