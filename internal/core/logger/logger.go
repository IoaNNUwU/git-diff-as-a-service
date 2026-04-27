package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

type Logger struct {
	slog.Logger

	file *os.File
}

const (
	Debug string = "DEBUG"
	Info  string = "INFO"
	Warn  string = "WARN"
	Error string = "ERROR"
)

const (
	StdOut = "STDOUT"
	None   = "NONE"
)

const (
	Json      = "JSON"
	PlainText = "PLAINTEXT"
)

func MustNewLogger(config Config) *Logger {

	if config.Stream == None {
		fmt.Println("Logging disabled for git-diff-app")
		slog := *slog.New(slog.NewTextHandler(io.Discard, nil))
		return &Logger{slog, nil}
	}

	var slogLvl slog.Level

	switch config.Level {
	case Debug:
		slogLvl = slog.LevelDebug
	case Info:
		slogLvl = slog.LevelInfo
	case Warn:
		slogLvl = slog.LevelWarn
	case Error:
		slogLvl = slog.LevelError
	default:
		panic(fmt.Sprintf(
			"Log level %q does not exist. Should be one of %q, %q, %q, %q", config.Level, Debug, Info, Warn, Error,
		))
	}
	opts := slog.HandlerOptions{}
	opts.Level = slogLvl

	if config.AddSource {
		opts.AddSource = true
	}

	if config.Format != Json && config.Format != PlainText {
		panic(fmt.Sprintf(
			"Log format %q does not exist. Should be one of %q, %q", config.Format, PlainText, Json,
		))
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15-04-05.00000")

	var file *os.File = nil
	if config.Folder != "" {
		logFilePath := filepath.Join(
			config.Folder,
			fmt.Sprintf("%s.log", timestamp),
		)
		if err := os.MkdirAll(config.Folder, 0755); err != nil {
			panic(fmt.Sprintf("unable to create logger: unable to create folder %q: %s", config.Folder, err))
		}
		logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(fmt.Sprintf("unable to create logger: unable to create log file %q: %s", logFilePath, err))
		}

		file = logFile
	}

	var stream io.Writer = nil
	if config.Stream == StdOut {
		stream = os.Stdout
	}

	var handler slog.Handler
	if file != nil && stream != nil {
		if config.Format == Json {
			handler = slog.NewMultiHandler(
				slog.NewJSONHandler(file, &opts),
				slog.NewJSONHandler(stream, &opts),
			)
		} else {
			handler = slog.NewMultiHandler(
				slog.NewTextHandler(file, &opts),
				slog.NewTextHandler(stream, &opts),
			)
		}
	} else if file != nil {
		if config.Format == Json {
			handler = slog.NewJSONHandler(file, &opts)
		} else {
			handler = slog.NewTextHandler(file, &opts)
		}
	} else if stream != nil {
		if config.Format == Json {
			handler = slog.NewJSONHandler(stream, &opts)
		} else {
			handler = slog.NewTextHandler(stream, &opts)
		}
	} else {
		panic("unable to create logger: both file and stream are nil")
	}

	slogger := slog.New(handler)
	return &Logger{*slogger, file}
}

func (l *Logger) Close() {
	if l.file != nil {
		if err := l.file.Close(); err != nil {
			fmt.Println("unable to close log file: ", err)
		}
	}
}

func FromContext(ctx context.Context) *Logger {
	log, ok := ctx.Value("log").(*Logger)
	if !ok {
		panic("unable to get logger from context. perhaps this function was called before AddLogger middleware")
	}
	return log
}

func (l *Logger) With(args ...any) *Logger {
	slog := l.Logger.With(args...)
	return &Logger{
		Logger: *slog,
	}
}

func String(key string, value string) slog.Attr {
	return slog.String(key, value)
}

func Err(key string, err error) slog.Attr {
	if err != nil {
		return slog.String(key, err.Error())	
	}
	return slog.String("", "")
}