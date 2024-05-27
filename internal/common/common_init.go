package common

const (
	DefaultConfigPath = "./config/config.json"
)

func Init() {
	initCfg()
	initDB()
	initScript()
	initListen()

}
