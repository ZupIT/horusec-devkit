#!/bin/bash
# Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

validateSwagger () {
    if ! swag &> /dev/null
    then
        go get -v github.com/swaggo/swag/cmd/swag@v1.7.0
    fi
    if ! swag &> /dev/null
    then
        echo "swag is not installed, please install and try again"
        exit 1
    fi

    echo "Lib swag installed with success!"
}

updateDocs () {
    swag init -g ./cmd/app/main.go --parseDependency
}

validateSwagger

updateDocs