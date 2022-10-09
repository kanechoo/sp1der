package channel

var UrlChannel = make(chan string, 20)
var HttpExecutorResultChannel = make(chan []byte, 10)
