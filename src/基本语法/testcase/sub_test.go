package main

import "testing"

func TestGetsub(t *testing.T) {
	res := getSub(2, 3)

	if res != 5 {
		t.Fatalf("getsub执行错误，期望为%v,实际为%v", 5, res)
	}

	t.Log("getsub执行正确")
}
