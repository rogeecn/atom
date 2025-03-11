package contracts

import (
	"github.com/riverqueue/river"
)

type CronJob interface {
	Kind() string
	Periodic() river.PeriodicSchedule
	JobArgs() river.JobArgs
	InsertOpts() river.InsertOpts
	RunOnStart() bool
}
