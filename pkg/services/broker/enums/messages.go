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

const (
	MessageFailedConnectBroker            = "{ERROR_BROKER} failed to connect"
	MessageFailedCreateChannelPublish     = "{ERROR_BROKER} failed to create channel while publishing"
	MessageFailedDeclareExchangePublish   = "{ERROR_BROKER} failed to declare exchange while publishing"
	MessageFailedCreateChannelConsume     = "{ERROR_BROKER} failed to create channel in consume"
	MessageFailedCreateQueueConsume       = "{ERROR_BROKER} error declaring queue in consumer"
	MessageFailedConsumeHandlingDelivery  = "{ERROR_BROKER} consume error while handling deliveries"
	MessageFailedSetConsumerPrefetch      = "{ERROR_BROKER} failed to set consumer prefetch"
	MessageFailedToDeclareExchangeQueue   = "{ERROR_BROKER} failed to declare exchange while declaring queue"
	MessageFailedBindQueueConsume         = "{ERROR_BROKER} failed to queue bind in consume"
	MessageWarningDefaultBrokerConnection = "{WARN} your user or password for connection with message broker " +
		"is default content, please change for you best security"
)
