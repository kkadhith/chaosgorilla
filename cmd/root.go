package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fault",
	Short: "Inject faults in your VM",
	Long:  "ChaosGorilla is a fault injection CLI tool to introduce CPU faults into VMs.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(injectCmd)
}
