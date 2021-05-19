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

package queues

import "strings"

type Queue string

const (
	HorusecAnalyticNewAnalysisByAuthors    Queue = "horusec-analytic::new-analysis-by-author"
	HorusecAnalyticNewAnalysisByRepository Queue = "horusec-analytic::new-analysis-by-repository"
	HorusecAnalyticNewAnalysisByLanguage   Queue = "horusec-analytic::new-analysis-by-language"
	HorusecAnalyticNewAnalysisByTime       Queue = "horusec-analytic::new-analysis-by-time"
	HorusecEmail                           Queue = "horusec-email"
	HorusecWebhook                         Queue = "horusec-webhook"
)

func Values() []Queue {
	return []Queue{
		HorusecAnalyticNewAnalysisByAuthors,
		HorusecAnalyticNewAnalysisByRepository,
		HorusecAnalyticNewAnalysisByLanguage,
		HorusecAnalyticNewAnalysisByTime,
		HorusecEmail,
		HorusecWebhook,
	}
}

func (q Queue) IsValid() bool {
	for _, value := range Values() {
		if q == value {
			return true
		}
	}

	return false
}

func (q Queue) IsInvalid() bool {
	return !q.IsValid()
}

func (q Queue) IsEqual(value string) bool {
	return strings.EqualFold(q.ToString(), value)
}

func ValueOf(value string) Queue {
	for _, queue := range Values() {
		if queue.IsEqual(value) {
			return queue
		}
	}

	return ""
}

func (q Queue) ToString() string {
	return string(q)
}
