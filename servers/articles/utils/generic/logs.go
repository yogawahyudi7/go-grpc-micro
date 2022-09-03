package utils

import (
	"bytes"
	"log"
)

var (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

var (
	Buf           bytes.Buffer
	InfoLogger    = log.New(&Buf, Green+"INFO: "+Green, log.Llongfile)
	ErrorLogger   = log.New(&Buf, Yellow+"ERROR: "+Yellow, log.Llongfile)
	WarningLogger = log.New(&Buf, Red+"WARNING: "+Red, log.Llongfile)
	Infof         = func(info string) {
		Buf.Reset()
		InfoLogger.Output(2, info+Reset)
	}
	Errorf = func(err string) {
		Buf.Reset()
		ErrorLogger.Output(2, err+Reset)
	}
	WarningF = func(warn string) {
		Buf.Reset()
		WarningLogger.Output(2, warn+Reset)
	}
)
