package util

import "log"

func LogNonFatal(err error) {
	if err != nil {
		log.Println(err)
	}
}
