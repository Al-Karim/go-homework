package main

import "testing"

func TestIsPrime(t *testing.T) {
	cases := map[int]bool{
		-7: false, 0: false, 1: false,
		2: true, 3: true, 4: false, 5: true, 9: false,
		17: true, 25: false, 97: true, 100: false,
		7919: true, 7921: false,
	}
	for n, want := range cases {
		if got := isPrime(n); got != want {
			t.Errorf("isPrime(%d) = %v, ожидалось %v", n, got, want)
		}
	}
}
