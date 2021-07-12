// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger/enums"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
)

func LogPanic(msg string, err error, args ...map[string]interface{}) {
	if err != nil {
		if len(args) > 0 {
			logrus.WithFields(args[0]).WithError(err).Panic(msg)
			return
		}

		logrus.WithError(err).Panic(msg)
	}
}

func LogError(msg string, err error, args ...map[string]interface{}) {
	if err != nil {
		if len(args) > 0 {
			logrus.WithFields(args[0]).WithError(err).Error(msg)
			return
		}

		logrus.WithError(err).Error(msg)
	}
}

func LogInfo(msg string, args ...interface{}) {
	if args != nil {
		logrus.Info(msg, args)
		return
	}
	logrus.Info(msg)
}

func LogWarn(msg string, args ...interface{}) {
	if args != nil {
		logrus.Warn(msg, args)
		return
	}

	logrus.Warn(msg)
}

func LogPrint(msg string) {
	log.SetFlags(0)
	log.Println(msg)
}

func SetLogLevel(level string) {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.Error(fmt.Sprintf(enums.MessageInvalidLogLevel, level, enums.InfoLevel.String()))
		logLevel = enums.InfoLevel
	}

	logrus.SetLevel(logLevel)
}

func LogPanicWithLevel(msg string, err error, args ...map[string]interface{}) {
	if logrus.IsLevelEnabled(enums.PanicLevel) && err != nil {
		if len(args) > 0 {
			logrus.WithFields(args[0]).WithError(err).Panic(msg)
		}

		logrus.WithError(err).Panic(msg)
	}
}

func LogErrorWithLevel(msg string, err error, args ...map[string]interface{}) {
	if logrus.IsLevelEnabled(enums.ErrorLevel) && err != nil {
		if len(args) > 0 {
			logrus.WithFields(args[0]).WithError(err).Error(msg)
			return
		}

		logrus.WithError(err).Error(msg)
	}
}

func LogWarnWithLevel(msg string, args ...interface{}) {
	if logrus.IsLevelEnabled(enums.WarnLevel) {
		if args != nil {
			logrus.Warn(msg, args)
			return
		}

		logrus.Warn(msg)
	}
}

func LogInfoWithLevel(msg string, args ...interface{}) {
	if logrus.IsLevelEnabled(enums.InfoLevel) {
		if args != nil {
			logrus.Info(msg, args)
			return
		}

		logrus.Info(msg)
	}
}

func LogDebugWithLevel(msg string, args ...interface{}) {
	if logrus.IsLevelEnabled(enums.DebugLevel) {
		if args != nil {
			logrus.Debug(msg, args)
			return
		}

		logrus.Debug(msg)
	}
}

func LogTraceWithLevel(msg string, args ...interface{}) {
	if logrus.IsLevelEnabled(enums.TraceLevel) {
		if args != nil {
			logrus.Trace(msg, args)
			return
		}

		logrus.Trace(msg)
	}
}

func LogStringAsError(msg string) {
	logrus.Error(msg)
}

func LogDebugJSON(message string, content interface{}) {
	if contentBytes, err := json.Marshal(content); err == nil {
		LogDebugWithLevel(message, string(contentBytes))
		return
	}

	LogDebugWithLevel(message, fmt.Sprintf("%v", content))
}

func LogSetOutput(stdout *bytes.Buffer, file *os.File) error {
	var err error
	if file == nil {
		logrus.SetOutput(stdout)
	} else {
		file, err = os.OpenFile(file.Name(), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			logrus.Error(err)
			return err
		}
		mw := io.MultiWriter(stdout, file)
		logrus.SetOutput(mw)
	}
	return nil
}
