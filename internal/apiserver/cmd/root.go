package cmd

import (
	"github.com/spf13/cobra"

	"github.com/luohu1/gin-apiserver/internal/apiserver"
	"github.com/luohu1/gin-apiserver/internal/apiserver/cmd/version"
)

func NewCmdAPIServer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apiserver",
		Short: "Gin APIServer",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := apiserver.RunServer()
			return err
		},
	}

	cmd.AddCommand(version.NewCmdVersion())

	return cmd
}
