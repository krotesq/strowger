package source

type handler struct {
	service *service
}

func newHandler(service *service) *handler {
	return &handler{service: service}
}
