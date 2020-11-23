package fx

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

func OrderCommandFactory(c Config, o *goanda.OandaConnection) (cli.Command, error) {
	return &orderCommand{
		ui:    Ui,
		oanda: o,
		conf:  c,
	}, nil
}
func CandleCommandFactory(c Config, o *goanda.OandaConnection) (cli.Command, error) {
	return &candleCommand{
		ui:    Ui,
		oanda: o,
		conf:  c,
	}, nil
}
func InstrumentsCommandFactory(c Config, o *goanda.OandaConnection) (cli.Command, error) {
	return &instrumentsCommand{
		ui:    Ui,
		oanda: o,
		conf:  c,
	}, nil
}
