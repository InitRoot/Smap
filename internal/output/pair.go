package output

import (
	"fmt"
	"os"

	g "github.com/s0md3v/smap/internal/global"
)

var openedPairFile *os.File

func StartPair() {
	if g.PairFilename != "-" {
		openedPairFile = OpenFile(g.PairFilename)
	}
}

func ContinuePair(result g.Output) {
	thisString := ""
	for _, port := range result.Ports {
		thisString += fmt.Sprintf("%s:%d,%s,%s\n", result.IP, port.Port, port.Service, port.Version)
	}
	Write(thisString, g.PairFilename, openedPairFile)
}

func EndPair() {
}
