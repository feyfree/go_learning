package array

import (
	"fmt"
	"testing"
	"time"
)

func TestArrayIndex(t *testing.T) {
	var notes [7]string
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[0])
	fmt.Println(notes[1])
	fmt.Println(notes[4])
}

func TestArray(t *testing.T) {
	var dates [3]time.Time
	dates[0] = time.Unix(1257894000, 0)
	dates[1] = time.Unix(1447920000, 0)
	fmt.Println(dates[1])
}

func TestArrayPrint(t *testing.T) {
	var notes = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes[3], notes[5])
	//
	newNotes := [3]string{"a", "b", "c"}
	fmt.Println(newNotes[0])
	fmt.Println(newNotes)
	fmt.Printf("%#v", newNotes)

}

func TestArrayLoop(t *testing.T) {
	var notes = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}

	for index, note := range notes {
		fmt.Println(index, note)
	}
}
