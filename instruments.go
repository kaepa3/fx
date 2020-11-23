package fx

import (
	"github.com/davecgh/go-spew/spew"
)

type instrumentsCommand struct {
	ui    cli.ui
	oanda *goanda.oandaconnection
}

func (c *instrumentsCommand) run(args []string) {
	accountInstruments := oanda.GetAccountInstruments(accountID)
	spew.Dump(accountInstruments)
}
