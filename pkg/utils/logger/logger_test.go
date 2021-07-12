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
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/utils/logger/enums"
)

func TestLogPanic(t *testing.T) {
	t.Run("should log error and panic", func(t *testing.T) {
		assert.Panics(t, func() { LogPanic("test error", errors.New("test")) })
	})

	t.Run("should log error with args and panic", func(t *testing.T) {
		args := map[string]interface{}{"test": "test"}
		assert.Panics(t, func() { LogPanic("test error", errors.New("test"), args) })
	})
}

func TestLogError(t *testing.T) {
	t.Run("should log error without panic", func(t *testing.T) {
		assert.NotPanics(t, func() { LogError("test error", errors.New("test")) })
	})

	t.Run("should log error with args without panic", func(t *testing.T) {
		args := map[string]interface{}{"test": "test"}
		assert.NotPanics(t, func() { LogError("test error", errors.New("test"), args) })
	})
}

func TestLogInfo(t *testing.T) {
	t.Run("should log information log without panic", func(t *testing.T) {
		assert.NotPanics(t, func() { LogInfo("test") })
	})

	t.Run("should log information log without panic", func(t *testing.T) {
		args := map[string]interface{}{"test": "test"}
		assert.NotPanics(t, func() { LogInfo("test", args) })
	})
}

func TestLogWarn(t *testing.T) {
	t.Run("should log warning log without panic", func(t *testing.T) {
		assert.NotPanics(t, func() { LogWarn("test") })
	})

	t.Run("should log warning log without panic", func(t *testing.T) {
		args := map[string]interface{}{"test": "test"}
		assert.NotPanics(t, func() { LogWarn("test", args) })
	})
}

func TestLogPrint(t *testing.T) {
	t.Run("should log print log without panic", func(t *testing.T) {
		assert.NotPanics(t, func() { LogPrint("test") })
	})
}

func TestSetLogLevel(t *testing.T) {
	t.Run("should success set level", func(t *testing.T) {
		assert.NotPanics(t, func() { SetLogLevel(enums.WarnLevel.String()) })
	})

	t.Run("should set info level when invalid value", func(t *testing.T) {
		assert.NotPanics(t, func() { SetLogLevel("test") })
	})
}

func TestLogPanicWithLevel(t *testing.T) {
	SetLogLevel(enums.PanicLevel.String())
	t.Run("should panic with error", func(t *testing.T) {
		assert.Panics(t, func() { LogPanicWithLevel("test", errors.New("test")) })
	})

	t.Run("should panic with args", func(t *testing.T) {
		assert.Panics(t, func() {
			LogPanicWithLevel("test", errors.New("test"),
				map[string]interface{}{})
		})
	})
}

func TestLogErrorWithLevel(t *testing.T) {
	SetLogLevel(enums.ErrorLevel.String())
	t.Run("should not panic", func(t *testing.T) {
		assert.NotPanics(t, func() { LogErrorWithLevel("test", errors.New("test")) })
	})

	t.Run("should not panic when log with args", func(t *testing.T) {
		assert.NotPanics(t, func() {
			LogErrorWithLevel("test", errors.New("test"),
				map[string]interface{}{})
		})
	})
}

func TestLogWarnWithLevel(t *testing.T) {
	SetLogLevel(enums.WarnLevel.String())
	t.Run("should not panic", func(t *testing.T) {
		assert.NotPanics(t, func() { LogWarnWithLevel("test") })
	})

	t.Run("should not panic when log with args", func(t *testing.T) {
		assert.NotPanics(t, func() { LogWarnWithLevel("test", map[string]interface{}{}) })
	})
}

func TestLogInfoWithLevel(t *testing.T) {
	SetLogLevel(enums.InfoLevel.String())
	t.Run("should not panic", func(t *testing.T) {
		assert.NotPanics(t, func() { LogInfoWithLevel("test") })
	})

	t.Run("should not panic when log with args", func(t *testing.T) {
		assert.NotPanics(t, func() { LogInfoWithLevel("test", map[string]interface{}{}) })
	})
}

func TestLogDebugWithLevel(t *testing.T) {
	SetLogLevel(enums.DebugLevel.String())
	t.Run("should not panic", func(t *testing.T) {
		assert.NotPanics(t, func() { LogDebugWithLevel("test") })
	})

	t.Run("should not panic when log with args", func(t *testing.T) {
		assert.NotPanics(t, func() { LogDebugWithLevel("test", map[string]interface{}{}) })
	})
}

func TestLogTraceWithLevel(t *testing.T) {
	SetLogLevel(enums.TraceLevel.String())
	t.Run("should not trace", func(t *testing.T) {
		assert.NotPanics(t, func() { LogTraceWithLevel("test") })
	})

	t.Run("should not panic when log with args", func(t *testing.T) {
		assert.NotPanics(t, func() { LogTraceWithLevel("test", map[string]interface{}{}) })
	})
}

func TestLogStringAsError(t *testing.T) {
	t.Run("should not panic", func(t *testing.T) {
		assert.NotPanics(t, func() { LogStringAsError("test") })
	})
}

func TestLogDebugJSON(t *testing.T) {
	t.Run("should log json bytes and not panic", func(t *testing.T) {
		assert.NotPanics(t, func() { LogDebugJSON("test", map[string]string{"test": "test"}) })
	})

	t.Run("should log json string when failed to marshal and not panic", func(t *testing.T) {
		assert.NotPanics(t, func() { LogDebugJSON("test", make(chan int)) })
	})
}

func TestLogSetOutput(t *testing.T) {
	t.Run("Should set output instance without file and get on read output", func(t *testing.T) {
		output := bytes.NewBufferString("")
		LogSetOutput(output)
		assert.Empty(t, output.String())

		const textLogged = "Some aleatory text logged"
		LogInfo(textLogged)

		assert.Contains(t, output.String(), textLogged)
	})
	t.Run("Should set output instance with file and get on read output and file", func(t *testing.T) {
		output := bytes.NewBufferString("")
		file, err := os.Create("./testSetOutput")
		if err != nil {
			t.Error(err)
		}
		defer func() {
			err := os.Remove(file.Name())
			if err != nil {
				t.Error()
			}
		}()
		LogSetOutput(output, file)
		assert.Empty(t, output.String())

		const textLogged = "Some aleatory text logged"
		LogInfo(textLogged)
		byteSlice, err := ioutil.ReadFile(file.Name())
		if err != nil {
			t.Error(err)
		}
		assert.Contains(t, output.String(), textLogged)
		assert.Contains(t, string(byteSlice), textLogged)

	})
	t.Run("Should set output instance with invalid writer and panic", func(t *testing.T) {
		assert.Panics(t, func() {
			LogSetOutput(nil, &os.File{})
			LogInfo("Should panic")
		})
	})

}
