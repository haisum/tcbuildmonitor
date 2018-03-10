package combined

import (
	"github.com/haisum/tcbuildmonitor/pkg/alert"
	"github.com/haisum/tcbuildmonitor/pkg/tc/build"
)

type combined struct {
	alerters []alert.Alerter
}

func New(alerters ...alert.Alerter) *combined {
	return &combined{alerters}
}

func (c *combined) Alert(b build.Build) error {
	for _, alerter := range c.alerters {
		err := alerter.Alert(b)
		if err != nil {
			return err
		}
	}
	return nil
}
