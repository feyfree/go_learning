package loop

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}

	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("even")
		case 1, 3:
			t.Log("odd")
		default:
			t.Log("it is not 0 - 3")
		}
	}

}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("even")
		case i%3 == 0:
			t.Log("odd")
		default:
			t.Log("it is not 0 - 3")
		}
	}

	m := 10
	t.Log(m % 4)
}
