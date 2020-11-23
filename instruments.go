package fx

import (
	"github.com/awoldes/goanda"
	"github.com/davecgh/go-spew/spew"
	"github.com/mitchellh/cli"
)

type instrumentsCommand struct {
	ui    cli.Ui
	oanda *goanda.OandaConnection
	conf  Config
}

func (c *instrumentsCommand) Run(args []string) int {
	accountInstruments := c.oanda.GetAccountInstruments(c.conf.Account)
	spew.Dump(accountInstruments)
	return 0
}
func (c *instrumentsCommand) Help() string {
	return "no help"
}
func (c *instrumentsCommand) Synopsis() string {
	return "no help"
}
