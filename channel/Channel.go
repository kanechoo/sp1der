package channel

var UrlChannel = make(chan string, 20)
var ExecutorResultChannel = make(chan []byte, 10)
