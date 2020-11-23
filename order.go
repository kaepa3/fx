package fx

import (
	"github.com/awoldes/goanda"
	"github.com/davecgh/go-spew/spew"
	"github.com/mitchellh/cli"
)

type ordercommand struct {
	ui    cli.ui
	oanda *goanda.oandaconnection
}

func (c *ordercommand) run(args []string) {
	order := goanda.orderpayload{
		order: goanda.orderbody{
			units:        10000,
			instrument:   "usd_jpy",
			timeinforce:  "gtc",
			type:         "market",
			positionfill: "default",
			price:        "1.25000",
		},
	}
	orderresult := oanda.createorder(order)
	spew.dump("%+v\n", orderresult)
}
