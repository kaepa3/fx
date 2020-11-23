package fy

import (
	"os"

	"github.com/awoldes/goanda"
	"github.com/mitchellh/cli"
)

var Ui cli.Ui

const (
	infoPrefix  = "INFO: "
	warnPrefix  = "WARN: "
	errorPrefix = "ERROR: "
)

func init() {
	Ui = &cli.PrefixedUi{
		InfoPrefix:  infoPrefix,
		WarnPrefix:  warnPrefix,
		ErrorPrefix: errorPrefix,
		Ui: &cli.BasicUi{
			Writer: os.Stdout,
		},
	}
}

func OrderCommandFactory(oanda *goanda.OandaConnection) (cli.Command, error) {
	return &orderCommand{
		ui: Ui,
	}, nil
}
func InstrumentsCommandFactory(oanda *goanda.OandaConnection) (cli.Command, error) {
	return &orderCommand{
		ui: Ui,
	}, nil
}
