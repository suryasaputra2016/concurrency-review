package main

import (
	basicgoroutine "github.com/suryasaputra2016/concurrency-review/1_basic_goroutine"
	waitgoroutine "github.com/suryasaputra2016/concurrency-review/2_wait_goroutine"
	usingchannels "github.com/suryasaputra2016/concurrency-review/3_using_channels"
	bufferedchannels "github.com/suryasaputra2016/concurrency-review/4_buffered_channels"
	selectwithchannels "github.com/suryasaputra2016/concurrency-review/5_select_with_channels"
	handlingerrors "github.com/suryasaputra2016/concurrency-review/6_handling_errors"
)

func main() {
	// 1
	basicgoroutine.BasicGoroutine()

	// 2
	waitgoroutine.WaitGoroutine()

	// 3
	usingchannels.UsingChannels()

	// 4
	bufferedchannels.BufferedChannels()

	// 5
	selectwithchannels.SelectWithChannels()

	// 6
	handlingerrors.HandlingErrors()

}
