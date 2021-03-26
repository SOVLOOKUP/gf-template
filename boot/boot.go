package boot

import (
	_ "gf-app/packed"
)

func init() {
	InitConfig()
	InitModules()
}
