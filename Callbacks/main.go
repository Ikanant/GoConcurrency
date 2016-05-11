package main

import "fmt"

func main() {
	po := new(PurchaseOrder)
	po.Value = 12.51

	ch := make(chan *PurchaseOrder)

	go SavePO(po, ch)

	newPo := <-ch

	fmt.Println("PO: ", newPo)
}

type PurchaseOrder struct {
	Number int
	Value  float64
}

func SavePO(po *PurchaseOrder, callback chan *PurchaseOrder) {
	po.Number = 1234

	callback <- po
}
