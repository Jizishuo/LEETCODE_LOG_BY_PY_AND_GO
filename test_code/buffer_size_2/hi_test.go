package main

import "testing"

func BencharkPipelineFan(b *testing.B)  {
	for i:=0;i<b.N;i++{
		PipelineFan(0)
	}
}
func BenchmarkPipelineSimple(b *testing.B)  {
	for i:=0;i<b.N;i++{
		PipelineSimple()
	}
}