package main

import (
	"archive/tar"
	_ "compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	flag.Parse()

	tarname := flag.Arg(0)

	if tarname == "" {

		fmt.Println("Please give some value")
	}

	foldername := flag.Arg(1)

	if foldername == "" {
		fmt.Println("Please give some value")
	}

	folcheck, _ := os.Open(foldername)

	folread, _ := folcheck.Readdir(0)

	fmt.Println(folread)

	fmt.Println(folcheck)
	fmt.Println(tarname)

	if strings.Index(tarname, ".tar") != -1 {
		fmt.Println("The Tar name is Testing")
	} else {
		fmt.Println("Sorry the file name is not Testing")
	}

	tarfile, err := os.Create(tarname)

	if err != nil {
		panic(err)
	}

	var fileWriter io.WriteCloser = tarfile

	tarfileWriter := tar.NewWriter(fileWriter)
	defer tarfileWriter.Close()

	for _, read := range folread {

		fmt.Println(read)

		if read.IsDir() {
			continue
		}

	}
}
