package utils

import "fmt"

func Green(text string) {
	fmt.Printf("\x1b[%dm"+fmt.Sprint(text)+"  \x1b[0m\n", 32)
}

func FillGray(text interface{}) {
	fmt.Printf("\x1b[%d;%dmtime consuming: \x1b[0m "+fmt.Sprint(text)+" \n", 47, 30)
}

func Red(text interface{}) {
	fmt.Printf("\x1b[%dm"+fmt.Sprint(text)+" \x1b[0m\n", 31)
}
