package cmd

import (
	"fmt"

	"github.com/prometheus/procfs"
	"github.com/spf13/cobra"
)

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Show CPU information and statistics for the host",
	RunE:  describeCPU,
}

func describeCPU(cmd *cobra.Command, args []string) error {
	fs, err := procfs.NewFS("/proc")
	if err != nil {
		return fmt.Errorf("failed to initalize the proc filesystem mount point. Aborting")
	}

	c, err := fs.CPUInfo()
	if err != nil {
		return fmt.Errorf("failed to read cpu information. Aborting")
	}

	fmt.Printf("CPU Model: %s\n", c[0].ModelName)
	fmt.Printf("Core Count: %d\n", len(c))
	// for _, cpu := range c {
	// 	fmt.Println(cpu.ModelName)
	// }

	return nil
}
