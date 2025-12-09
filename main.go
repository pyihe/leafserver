package main

import (
	"github.com/pyihe/leaf"
	lconf "github.com/pyihe/leaf/conf"

	"github.com/pyihe/leafsearver/conf"
	"github.com/pyihe/leafsearver/game"
	"github.com/pyihe/leafsearver/gate"
	"github.com/pyihe/leafsearver/login"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
