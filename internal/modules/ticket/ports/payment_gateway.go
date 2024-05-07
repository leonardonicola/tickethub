package ports

type TicketPaymentGateway[T interface{}, PT interface{}] interface {
	CreateProduct(product T) (PT, error)
	GetProduct(id string) (PT, error)
}
