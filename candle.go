package fx

import (
	"github.com/awoldes/goanda"
	"github.com/davecgh/go-spew/spew"
	"github.com/mitchellh/cli"
)

type candleCommand struct {
	ui    cli.Ui
	oanda *goanda.OandaConnection
	conf  Config
}

func (c *candleCommand) Run(args []string) int {
	oanda := goanda.NewConnection(c.conf.Account, c.conf.Token, false)
	history := oanda.GetCandles("EUR_USD", "10", "S5")
	spew.Dump(history)
	return 0
}

func (c *candleCommand) Help() string {
	return "no help"
}

func (c *candleCommand) Synopsis() string {
	return "no help"
}
