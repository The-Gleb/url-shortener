package config

import (
	"flag"
	// "os"
)

type Config struct {
	ServerAddress string
	BaseAddress   string
}

type ConfigBuilder struct {
	config Config
}

func (b ConfigBuilder) SetServerAddres(address string) ConfigBuilder {
	b.config.ServerAddress = address
	return b
}

func (b ConfigBuilder) SetBaseAddress(address string) ConfigBuilder {
	b.config.BaseAddress = address
	return b
}

func NewConfigFromFlags() Config {
	// flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	var serverAddress string
	flag.StringVar(&serverAddress, "a", ":8080", "address and port to run server")

	var baseAddress string
	flag.StringVar(&baseAddress, "b", "http://localhost:8080/", "address before shortened url")

	flag.Parse()

	var builder ConfigBuilder

	builder = builder.SetServerAddres(serverAddress).
		SetBaseAddress(baseAddress)

	// if envServerAddress := os.Getenv("ADDRESS"); envServerAddress != "" {
	// 	builder = builder.SetServerAddres(envServerAddress)
	// }
	// if envBaseAddress := os.Getenv("POLL_INTERVAL"); envBaseAddress != "" {
	// 	builder = builder.SetBaseAddress(envBaseAddress)
	// }

	return builder.config
}
