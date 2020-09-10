package test

import (
	"fmt"
	"testing"
)

func TestMap1(t *testing.T) {
	// declare a map variable not create map
	var ranks map[string]int
	// create map and declare a variable to hold it
	ranks = make(map[string]int)
	ranks["gold"] = 1
	ranks["silver"] = 2
	fmt.Println(ranks["gold"])
	fmt.Println(ranks["bronze"])
}

func TestMap2(t *testing.T) {
	myMap := map[string]int{"a": 1, "b": 2, "d": 4}
	emptyMap := map[string]float64{}
	fmt.Println(emptyMap)
	var value int
	var ok bool
	value, ok = myMap["c"]
	a, b := myMap["d"]
	fmt.Println(value, ok)
	fmt.Println(a, b)
	delete(myMap, "a")
	fmt.Println(myMap)
	for _, grade := range myMap {
		fmt.Println(grade)
	}

}
