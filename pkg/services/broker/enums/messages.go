package enums

const (
	MessageFailedConnectBroker           = "{ERROR_BROKER} failed to connect"
	MessageFailedCreateChannelPublish    = "{ERROR_BROKER} failed to create channel while publishing"
	MessageFailedDeclareExchangePublish  = "{ERROR_BROKER} failed to declare exchange while publishing"
	MessageFailedCreateChannelConsume    = "{ERROR_BROKER} failed to create channel in consume"
	MessageFailedCreateQueueConsume      = "{ERROR_BROKER} error declaring queue in consumer"
	MessageFailedConsumeHandlingDelivery = "{ERROR_BROKER} consume error while handling deliveries"
	MessageFailedSetConsumerPrefetch     = "{ERROR_BROKER} failed to set consumer prefetch"
	MessageFailedToDeclareExchangeQueue  = "{ERROR_BROKER} failed to declare exchange while declaring queue"
	MessageFailedBindQueueConsume        = "{ERROR_BROKER} failed to queue bind in consume"
)
