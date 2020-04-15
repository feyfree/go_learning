package pipefilter

import "fmt"

// NewStraightPipeline create a new StraightPipelineWithWallTime
func NewStraightPipeline(name string, filters ...Filter) *StraightPipeline {
	return &StraightPipeline{
		Name:    name,
		Filters: &filters,
	}
}

// StraightPipeline is composed of the filters, and the filters are piled as a straigt line.
type StraightPipeline struct {
	Name    string
	Filters *[]Filter
}

// Process is to process the coming data by the pipeline
func (f *StraightPipeline) Process(data Request) (Response, error) {
	var ret interface{}
	var err error
	fmt.Println(f.Filters)
	// 注意数据类型   *f.Filters 计算顺序是 *(f.Filters)
	for _, filter := range *f.Filters {
		ret, err = filter.Process(data)
		if err != nil {
			return ret, err
		}
		data = ret
	}
	return ret, err
}
