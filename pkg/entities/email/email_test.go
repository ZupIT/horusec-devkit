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

package email

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBytes(t *testing.T) {
	t.Run("should success parse to bytes", func(t *testing.T) {
		message := &Message{To: "test", Subject: "test", TemplateName: "test", Data: "test"}

		assert.NotEmpty(t, message.ToBytes())
	})
}
