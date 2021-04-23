package test

import (
	"errors"
)

func Test() error  {

	if err:= Test1();err!=nil {
		return err
	}

	return nil
}


func Test1() error  {

	return errors.New("EOF")

}