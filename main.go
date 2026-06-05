package main

import (
	"fmt"

	"github.com/duc-huy-ly/aggregator/internal/config"
)

func main() {
	myconfig := config.Read()
	fmt.Printf("%v\n", myconfig.Db_url)
}

