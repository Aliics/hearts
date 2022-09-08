package main

import "log"

func logNonFatal(err error) {
	if err != nil {
		log.Println(err)
	}
}
