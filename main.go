package main

import (
	"fmt"

	"github.com/duc-huy-ly/aggregator/internal/config"
)

func main() {
	fmt.Printf("Welcome to blog aggretator\n")
	myconfig := config.Read()
	fmt.Printf("%v\n%v\n", myconfig.Db_url, myconfig.Current_user_name)
	myconfig.SetUser("tom cruise")
	myconfig = config.Read()
	fmt.Printf("%v\n%v\n", myconfig.Db_url, myconfig.Current_user_name)
}
