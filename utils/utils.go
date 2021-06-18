package utils

import "errors"

func PanicWithError(err error, msg string) {
	if err != nil {
		panic(errors.New(msg))
	}
}
