package arand_test

import (
	"github.com/orangenumber/arand"
	"testing"
)

func TestRandStr(t *testing.T) {
	numTest := 2000
	var acceptanceRate float32 = 0.999
	length := 5
	// int
	{
		out := make(map[string]int)
		for i := 0; i < numTest; i++ {
			r := arand.RandStr(arand.R_ALPHANUMERIC, length)
			if v, ok := out[r]; ok {
				out[r] = v + 1
			} else {
				out[r] = 1
			}
		}
		if float32(len(out)) < float32(numTest)*acceptanceRate {
			t.Logf("RandInt.collideRate: %.2f%% (%d/%d)", float32(numTest-len(out))/float32(numTest)*100, numTest-len(out), numTest)
			t.Fail()
		}
	}
}

func TestRandInt(t *testing.T) {
	numTest := 20
	numRange := 500
	// int
	{
		out := make(map[int]int)
		for i := 0; i < numTest; i++ {
			r := arand.RandInt(0, numRange)
			if v, ok := out[r]; ok {
				out[r] = v + 1
			} else {
				out[r] = 1
			}
		}

		total := 0
		for _, v := range out {
			if v > 1 {
				total += v - 1
			}
		}

		if collide := total * 100 / numTest; collide > 5 {
			t.Logf("RandInt.collideRate: %d%% (%d/%d)", collide, total, numTest)
			t.Fail()
		}
	}

	// i64
	{
		out := make(map[int64]int)
		for i := 0; i < numTest; i++ {
			r := arand.RandInt64(0, int64(numRange))
			if v, ok := out[r]; ok {
				out[r] = v + 1
			} else {
				out[r] = 1
			}
		}
		total := 0
		for _, v := range out {
			if v > 1 {
				total += v - 1
			}
		}
		if collide := total * 100 / numTest; collide > 5 {
			t.Logf("RandInt64.collideRate: %d%% (%d/%d)", collide, total, numTest)
			t.Fail()
		}
	}
}
