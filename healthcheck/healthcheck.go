package healthcheck

import (
	"../common"
	"../config"
)

type HealthChecker interface {
	Run()
}
