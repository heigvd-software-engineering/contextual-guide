package httpserver

import (
	"fmt"
	"main/src/internal/storage"
	"net/http"
)


type Router interface {

}

type Service interface {
	GetAll() []Model
	
}

type Repository interface {
	Get() ([]Model,bool)
	Create(model *Model) (*Model,bool)
}

type Storage interface {

}

type Model interface {

}

type Resource struct {
	Name   string
	Model  Model
	Router Router
	Service Service
	Repository Repository
}
type ResourceName = string
type ResourceList = map[ResourceName]*Resource

type Configuration struct {
	port int
}


type httpserver struct {
	resources ResourceList
	storage storage.Storage
	configuration Configuration
}

func (h *httpserver) Register(resource *Resource) {
	h.resources[resource.Name] = resource

	h.updateRouter()
}

func New(port int, store storage.Storage) *httpserver {
	conf := Configuration{
		port: port,
	}

	return &httpserver{
		resources:     nil,
		storage:       store,
		configuration: conf,
	}
}

func (h *httpserver) updateRouter() {
}

func (h *httpserver) start()  {
	_ = http.ListenAndServe(fmt.Sprintf(":%d",h.configuration.port), nil)
}
