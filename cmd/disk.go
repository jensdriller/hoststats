package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/prometheus/procfs"
	"github.com/spf13/cobra"
)

const diskstatsFilename = "diskstats"

var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Show Disk information and statistics for the host",
	RunE:  describeDisk,
}

var procPath = procfs.DefaultMountPoint

func describeDisk(cmd *cobra.Command, args []string) error {
	invalid, _ := regexp.Compile("^(ram|loop|fd|(h|s|v|xv)d[a-z]|nvme\\d+n\\d+p)\\d+$")

	data, err := getDiskStats()
	if err != nil {
		return err
	}

	for dev, _ := range data {
		if invalid.MatchString(dev) {
			continue
		}
		fmt.Printf("Device: %s\n", dev)
		// fmt.Println(info)
		fmt.Println()
	}

	return nil
}

func procFilePath(name string) string {
	return filepath.Join(procPath, name)
}

func getDiskStats() (map[string][]string, error) {
	file, err := os.Open(procFilePath(diskstatsFilename))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return parseDiskStats(file)
}

func parseDiskStats(r io.Reader) (map[string][]string, error) {
	var (
		diskStats = map[string][]string{}
		scanner   = bufio.NewScanner(r)
	)

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) < 4 { // we strip major, minor and dev
			return nil, fmt.Errorf("invalid line in %s: %s", procFilePath(diskstatsFilename), scanner.Text())
		}
		dev := parts[2]
		diskStats[dev] = parts[3:]
	}

	return diskStats, scanner.Err()
}
