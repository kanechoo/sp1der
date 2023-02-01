package channel

var UrlChan = make(chan string, 20)
var DocumentChan = make(chan []byte, 10)
