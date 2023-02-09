package pack

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func PrintCurrentTime(host string) {
	time, err := ntp.Time(host)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(time)
}
