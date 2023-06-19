package atom

import (
	"github.com/rogeecn/atom-addons/services/queue"
	"github.com/spf13/cobra"
)

func withQueueCommand(rootCmd *cobra.Command) *cobra.Command {
	queueCmd := &cobra.Command{
		Use:   "queue",
		Short: "queue consumer",
		RunE: func(cmd *cobra.Command, args []string) error {
			return queue.Serve()
		},
	}
	rootCmd.AddCommand(queueCmd)
	return rootCmd
}
