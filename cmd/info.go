package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Info")
		},
	}

	return cmd
}
