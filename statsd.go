package main

import (
	"fmt"
	"github.com/cactus/go-statsd-client/statsd"
	"github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/process"
	_ "io/ioutil"
	_ "math/rand"
	"os"
	_ "os/exec"
	_ "strconv"
	"time"
)

func hostname() string {
	host, _ := os.Hostname()
	return host
}

func apachePID() (pid []int) {
	k, _ := ps.Processes()

	for _, exe := range k {

		if exe.Executable() == "apache2" {

			pid = append(pid, exe.Pid())

		}
	}
	return pid
}

func memoryUtil() (value int64, rate float32) {

	pid := apachePID()

	total := make([]int64, len(pid))

	var sum int64

	for i := 0; i < len(pid); i++ {

		k, _ := process.NewProcess(int32(pid[i]))
		rss, _ := k.MemoryInfo()

		val := rss.RSS / 1024

		total[i] = int64(val)

	}

	for k := 0; k < len(total); k++ {

		sum = sum + total[k]

	}

	rat := (sum / 1024) * 100

	fmt.Println(sum, rat)

	return sum / 1024, float32(rat)
}

func main() {

	for {

		client, err := statsd.NewClient("104.130.20.4:8125", hostname())
		if err != nil {
			fmt.Println(err.Error())
		}

		memoryUtil()

		defer client.Close()

		value, _ := memoryUtil()

		fmt.Println(value)

		client.Gauge("Apache2", value, 0.4)

		fmt.Println(hostname())

		time.Sleep(1 * time.Second)

	}
}
