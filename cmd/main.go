package main

import (
	"log"
	"os"

	"github.com/awoldes/goanda"
	"github.com/kaepa3/fx"
	"github.com/mitchellh/cli"
)

func main() {

	conf, err := fx.LoadConfig("./token.toml")
	if err != nil {
		log.Println(err)
		return
	}

	key := conf.Token
	accountID := conf.Account
	oanda := goanda.NewConnection(accountID, key, false)

	c := cli.NewCLI("fx", "version 1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"ORDER": func() (cli.Command, error) {
			return fx.OrderCommandFactory(*conf, oanda)
		},
		"INST": func() (cli.Command, error) {
			return fx.InstrumentsCommandFactory(*conf, oanda)
		},
		"CANDLE": func() (cli.Command, error) {
			return fx.CandleCommandFactory(*conf, oanda)
		},
	}

	exitCode, err := c.Run()
	if err != nil {
		log.Printf("Error executing CLI: %s", err.Error())
	}

	os.Exit(exitCode)
}
