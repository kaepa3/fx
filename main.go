package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/awoldes/goanda"
	"github.com/davecgh/go-spew/spew"
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

	if len(os.Args) <= 1 {
		accountInstruments := oanda.GetAccountInstruments(accountID)
		spew.Dump(accountInstruments)
		return
	}
	cmd := os.Args[1]
	if cmd == "order" {
		Order(oanda, conf, 10)
	} else {
		fmt.Println("no cmd")
	}
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

func Order(oanda *goanda.OandaConnection, conf *Config, yen int) {
	order := goanda.OrderPayload{
		Order: goanda.OrderBody{
			Units:        10000,
			Instrument:   "USD_JPY",
			TimeInForce:  "GTC",
			Type:         "MARKET",
			PositionFill: "DEFAULT",
			Price:        "1.25000",
		},
	}
	orderResult := oanda.CreateOrder(order)
	spew.Dump("%+v\n", orderResult)
}
func Sell(dolers int) {

}
