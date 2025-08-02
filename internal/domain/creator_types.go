package domain

type PaymentMethodCreator func(details map[string]string) (PaymentMethod, error)

type PaymentGatewayCreator func() PaymentGateway
