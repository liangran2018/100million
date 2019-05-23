package base

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/liangran2018/lived/base"
)

type LogLvl int

var logFile = &os.File{}

const (
	Info LogLvl = iota
	Warning
	Wrong
)

func NewLogFile() {
	if ok := base.Exists("./operaLog"); !ok {
		if err := os.Mkdir("./operaLog", os.ModePerm); err != nil {
			panic(err)
		}
	}
	
	files, err := ioutil.ReadDir("./operaLog")
	if err != nil {
		panic(err)
	}

	i := 1
	for _, file := range files {
		if file.IsDir() || base.Exists(fmt.Sprintf("./operaLog/log%d.txt", i)) {
			i++
			continue
		}

		break
	}

	file, err := os.Create(fmt.Sprintf("./operaLog/log%d.txt", i))
	if err != nil {
		panic(err)
	}

	logFile = file
}

//func OpenLogFile(i int) {
//	file, err := os.Open(fmt.Sprintf("./operaLog/log%s.txt", strconv.Itoa(i)))
//	if err != nil {
//		panic(err)
//	}
//
//	logFile.l = file
//}

func notice(lvl LogLvl) string {
	switch lvl {
	case Info:
		return "[info] "
	case Warning:
		return "[Warn] "
	case Wrong:
		return "[ERR] "
	default:
	}

	return ""
}

func Log(lvl LogLvl, opera ...interface{}) {
	logFile.WriteString(notice(lvl))
	logFile.WriteString(time.Now().Format("2006-01-02 15:04:05") + " ")
	if len(opera) == 0 {
		return
	}

	for _, o := range opera {
		logFile.WriteString(base.StrVal(o) + " ")
	}
	logFile.WriteString("\n")
}

func CloseLog() {
	logFile.Close()
}
