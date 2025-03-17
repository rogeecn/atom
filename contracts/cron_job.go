package contracts

import (
	"github.com/riverqueue/river"
)

type CronJob interface {
	Args() []CronJobArg
}

type CronJobArgInterface interface {
	river.JobArgs
	river.JobArgsWithInsertOpts
}

type CronJobArg struct {
	RunOnStart       bool
	PeriodicInterval river.PeriodicSchedule

	Arg CronJobArgInterface
}
