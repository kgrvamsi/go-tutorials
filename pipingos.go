package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type ProcesPids struct {
	Process []struct {
		Pid int `json:"pid"`
	} `json:"process"`
}

func getPids() ([5]string, error) {

	// pids := &ProcesPids{}
	cmd := "ps aux | awk '{print $2, $4, $11}' | sort -k2rn | head -n 5|awk '{print $1}'"

	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}

	f, _ := os.Create("pids.txt")

	f.Write(out)

	inFile, _ := os.Open("pids.txt")

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	var data [5]string
	for scanner.Scan() {

		// for i := 0; i < 4; i++ {

		// if data[i] == scanner.Text() {

		// 	fmt.Println(scanner.Text())

		// 	data[i] = scanner.Text()

		// } else {
		// 	fmt.Println("Duplicate entry")
		// 	data[i] = scanner.Text()
		// }

		b := scanner.Text()

		fmt.Println(b)

		for i := 0; i < 5; i++ {

			data[i] = b
		}

		// }

		// fmt.Println(scanner.Text(), "We are Here")

		// data[0] = scanner.Text()
		// data[1] = scanner.Text()
		// data[2] = scanner.Text()
		// data[3] = scanner.Text()
		// data[4] = scanner.Text()

	}
	return data, nil

}

func main() {
	c1 := exec.Command("ls")
	c2 := exec.Command("wc", "-l")

	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r

	var b2 bytes.Buffer
	c2.Stdout = &b2

	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()
	io.Copy(os.Stdout, &b2)

	k, _ := getPids()

	fmt.Println(k, "We are Here")

	// scanner := bufio.Reader(pid)

	// fmt.Println(scanner)

}
