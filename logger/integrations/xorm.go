// Package integrations has every 'middleware' to convert a library specific logger format to the logrus format.
package integrations

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"hextechdocs-be/logger"
	"runtime"
	"strings"
	"xorm.io/xorm/log"
)

// XormLogger is responsible for translating between the xorm logger format and logrus
type XormLogger struct {
	level		log.LogLevel
	showSql		bool
}

// Debug logs a message at level Debug on the standard logger.
func (XormLogger) Debug(args ...interface{}) {
	entry := logger.HexLogger.WithFields(logrus.Fields{
		"file-line": fileInfo(2),
	})
	entry.Debug(args)
}

// Debug logs a message with fields at level Debug on the standard logger.
func (XormLogger) Debugf(format string, args ...interface{}) {
	entry := logger.HexLogger.WithFields(logrus.Fields{
		"file-line": fileInfo(2),
	})
	entry.Debugf(format, args...)
}

// Debug logs a message at level Debug on the standard logger.
func (XormLogger) Info(args ...interface{}) {
	entry := logger.HexLogger.WithFields(logrus.Fields{
		"file-line": fileInfo(2),
	})
	entry.Info(args)
}

// Debug logs a message with fields at level Debug on the standard logger.
func (XormLogger) Infof(format string, args ...interface{}) {
	entry := logger.HexLogger.WithFields(logrus.Fields{
		"file-line": fileInfo(2),
	})
	entry.Infof(format, args...)
}

// Debug logs a message at level Debug on the standard logger.
func (XormLogger) Warn(args ...interface{}) {
	entry := logger.HexLogger.WithFields(logrus.Fields{
		"file-line": fileInfo(2),
	})
	entry.Warn(args)
}

// Debug logs a message with fields at level Debug on the standard logger.
func (XormLogger) Warnf(format string, args ...interface{}) {
	entry := logger.HexLogger.WithFields(logrus.Fields{
		"file-line": fileInfo(2),
	})
	entry.Warnf(format, args...)
}

// Debug logs a message at level Debug on the standard logger.
func (XormLogger) Error(args ...interface{}) {
	entry := logger.HexLogger.WithFields(logrus.Fields{
		"file-line": fileInfo(2),
	})
	entry.Error(args)
}

// Debug logs a message with fields at level Debug on the standard logger.
func (XormLogger) Errorf(format string, args ...interface{}) {
	entry := logger.HexLogger.WithFields(logrus.Fields{
		"file-line": fileInfo(2),
	})
	entry.Errorf(format, args...)
}

func (l *XormLogger) Level() log.LogLevel {
	return l.level
}
func (l *XormLogger) SetLevel(newLevel log.LogLevel) {
	l.level = newLevel
}

func (l *XormLogger) ShowSQL(show ...bool) {
	l.showSql = show[0]
}

func (l *XormLogger) IsShowSQL() bool {
	return l.showSql
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}
