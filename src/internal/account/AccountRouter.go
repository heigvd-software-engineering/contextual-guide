package account

import "main/src/pkg/httpserver"

type Router struct {
	service httpserver.Service
}

func NewRouter(service *service) *Router {
	return &Router{service: service}
}

func (r *Router) getAll() []httpserver.Model{
	return r.service.GetAll()
}