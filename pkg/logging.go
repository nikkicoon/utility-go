package pkg

import (
	"fmt"
	phuslu "github.com/phuslu/log"
	"io"
	"log"
	"log/slog"
	"net"
	"os"
)

// InitializeSlogLogger initializes a new `log/slog` logger.
// level: loglevel (debug, info, warning, error)
// logfile: either stdout, stderr or file
// format: json or text, every other argument defaults to text
func InitializeSlogLogger(level string, logfile string, format string) *slog.Logger {
	var logFile io.Writer
	switch logfile {
	case "stdout":
		logFile = os.Stdout
	case "stderr":
		logFile = os.Stderr
	case "file":
		f, err := os.OpenFile("/var/log/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		logFile = f
	}
	var l = new(slog.LevelVar)
	switch level {
	case "debug":
		l.Set(slog.LevelDebug)
	case "info":
		l.Set(slog.LevelInfo)
	case "warning":
		l.Set(slog.LevelWarn)
	case "error":
		l.Set(slog.LevelError)
	}
	handlerOptions := &slog.HandlerOptions{Level: l}
	switch format {
	case "json":
		return slog.New(slog.NewJSONHandler(logFile, handlerOptions))
	case "text":
		return slog.New(slog.NewTextHandler(logFile, handlerOptions))
	default:
		return slog.New(slog.NewTextHandler(logFile, handlerOptions))
	}
}

type SyslogConn struct {
	Protocol string
	Port     string
	Address  string
}

// InitializeLogger initializes a new logger.
// level: loglevel (trace, debug, info, warn, error, fatal, panic)
// logType: simple, consolewriter (same as text), consolewritermulti, file, systemd (only available on linux)
// logFile: file to log to if logType is file, stdout or stderr for simple (defaults to stderr if empty)
// localtime: boolean, use local time in logs
// caller: add `file:line` to log output
func InitializeLogger(level string, logType string, logFile string, localtime bool, caller int, sysConn SyslogConn, writer ...*phuslu.Writer) *phuslu.Logger {
	var f *os.File
	switch logFile {
	case "stdout":
		f = os.Stdout
	case "stderr":
		f = os.Stderr
	default:
		f = os.Stderr
	}
	switch logType {
	case "simple":
		return &phuslu.Logger{
			Level:  phuslu.ParseLevel(level),
			Caller: caller,
			Writer: &phuslu.IOWriter{Writer: f},
		}
	case "text", "consolewriter":
		return &phuslu.Logger{
			Level:  phuslu.ParseLevel(level),
			Caller: caller,
			Writer: &phuslu.ConsoleWriter{
				Formatter: phuslu.LogfmtFormatter{TimeField: "time"}.Formatter,
				Writer:    &phuslu.IOWriter{Writer: f},
			},
		}
	case "textmulti", "consolewritermulti":
		return &phuslu.Logger{
			Level:  phuslu.ParseLevel(level),
			Caller: caller,
			Writer: &phuslu.ConsoleWriter{
				Formatter: phuslu.LogfmtFormatter{TimeField: "time"}.Formatter,
				Writer:    io.MultiWriter(os.Stdout, os.Stderr),
			},
		}
	case "file":
		if logFile == "" {
			logFile = "/var/log/out.log"
		}
		return &phuslu.Logger{
			Level:  phuslu.ParseLevel(level),
			Caller: caller,
			Writer: &phuslu.AsyncWriter{
				ChannelSize: 4096,
				Writer: &phuslu.FileWriter{
					Filename:   logFile,
					FileMode:   0600,
					MaxSize:    50 * 1024 * 1024,
					MaxBackups: 7,
					LocalTime:  localtime,
				},
			},
		}
	case "rsyslogd":
		if IfAnyEmpty(sysConn) {
			_, _ = fmt.Fprintf(os.Stderr, "rsyslogd logger: conn is empty")
			return nil
		}
		return &phuslu.Logger{
			Level:      phuslu.ParseLevel(level),
			Caller:     caller,
			TimeField:  "ts",
			TimeFormat: phuslu.TimeFormatUnixMs,
			Writer: &phuslu.SyslogWriter{
				Network: sysConn.Protocol,
				Address: sysConn.Address + ":" + sysConn.Port,
				Tag:     "",
				Marker:  "",
				Dial:    net.Dial,
			},
		}
	default:
		if len(writer) > 0 {
			return &phuslu.Logger{
				Level:  phuslu.ParseLevel(level),
				Caller: caller,
				Writer: *writer[0],
			}
		} else {
			_, _ = fmt.Fprintf(os.Stderr, "no logger type matched")
			return nil
		}
	}
}
