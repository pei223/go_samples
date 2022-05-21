package main

import (
	"go_samples/array_sample"
	"go_samples/goroutine_sample"
	"go_samples/interface_sample"
	"go_samples/struct_sample"
)

func main() {
	struct_sample.DoStructSamples()
	struct_sample.DoStructExtendSamples()
	array_sample.DoArraySample()
	interface_sample.DoInterfaceSample()
	goroutine_sample.DoGoroutineSample()
}
