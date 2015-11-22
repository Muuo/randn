package randn

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestLookup(t *testing.T) {
	t.Log(Lookup(1))
	t.Log(Lookup(2))
	t.Log(Lookup(3))
	t.Log(Lookup(4))
	t.Log(Lookup(5))
	t.Log(Lookup(6))
	t.Log(Lookup(7))
	t.Log(Lookup(8))
	t.Log(Lookup(9))
	t.Log(Lookup(10))
}

func TestBLookup(t *testing.T) {
	t.Log(Lookup(1))
	t.Log(Lookup(2))
	t.Log(Lookup(3))
	t.Log(Lookup(4))
	t.Log(Lookup(5))
	t.Log(Lookup(6))
	t.Log(Lookup(7))
	t.Log(Lookup(8))
	t.Log(Lookup(9))
	t.Log(Lookup(10))
}

func TestALookup(t *testing.T) {
	t.Log(Lookup(1))
	t.Log(Lookup(2))
	t.Log(Lookup(3))
	t.Log(Lookup(4))
	t.Log(Lookup(5))
	t.Log(Lookup(6))
	t.Log(Lookup(7))
	t.Log(Lookup(8))
	t.Log(Lookup(9))
	t.Log(Lookup(10))
}

func TestAALookup(t *testing.T) {
	t.Log(Lookup(1))
	t.Log(Lookup(2))
	t.Log(Lookup(3))
	t.Log(Lookup(4))
	t.Log(Lookup(5))
	t.Log(Lookup(6))
	t.Log(Lookup(7))
	t.Log(Lookup(8))
	t.Log(Lookup(9))
	t.Log(Lookup(10))
}

func BenchmarkLookup(b *testing.B) {
	// run the Lookup function b.N times
	for i := 0; i < b.N; i++ {
		Lookup(10)
	}
}

func BenchmarkBLookup(b *testing.B) {
	// run the Lookup function b.N times
	for i := 0; i < b.N; i++ {
		BLookup(10)
	}
}

func BenchmarkALookup(b *testing.B) {
	// run the Lookup function b.N times
	for i := 0; i < b.N; i++ {
		ALookup(10)
	}
}

func BenchmarkAALookup(b *testing.B) {
	// run the Lookup function b.N times
	for i := 0; i < b.N; i++ {
		AALookup(10)
	}
}
