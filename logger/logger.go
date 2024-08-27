/* SPDX-License-Identifier: AGPL-3.0-or-later
 *
 * SPDX-FileCopyrightText: 2024 Tristan Partin <tristan@partin.io>
 */

package logger

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var global *zap.SugaredLogger
var logLevel zap.AtomicLevel
var closeLogFiles func()

func Setup(debug bool) {
	closeLogFiles = func() {}

	var encoder zapcore.Encoder
	var output zapcore.WriteSyncer

	if debug {
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		logLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	logLocation := viper.GetString("log.location")
	if logLocation == "file" {
		logFilepath := viper.GetString("log.filepath")
		if logFilepath == "" {
			fmt.Fprintln(os.Stderr, "no log file path provided")
			os.Exit(1)
		}

		var err error
		output, closeLogFiles, err = zap.Open(logFilepath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open %s: %s", logFilepath, err)
			os.Exit(1)
		}
	} else if logLocation == "stdout" {
		output = zapcore.AddSync(os.Stdout)
	} else if logLocation == "stderr" {
		output = zapcore.AddSync(os.Stderr)
	} else if logLocation == "syslog" { //nolint:staticcheck
		if runtime.GOOS == "windows" || runtime.GOOS == "macos" {
			fmt.Fprintf(os.Stderr, "log.location cannot be set to syslog on %s\n", runtime.GOOS)
			os.Exit(1)
		}

		fmt.Fprintln(os.Stderr, "log.location set to syslog is currently unsupported")
		os.Exit(1)
	}

	core := zapcore.NewCore(encoder, output, logLevel)
	logger := zap.New(core)

	zap.ReplaceGlobals(logger)
	global = logger.Sugar()
}

func Debug(args ...any) {
	global.Debug(args)
}

func Debugf(template string, a ...any) {
	global.Debugf(template, a)
}

func Error(args ...any) {
	global.Error(args)
}

func Errorf(template string, a ...any) {
	global.Errorf(template, a)
}

func Fatal(args ...any) {
	global.Fatal(args)
}

func Fatalf(template string, a ...any) {
	global.Fatalf(template, a)
}

func Finish() {
	_ = global.Sync()
	closeLogFiles()
}

func HTTPEndpoint() func(http.ResponseWriter, *http.Request) {
	return logLevel.ServeHTTP
}

func Info(args ...any) {
	global.Info(args)
}

func Infof(template string, a ...any) {
	global.Infof(template, a)
}

func Log(lvl zapcore.Level, args ...any) {
	global.Log(lvl, args)
}

func Logf(lvl zapcore.Level, template string, a ...any) {
	global.Logf(lvl, template, a)
}

func Warn(args ...any) {
	global.Warn(args)
}

func Warnf(template string, a ...any) {
	global.Warn(template, a)
}
