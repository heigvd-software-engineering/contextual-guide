package uri

import "main/src/internal/account"

type Model struct {
	id int
	metadata string
	account *account.Model
}

func Create(uri *Model) (bool,*Model){
	return true,nil
}

func Read() (bool, []Model){
	return true,nil
}

func Update(id int, uri *Model) (bool,*Model){
	return true,nil
}

func Delete(id int) bool{
	return true

}


