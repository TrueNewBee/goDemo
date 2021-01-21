package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 在文件里面写日志相关代码

var (
	// MaxSize 通道缓冲区的最大容量
	MaxSize int = 50000
)

// FileLogger 结构体
type FileLogger struct {
	level       LogLevel
	filePath    string   // 日志文件保存的路径
	fileName    string   // 日志文件保存的文件名
	fileObj     *os.File // 文件对象
	errFileObj  *os.File
	maxFileSize int64
	logChan     chan *logMsg // string类型太大,故用结构体
}

// logMsg :创建一个chan的结构体
type logMsg struct {
	level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	line      int
}

// NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
		logChan:     make(chan *logMsg, MaxSize), // 初始化管道
	}
	fl.initFile() // 按照文件路径和文件名将文件打开
	return fl
}

// 根据指定的日志文件路径和文件名打开日志文件
func (f *FileLogger) initFile() error {
	// 找到路径
	fullFileName := path.Join(f.filePath, f.fileName)
	// 正确文件
	fileObje, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return err
	}
	// 错误文件
	errfileObje, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed, err:%v\n", err)
		return err
	}
	// 日志文件都打开
	f.fileObj = fileObje
	f.errFileObj = errfileObje
	// 开启5个后台的goroutine去写日志
	// 由于需要 check 所以开一个goroutine就好了
	// 我试了开2 goroutine个和普通的不异步操作谁快
	// 异步的快  开一个 写的文件
	// for i := 0; i < 1; i++ {
	// 	go f.writeLogBackground()
	// }
	go f.writeLogBackground()
	return nil
}

// 判断是否需要记录该文件
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.level
}

// 判断文件文件大小是否需要切割
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("open file info  failed , err:%v\n", err)
		return false
	}
	// 如果当前文件爱你大小  大于等于 日志文件的最大值 就应该返回True
	return fileInfo.Size() >= f.maxFileSize
}

// 切割文件
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 需要切割日志文件
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name()) // 拿到当前文件完整的名字
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
	// 1. 关闭当前的日志文件
	file.Close()
	// 2. 备份一下 rename  xx.log -> xx.log.bak202101182023
	os.Rename(logName, newLogName)
	// 3. 打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed,err:%v\n", err)
		return nil, err
	}
	// 4. 将打开的新日志文件对象赋值给  f.fileObj
	return fileObj, nil
}

// 后台写日志
func (f *FileLogger) writeLogBackground() {
	for {
		if f.checkSize(f.fileObj) {
			// 切割文件
			newFile, err := f.splitFile(f.fileObj) // 日志文件
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		select {
		case logTmp := <-f.logChan: // 先把通道日志内容取出
			// 把日志先拼出来
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d]%s\n", logTmp.timestamp, getLogString(logTmp.level), logTmp.funcName, logTmp.fileName, logTmp.line, logTmp.msg)
			fmt.Fprintf(f.fileObj, logInfo)
			if logTmp.level >= ERROR {
				if f.checkSize(f.errFileObj) {
					// 切割文件
					newFile, err := f.splitFile(f.errFileObj) // 日志文件
					if err != nil {
						return
					}
					f.errFileObj = newFile
				}
				// 如果要记录的日志大于等于ERROR级别,我要在err日志文件中再记录一遍
				fmt.Fprintf(f.errFileObj, logInfo)
			}
		default:
			// 取不到就休息500毫秒
			time.Sleep(time.Millisecond * 500)
		}
	}
}

// log 记录日志的方法  传入格式化参数和空接口(任意个和任意类型参数都可) 仿照的print()传参
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funceName, fileName, lineNo := getInfo(3)
		// 先把日志发送到日志通道中
		// 1. 造一个logMsg对象
		logtTmp := &logMsg{
			level:     lv,
			msg:       msg,
			funcName:  funceName,
			fileName:  fileName,
			timestamp: now.Format("2006-01-02 15:04:05"),
			line:      lineNo,
		}
		select {
		case f.logChan <- logtTmp:
		default:
			// 把日志丢掉保证不阻塞
		}
	}
}

// Debug 方法
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

// Info 方法
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

// Warning 方法
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

// Error 方法
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

// Fatal 方法
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}

// Close  关闭文件
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
