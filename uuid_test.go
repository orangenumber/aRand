package arand_test

import (
	"github.com/orangenumber/arand"
	"testing"
)

func TestUUID(t *testing.T) {
	totalTest := 1_000

	m := make(map[string]int)

	for i := 0; i < totalTest; i++ {
		a := string(arand.UUID())
		if len(a) != 36 {
			t.Fail()
		}
		if v, ok := m[a]; ok {
			m[a] = v + 1
		} else {
			m[a] = 1
		}
	}

	collide := 0
	for _, v := range m {
		if v > 1 {
			collide += v - 1
		}
	}

	if collide > 0 { // 0.1%
		t.Logf("UUID.collide: %.1f", float32(collide)/float32(totalTest)*100)
		t.Fail()
	}
}
