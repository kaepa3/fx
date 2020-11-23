package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/awoldes/goanda"
	"github.com/kaepa3/fx"
	"github.com/mitchellh/cli"
)

func main() {

	conf, err := LoadConfit("./token.toml")
	if err != nil {
		log.Println(err)
		return
	}

	key := conf.Token
	accountID := conf.Account
	oanda := goanda.NewConnection(accountID, key, false)

	c := cli.NewCLI("fx-py", "version 1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"order":       fx.OrderCommandFactory(oanda),
		"instruments": fx.InstrumentsCommandFactory(oanda),
	}

	exitCode, err := c.Run()
	if err != nil {
		log.Printf("Error executing CLI: %s", err.Error())
	}

	os.Exit(exitCode)
}

type Config struct {
	Token   string
	Account string
}

func LoadConfit(path string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
