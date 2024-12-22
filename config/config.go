package config

import "flag"

var (
	Port string
)

func InitConfigs() {
	portFlag := flag.String("port", "8080", "Port number")
	flag.Parse()

	if *portFlag != "" {
		Port = ":" + *portFlag
	}
}
