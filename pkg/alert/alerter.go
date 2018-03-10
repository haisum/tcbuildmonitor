package alert

import (
	"github.com/haisum/tcbuildmonitor/pkg/tc/build"
)

type Alerter interface {
	Alert(build.Build) error
}
