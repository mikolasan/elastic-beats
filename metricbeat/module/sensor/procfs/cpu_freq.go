package procfs

import (
    "bufio"
    "fmt"
    "os"
    "strings"
	"strconv"

	// "github.com/pkg/errors"
)

func readCurFreq() (float64, error) {
    file, err := os.Open("/proc/cpuinfo")
    if err != nil {
        fmt.Println("Failed to open /proc/cpuinfo")
        return 0.0, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, "cpu MHz") {
            parts := strings.Split(line, ":")
            freq := strings.TrimSpace(parts[1])
			f, err := strconv.ParseFloat(freq, 64)
            // fmt.Printf("CPU frequency: %s MHz\n", freq)
            return f, err
        }
    }
    fmt.Printf("Failed to find CPU frequency in /proc/cpuinfo")
	return 0.0, err
}
