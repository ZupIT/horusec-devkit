// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

package enums

import "errors"

var ErrorNotFoundRecords = errors.New("{ERROR_DATABASE} database not found records")

//nolint
var ErrorConnectingToDB = errors.New("{ERROR_DATABASE} error connecting to db, use this format" +
	" string for connection in " + EnvRelationalURI +
	": 'host=localhost user=username password=user_password dbname=db_name port=5432 sslmode=disable TimeZone=Asia/Shanghai'")
