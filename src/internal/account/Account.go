package account

import "main/src/pkg/httpserver"

type accountModel struct {
	GoTrueId string
}

func NewModel(id string) httpserver.Model{
	return &accountModel{GoTrueId: id }
}

func (m * accountModel) print(){

}
