package util

import "log"

func LogNonFatal(err error) {
	if err != nil {
		log.Println(err)
	}
}

func Try0(err error) {
	if err != nil {
		log.Panic(err)
	}
}
