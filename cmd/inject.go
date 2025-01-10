package cmd

import (
	"errors"
	"fmt"
	"injection/internal/faults"

	"github.com/spf13/cobra"
)

var vmName string
var faultType string
var duration int

var injectCmd = &cobra.Command{
	Use:   "inject",
	Short: "Inject a fault into a VM",
	Long:  "Inject a specified fault (CPU spike or memory leak) into a VM for a given duration.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vmName == "" || faultType == "" || duration <= 0 {
			return errors.New("please specify the targetted --vm, --type, and --duration (in seconds)")
		}

		switch faultType {
		case "cpu":
			return faults.InjectCPU(vmName, duration)
		default:
			return fmt.Errorf("unknown fault type: %s", faultType)
		}
	},
}

func init() {
	injectCmd.Flags().StringVar(&vmName, "vm", "", "Name of the VM to target (In progress. For now, any name will suffice)")
	injectCmd.Flags().StringVar(&faultType, "type", "", "Type of fault to inject (cpu)")
	injectCmd.Flags().IntVar(&duration, "duration", 0, "duration of fault (seconds)")
}
