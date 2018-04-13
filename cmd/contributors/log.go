package main

import "log"

func debugf(format string, v ...interface{}) {
	if *flagDebug {
		log.Printf(format, v...)
	}
}

func warnf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
