package main

import (
	"errors"
	"fmt"
)

func main() {

	po := new(PurchaseOrder)
	po.Value = 45.12

	SavePO(po, false).Then(func(obj interface{}) error {
		po := obj.(*PurchaseOrder)
		fmt.Printf("Purchase Order saved with ID: %d\n", po.Number)

		return nil
	}, func(err error) {
		fmt.Printf("Failed to save Purchase Order: " + err.Error() + "\n")
	})

	fmt.Scanln()

}

type PurchaseOrder struct {
	Number int
	Value  float64
}

func SavePO(po *PurchaseOrder, shouldFail bool) *Promise {
	result := new(Promise)

	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func() {
		if shouldFail {
			result.failureChannel <- errors.New("Failed to save purchase order")
		} else {
			po.Number = 1234
			result.successChannel <- po
		}
	}()

	return result
}

type Promise struct {
	successChannel chan interface{} // Generic promise. So it can pass any type of result data out of itself
	failureChannel chan error
}

func (this *Promise) Then(success func(interface{}) error, failure func(error)) *Promise {

	result := new(Promise)

	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func() {
		select {
		case obj := <-this.successChannel:
			newErr := success(obj)
			if newErr == nil {
				result.successChannel <- obj
			} else {
				result.failureChannel <- newErr
			}
		case err := <-this.failureChannel:
			failure(err)
			result.failureChannel <- err
		}
	}()

	return result

}
