package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 在文件里面写日志相关代码

// FileLogger 结构体
type FileLogger struct {
	level       LogLevel
	filePath    string   // 日志文件保存的路径
	fileName    string   // 日志文件保存的文件名
	fileObj     *os.File // 文件对象
	errFileObj  *os.File
	maxFileSize int64
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

// log 记录日志的方法  传入格式化参数和空接口(任意个和任意类型参数都可) 仿照的print()传参
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funceName, fileName, lineNo := getInfo(3)
		if f.checkSize(f.fileObj) {
			// 切割文件
			newFile, err := f.splitFile(f.fileObj) // 日志文件
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d]%s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funceName, fileName, lineNo, msg)
		if lv >= ERROR {
			if f.checkSize(f.errFileObj) {
				// 切割文件
				newFile, err := f.splitFile(f.errFileObj) // 日志文件
				if err != nil {
					return
				}
				f.errFileObj = newFile
			}
			// 如果要记录的日志大于等于ERROR级别,我要在err日志文件中再记录一遍
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d]%s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funceName, fileName, lineNo, msg)
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
