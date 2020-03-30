package client

import "testing"
import "ch15/series"

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSeries(5))
}
