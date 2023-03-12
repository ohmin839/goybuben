package goybuben

import (
	"github.com/ohmin839/goybuben/goybubenapi/collector"
	"github.com/ohmin839/goybuben/goybubenapi/converter"
)

func ToAybuben(text string) string {
	converter := &converter.Converter{Buffer: text}
	converter.Init()
	converter.Parse()
	return converter.Evaluate()
}

func ToHayerenWords(text string) []string {
	collector := &collector.Collector{Buffer: text}
	collector.Init()
	collector.Parse()
	return collector.Evaluate()
}
