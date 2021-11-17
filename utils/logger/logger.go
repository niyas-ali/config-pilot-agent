package logger

import (
	"log"

	"github.com/logrusorgru/aurora"
)

var (
	success = aurora.Green("[✓]").String()
	failed  = aurora.Red("[✘]").String()
)

func Println(v ...interface{}) {
	log.Println(v...)
}
func PrintSuccessln(v ...interface{}) {
	log.Println(success, v)
}
func PrintErrorln(v ...interface{}) {
	log.Println(failed, v)
}
func Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
func Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}
func Fatalln(v ...interface{}) {
	log.Fatalln(v...)
}
