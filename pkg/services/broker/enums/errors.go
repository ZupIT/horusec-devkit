package enums

const (
	FailedConnectBroker           = "{ERROR_BROKER} failed to connect"
	FailedCreateChannelPublish    = "{ERROR_BROKER} failed to create channel while publishing"
	FailedDeclareExchangePublish  = "{ERROR_BROKER} failed to declare exchange while publishing"
	FailedCreateChannelConsume    = "{ERROR_BROKER} failed to create channel in consume"
	FailedCreateQueueConsume      = "{ERROR_BROKER} error declaring queue in consume"
	FailedConsumeHandlingDelivery = "{ERROR_BROKER} consume error while handling deliveries"
	FailedSetConsumerPrefetch     = "{ERROR_BROKER} failed to set consumer prefetch"
	FailedToDeclareExchangeQueue  = "{ERROR_BROKER} failed to declare exchange while declaring queue"
	FailedBindQueueConsume        = "{ERROR_BROKER} failed to queue bind in consume"
)
