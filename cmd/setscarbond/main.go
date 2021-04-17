package main

import (
	"os"

	"github.com/MasterPi-2124/setscarbon/app"
	"github.com/MasterPi-2124/setscarbon/cmd/setscarbond/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
