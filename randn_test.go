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
}

func TestBLookup(t *testing.T) {
	t.Log(BLookup(1))
	t.Log(BLookup(2))
	t.Log(BLookup(3))
	t.Log(BLookup(4))
	t.Log(BLookup(5))
}

func TestALookup(t *testing.T) {
	t.Log(ALookup(1))
	t.Log(ALookup(2))
	t.Log(ALookup(3))
	t.Log(ALookup(4))
	t.Log(ALookup(5))
}

func TestAALookup(t *testing.T) {
	t.Log(AALookup(1))
	t.Log(AALookup(2))
	t.Log(AALookup(3))
	t.Log(AALookup(4))
	t.Log(AALookup(5))
}

func BenchmarkLookup(b *testing.B) {
	// run the Lookup function b.N times
	for i := 0; i < b.N; i++ {
		Lookup(5)
	}
}

func BenchmarkBLookup(b *testing.B) {
	// run the Lookup function b.N times
	for i := 0; i < b.N; i++ {
		BLookup(5)
	}
}

func BenchmarkALookup(b *testing.B) {
	// run the Lookup function b.N times
	for i := 0; i < b.N; i++ {
		ALookup(5)
	}
}

func BenchmarkAALookup(b *testing.B) {
	// run the Lookup function b.N times
	for i := 0; i < b.N; i++ {
		AALookup(5)
	}
}
