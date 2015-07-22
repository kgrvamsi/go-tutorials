package main

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Specification struct {
	Debug bool
	Port  int
	User  string
	Rate  float32
	Go    string
}

type JavaSpec struct {
	Version string
	Path    string
}

func main() {
	var s Specification
	err := envconfig.Process("myapp", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	format := "Debug: %v\nPort: %d\nUser: %s\nRate: %f\nGo: %s\n"
	_, err = fmt.Printf(format, s.Debug, s.Port, s.User, s.Rate, s.Go)
	if err != nil {
		log.Fatal(err.Error())
	}

	var j JavaSpec
	errr := envconfig.Process("java", &j)
	if errr != nil {
		log.Fatal(errr.Error())
	}

	_, er := fmt.Printf("Java Version: %s\nJava Path: %s\n", j.Version, j.Path)
	if er != nil {
		log.Fatal(er.Error())
	}

}
