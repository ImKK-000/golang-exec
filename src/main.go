package main

import (
	"fmt"
	"golang-exec/loadconfig"
)

func main() {
	configs := loadconfig.LoadConfig()
	for index, config := range configs {
		fmt.Println(index, config.Name, config.ExecuteCommand)
	}
}
