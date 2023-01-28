package config

import "atom/utils"

type App struct {
	Mode string
}

func (a App) IsDebug() bool {
	return a.Mode == "debug"
}

func (a App) IsRelease() bool {
	return a.Mode == "release"
}
func (a App) IsTesting() bool {
	return a.Mode == "testing" || utils.IsInTesting()
}
