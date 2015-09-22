package main

import (
	slog "sphire/log"
)

func main() {
	slog.Init("json", "DEV", "/tmp/sphire.log")
	slog.Log(nil, "Error connecting to database", "debug")
}

