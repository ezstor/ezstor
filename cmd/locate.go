package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newLocateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "locate",
		Short: "",
		Long:  "",
	}

	cmd.AddCommand(newLocateOnCmd())
	cmd.AddCommand(newLocateOffCmd())

	return cmd
}

func newLocateOnCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "on",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Locate On")
		},
	}

	return cmd
}

func newLocateOffCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "off",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Locate Off")
		},
	}

	return cmd
}
