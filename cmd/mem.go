package cmd

import (
	"fmt"

	"github.com/prometheus/procfs"
	"github.com/spf13/cobra"
)

var memoryCmd = &cobra.Command{
	Use:   "memory",
	Short: "Show Memory information and statistics for the host",
	RunE:  describeMemory,
}

func describeMemory(cmd *cobra.Command, args []string) error {
	fs, err := procfs.NewFS("/proc")
	if err != nil {
		return fmt.Errorf("failed to initalize the proc filesystem mount point. Aborting")
	}

	m, err := fs.Meminfo()
	if err != nil {
		return fmt.Errorf("failed to read mem stat. Aborting")
	}
	fmt.Printf("Total: %d\n", m.MemTotal)
	fmt.Printf("Available: %d\n", m.MemAvailable)
	fmt.Printf("Free: %d\n", m.MemFree)

	return nil
}
