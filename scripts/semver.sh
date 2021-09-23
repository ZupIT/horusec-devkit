#!/bin/bash
# Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain versionArray copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Increment versionArray version string using Semantic Versioning (SemVer) terminology.

# Parse command line options.

while getopts ":Mmp" Option
do
  case $Option in
    M ) major=true;;
    m ) minor=true;;
    p ) patch=true;;
  esac
done

shift $(($OPTIND - 1))

repositoryName=$1

echo "cd to github workspace"
cd ${GITHUB_WORKSPACE}

version=$(curl -sL https://api.github.com/repos/$repositoryName/releases/latest | jq -r ".tag_name")
echo "Actual version: ${version}"
echo "::set-output name=actualVersion::${version}"

if [ -z ${version} ]
then
    echo "Couldn't determine version"
    exit 1
fi
# Build array from version string.

versionArray=( ${version//./ } )
major_version=0

# If version string is missing or has the wrong number of members, show usage message.

if [ ${#versionArray[@]} -ne 3 ]
then
  echo "usage: $(basename $0) [-Mmp] major.minor.patch"
  exit 1
fi

# Increment version numbers as requested.

if [ ! -z $major ]
then
# Check for v in version (e.g. v1.0 not just 1.0)
  if [[ ${versionArray[0]} =~ ([vV]?)([0-9]+) ]]
  then
    v="${BASH_REMATCH[1]}"
    major_version=${BASH_REMATCH[2]}
    ((major_version++))
    versionArray[0]=${v}${major_version}
  else
    ((versionArray[0]++))
    major_version=versionArray[0]
  fi

  versionArray[1]=0
  versionArray[2]=0
fi

if [ ! -z $minor ]
then
  ((versionArray[1]++))
  versionArray[2]=0
fi

if [ ! -z $patch ]
then
  ((versionArray[2]++))
fi

echo "New version ${versionArray[0]}.${versionArray[1]}.${versionArray[2]}"

version=$(echo "${versionArray[0]}.${versionArray[1]}.${versionArray[2]}")
echo "::set-output name=version::${version}"

echo "::set-output name=strippedVersion::$(echo ${version} | sed 's/v//')"

releaseBranchName=$(echo "${versionArray[0]}.${versionArray[1]}")
echo "::set-output name=releaseBranchName::release/${releaseBranchName}"
