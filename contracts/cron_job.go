package contracts

import (
	"github.com/riverqueue/river"
)

type CronJob interface {
	Args() []CronJobArg
}

type JobArgs interface {
	river.JobArgs
	river.JobArgsWithInsertOpts
}

type CronJobArg struct {
	RunOnStart       bool
	PeriodicInterval river.PeriodicSchedule

	Arg JobArgs
}
