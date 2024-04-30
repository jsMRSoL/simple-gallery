package main

import (
	"flag"
	"strings"
)

func parseFlags() (debug bool, addr, port, mediaRoot string) {
	flag.BoolVar(&debug, "debug", false, "Enable debug mode")
	flag.StringVar(&addr, "addr", "", "Specify (available) IP address")
	flag.StringVar(&port, "port", "8080", "Specify (available) port")
	flag.StringVar(&mediaRoot, "dir", "~/Pictures", "Specify photo gallery")
	flag.Parse()

	if !strings.HasSuffix(mediaRoot, "/") {
		mediaRoot = mediaRoot + "/"
	}

	return debug, addr, port, mediaRoot
}
