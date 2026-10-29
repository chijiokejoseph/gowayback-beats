package utils

import "log"

func Must[T any](x T, err error) T {
	if err != nil {
		log.Fatalln(err)
	}
	return x
}

func Handle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
