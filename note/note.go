/*Package log 自己的记录类，方便记录特定领域的信息
 */
package log

import (
	"fmt"
	"os"
	"sync"

	logMod "hippogames.com.cn/common/log"
	"hippogames.com.cn/common/str"
	"hippogames.com.cn/common/time"
)

var (
	logger   *logMod.Logger
	curDay   int
	outFile  *os.File
	syncLock sync.Mutex
)

func init() {
	fileName := calFileName() // 生成文件名
	outFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	logger = logMod.New(outFile, "", logMod.Lshortfile|logMod.Ltime|logMod.Ldebug)

	logLvl := logMod.Linfo
	logger.SetFlags(logger.Flags() | logLvl)

	if err != nil {
		logger.Fatal(err.Error())
	}

	//修改默认的日志输出对象
	logMod.SetOutput(outFile)
}

func switchFile() {
	syncLock.Lock()
	defer syncLock.Unlock()

	nowDay := time.GetDay()
	if nowDay == curDay { // 没变天
		return
	}

	newFileName := calFileName()
	outFD, err := os.OpenFile(newFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if nil == err {
		logger.SetFile(outFD)
		outFile.Close()
		outFile = outFD
	}
	fmt.Printf("[note.switch.file] path(%s)", newFileName)
}

// 生成文件名
func calFileName() string {
	y, m, d := time.GetDate()
	curDay = d
	fileName := str.Str(y) + "_" + str.Str(m) + "_" + str.Str(d) + ".log"
	{ // 写入指定的路径
		prePath := "./logs/"
		fileName = prePath + fileName
	}
	return fileName
}

// SetPrefix 设置前缀
func SetPrefix(prefix string) {
	logger.SetPrefix(prefix)
}

// Debug 调试
func Debug(format string, v ...interface{}) {
	switchFile()
	logger.Debug(fmt.Sprintf(format, v...))
}

// Log 日志
func Log(format string, v ...interface{}) {
	switchFile()
	logger.Info(fmt.Sprintf(format, v...))
}

// Warn 警告
func Warn(format string, v ...interface{}) {
	switchFile()
	logger.Warn(fmt.Sprintf(format, v...))
}

// Err 错误
func Err(format string, v ...interface{}) {
	switchFile()
	logger.Error(fmt.Sprintf(format, v...))
}

// Panic 抛出异常
func Panic(format string, v ...interface{}) {
	switchFile()
	logger.Panic(fmt.Sprintf(format, v...))
}

// Fatal 退出程序，输出
func Fatal(format string, v ...interface{}) {
	switchFile()
	logger.Fatal(fmt.Sprintf(format, v...))
}
