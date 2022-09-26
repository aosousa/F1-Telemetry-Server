package logger

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

const (
	DEBUG = "Debug"
	ERROR = "Error"
)

func getCallerInformation() (string, string, string, int) {
	pc, file, line, _ := runtime.Caller(2)

	_, fileName := path.Split(file)
	funcForPC := runtime.FuncForPC(pc).Name()
	funcForPCSplit := strings.Split(funcForPC, ".")
	packageName := funcForPCSplit[0]
	funcName := funcForPCSplit[1]

	return packageName, funcName, fileName, line
}

func logToFile(logType string, message string) {
	var d, y, m, folderPath, filePath string

	packageName, funcName, fileName, line := getCallerInformation()

	currentTime := time.Now()
	d, y, m = currentTime.Format("02"), currentTime.Format("2006"), currentTime.Format("01")

	folderPath = fmt.Sprintf("./logs/%s/%s", y, m)
	filePath = fmt.Sprintf("./logs/%s/%s/%s.log", y, m, d)

	// check if the directory exists, otherwise create it
	if err := os.MkdirAll(folderPath, 0644); err != nil {
		// if an error occurred, log it to the command line
		log.Panic(err)
	} else {
		logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Panic(err)
		}
		defer logFile.Close()

		log.SetOutput(logFile)
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		log.Printf("%v\t%v\tCalled in function %v in line %v of file %v (package %v)",
			logType,
			message,
			funcName,
			line,
			fileName,
			packageName,
		)
	}
}

func Debug(message string) {
	logToFile(DEBUG, message)
}

func Error(message string) {
	logToFile(ERROR, message)
}
