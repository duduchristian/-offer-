package offer

import "testing"

func TestEvalRPN(t *testing.T) {
	res := EvalRPN([]string{"2", "1", "3", "*", "+"})
	if res != 5 {
		t.Fatal("error!")
	}
}
