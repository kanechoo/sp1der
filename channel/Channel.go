package channel

import "sp1der/distributor"

var UrlChannel = make(chan distributor.QueryParams, 20)
var ExecutorResultChannel = make(chan []byte, 10)
