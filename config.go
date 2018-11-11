package main

import (
	"os"
)

type Configuration struct {
	Address string
}

func configurationFromEnv() Configuration {
	c := Configuration{}
	c.Address = os.Getenv("SUBSCRIBER_ADDRESS")
	return c
}

var Config = configurationFromEnv()
