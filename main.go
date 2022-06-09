package main

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
)

func main() {
	c := CustomerRepositoryMock{}
	c.On("GetCustomer", 1).Return("Ath", 25, nil)
	c.On("GetCustomer", 2).Return("", 0, errors.New("not found"))

	//======
	name, age, err := c.GetCustomer(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name, age)

}

type CustomerRepositoryMock struct {
	mock.Mock
}

func (c *CustomerRepositoryMock) GetCustomer(id int) (name string, age int, err error) {
	args := c.Called(id)
	return args.String(0), args.Int(1), args.Error(2)
}
