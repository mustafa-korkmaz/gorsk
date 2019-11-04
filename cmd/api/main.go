package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/mustafa-korkmaz/movies/pkg/api"
	"github.com/mustafa-korkmaz/movies/pkg/utl/config"
)

func main() {

	fmt.Println(getBinPath())

	cfgPath := flag.String("p", "./cmd/api/conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)

	checkErr(api.Start(cfg))
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func getBinPath() string {
	e, err := os.Executable()
	if err != nil {
		panic(err)
	}
	path := path.Dir(e)
	return path
}
