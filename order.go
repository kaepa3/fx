package fx

import (
	"github.com/awoldes/goanda"
	"github.com/davecgh/go-spew/spew"
	"github.com/mitchellh/cli"
)

type orderCommand struct {
	ui    cli.Ui
	oanda *goanda.OandaConnection
	conf  Config
}

func (c *orderCommand) Run(args []string) int {
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
	orderresult := c.oanda.CreateOrder(order)
	spew.Dump("%+v\n", orderresult)
	return 0
}

func (c *orderCommand) Help() string {
	return "no help"
}

func (c *orderCommand) Synopsis() string {
	return "no help"
}
