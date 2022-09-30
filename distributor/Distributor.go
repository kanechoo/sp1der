package distributor

import "fmt"

type QueryParams struct {
	Url string
}
type Distributor[T any] struct {
	data    *[]T
	urlChan *chan T
}
type NextQueryParamsFunc[T any] func() (bool, T)

func (d *Distributor[T]) Target(urlChannel *chan T) *Distributor[T] {
	d.urlChan = urlChannel
	return d
}
func (d *Distributor[T]) NextQueryParams(paramsFunc NextQueryParamsFunc[T]) *Distributor[T] {
	l := make([]T, 0)
	for {
		ok, p := paramsFunc()
		if ok {
			l = append(l, p)
		} else {
			break
		}
	}
	return &Distributor[T]{data: &l}
}
func (d *Distributor[T]) Run() {
	if nil == *d.urlChan {
		panic(fmt.Sprintf("must be set url chan for env before run"))
	}
	if nil == *d.data {
		panic(fmt.Sprintf("must be implement NextQueryParams func"))
	}
	for i := range *d.data {
		//send data to chan
		*d.urlChan <- (*d.data)[i]
	}
}
