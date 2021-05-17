#!/bin/sh
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

THRESHOLD=$1
PROJECT=$2

go clean -testcache
go test -v -coverpkg "$PROJECT/..." -coverprofile coverage.out "$PROJECT/..."

sed -i '/mock.go/d' coverage.out
sed -i '/pb.go/d' coverage.out
sed -i '/wire_gen.go/d' coverage.out
sed -i '/wire.go/d' coverage.out
sed -i '/docs.go/d' coverage.out
sed -i '/main.go/d' coverage.out
sed -i '/libs/d' coverage.out
sed -i '/cmd\/migrate/d' coverage.out

COVERAGE=$(go tool cover -func=coverage.out | grep total: | awk '{print $3}')
COVERAGE=${COVERAGE%\%}

if [ 1 -eq "$(echo "$COVERAGE" '>=' "$THRESHOLD" | bc -l)" ]; then
  echo "SUCCESS! Coverage above threshold"
  echo "coverage: ${COVERAGE} - threshold: ${THRESHOLD}"
  rm coverage.out
  exit 0
fi

echo "FAILS! Coverage below threshold"
echo "coverage: ${COVERAGE} - threshold: ${THRESHOLD}"
rm coverage.out
exit 1
