package distributor

import "fmt"

type Distributor struct {
	data    *[]string
	urlChan *chan string
}

func (d *Distributor) Target(urlChannel *chan string) *Distributor {
	d.urlChan = urlChannel
	return d
}
func (d *Distributor) NextUrlFunc(f func() string) *Distributor {
	l := make([]string, 0)
	for {
		s := f()
		if "" != s {
			l = append(l, s)
		} else {
			break
		}
	}
	return &Distributor{data: &l}
}
func (d *Distributor) Run() {
	if nil == *d.urlChan {
		panic(fmt.Sprintf("must be set url chan for env before run"))
	}
	if nil == *d.data {
		panic(fmt.Sprintf("must be implement NextUrlFunc func"))
	}
	for i := range *d.data {
		//send data to chan
		*d.urlChan <- (*d.data)[i]
	}
}
