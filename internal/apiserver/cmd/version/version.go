package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the client version information",
		Long:  "Print the client and server version information for the current context",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("v0.1.0")
		},
	}
}
