package account


type Model struct {
	gotrueId string
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


