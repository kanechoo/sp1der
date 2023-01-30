package channel

var HttpUrlChannel = make(chan string, 20)
var HtmlDocChannel = make(chan []byte, 10)
