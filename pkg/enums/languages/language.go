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

package languages

import "strings"

type Language string

const (
	Go         Language = "Go"
	CSharp     Language = "C#"
	Dart       Language = "Dart"
	Ruby       Language = "Ruby"
	Python     Language = "Python"
	Java       Language = "Java"
	Kotlin     Language = "Kotlin"
	Javascript Language = "Javascript"
	Typescript Language = "Typescript"
	Leaks      Language = "Leaks"
	HCL        Language = "HCL"
	C          Language = "C"
	PHP        Language = "PHP"
	HTML       Language = "HTML"
	Generic    Language = "Generic"
	Yaml       Language = "YAML"
	Elixir     Language = "Elixir"
	Shell      Language = "Shell"
	Unknown    Language = "Unknown"
)

func ParseStringToLanguage(value string) (l Language) {
	for key, lang := range mapEnableLanguages() {
		if strings.EqualFold(key, value) {
			return lang
		}
	}

	return Unknown
}

//nolint:funlen // method need to have more then 15 lines
func Values() []Language {
	return []Language{
		Go,
		CSharp,
		Dart,
		Ruby,
		Python,
		Java,
		Kotlin,
		Javascript,
		Typescript,
		Leaks,
		HCL,
		C,
		PHP,
		HTML,
		Generic,
		Yaml,
		Elixir,
		Shell,
		Unknown,
	}
}

//nolint:funlen // method need to have more then 15 lines
func mapEnableLanguages() map[string]Language {
	return map[string]Language{
		Go.ToString():         Go,
		CSharp.ToString():     CSharp,
		Dart.ToString():       Dart,
		Ruby.ToString():       Ruby,
		Python.ToString():     Python,
		Java.ToString():       Java,
		Kotlin.ToString():     Kotlin,
		Javascript.ToString(): Javascript,
		Typescript.ToString(): Typescript,
		Leaks.ToString():      Leaks,
		HCL.ToString():        HCL,
		C.ToString():          C,
		PHP.ToString():        PHP,
		HTML.ToString():       HTML,
		Generic.ToString():    Generic,
		Yaml.ToString():       Yaml,
		Elixir.ToString():     Elixir,
		Shell.ToString():      Shell,
		Unknown.ToString():    Unknown,
	}
}

func (l Language) ToString() string {
	return string(l)
}

func (l Language) GetCustomImagesKeyByLanguage() string {
	return l.mapConfigCustomImageJSONByLanguage()[l]
}

//nolint:funlen // method need to have more then 15 lines
func (l Language) mapConfigCustomImageJSONByLanguage() map[Language]string {
	return map[Language]string{
		CSharp:     "csharp",
		Leaks:      "leaks",
		Go:         "go",
		Javascript: "javascript",
		Python:     "python",
		Ruby:       "ruby",
		HCL:        "hcl",
		Generic:    "generic",
		PHP:        "php",
		Elixir:     "elixir",
		Shell:      "shell",
		C:          "c",
		Java:       "java",
		Kotlin:     "kotlin",
		Yaml:       "yaml",
		Dart:       "dart",
	}
}
