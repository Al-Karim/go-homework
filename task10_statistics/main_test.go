package main

import "testing"

func TestComputeStats(t *testing.T) {
	got := computeStats([]int{4, -3, 10, 0, 7})
	want := Stats{Sum: 18, Average: 3.6, Max: 10, Min: -3, Even: 3, Odd: 2}
	if got != want {
		t.Errorf("computeStats = %+v, ожидалось %+v", got, want)
	}
}

func TestComputeStatsSingle(t *testing.T) {
	got := computeStats([]int{-5})
	want := Stats{Sum: -5, Average: -5, Max: -5, Min: -5, Even: 0, Odd: 1}
	if got != want {
		t.Errorf("computeStats = %+v, ожидалось %+v", got, want)
	}
}
