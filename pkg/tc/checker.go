package tc

import (
	"github.com/haisum/tcbuildmonitor/pkg/alert"
	"github.com/haisum/tcbuildmonitor/pkg/tc/build"
)

const (
	FAILURE = "FAILURE"
	SUCCESS = "SUCCESS"
)

type TcAPI interface {
	Get(buildTypeID string) ([]build.Build, error)
}

func Check(tc TcAPI, alerter alert.Alerter, buildTypesIDs ...string) error {
	for _, buildTypeID := range buildTypesIDs {
		builds, err := tc.Get(buildTypeID)
		if err != nil {
			return err
		}
		if len(builds) > 0 {
			// builds[0].Status == FAILURE && builds[1].Status == SUCCESS) || (builds[0].TestOccurrences.NewFailed > 0
			if builds[0].Status == FAILURE {
				err = alerter.Alert(builds[0])
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
