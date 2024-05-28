package helpers

import "log"

func Trace(s string) string {
	log.Println(s)
	return s
}

func Untrace(s string) string {
	log.Println("leaving", s)
	return s
}
