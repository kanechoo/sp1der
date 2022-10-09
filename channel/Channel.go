package channel

var HttpUrl = make(chan string, 20)
var HtmlDoc = make(chan []byte, 10)
