package main

import (
	"flag"
	"os"
	"strings"
)

func parseFlags() (debug bool, addr, port, mediaRoot string) {
	flag.BoolVar(&debug, "debug", false, "Enable debug mode")
	flag.StringVar(&addr, "addr", "127.0.0.1", "Specify (available) IP address")
	flag.StringVar(&port, "port", "8080", "Specify (available) port")
	flag.StringVar(
		&mediaRoot,
		"dir",
		os.ExpandEnv("${HOME}/Pictures/"),
		"Specify photo gallery",
	)
	flag.Parse()

	if !strings.HasSuffix(mediaRoot, "/") {
		mediaRoot = mediaRoot + "/"
	}

	return debug, addr, port, mediaRoot
}
