package distributor

import (
	"fmt"
	"time"
)

type Distributor struct {
	urlChan *chan string
}

func (d *Distributor) UrlReceivedChannel(urlChannel *chan string) *Distributor {
	d.urlChan = urlChannel
	return d
}
func (d *Distributor) NextUrlFunc(f func() string) *Distributor {
	if nil == d.urlChan {
		panic("Please defined a url channel for the Distributor struct first then run this function")
	}
	go func() {
		for {
			s := f()
			if "" == s {
				break
			}
			select {
			case *d.urlChan <- s:
				fmt.Printf("Added request url to chan : %s\n", s)
			case <-time.After(10 * time.Second):
				fmt.Printf("Timeout\n")
			}
		}
	}()
	return &Distributor{}
}
