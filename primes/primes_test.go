package main

import "testing"

var tables = []struct {
	in  int
	out []int
}{
	{1, []int{2}},
	{10, []int{2, 3, 5, 7}},
	{11, []int{2, 3, 5, 7}},
	{500, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43,
		47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103,
		107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163,
		167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227,
		229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281,
		283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353,
		359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421,
		431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487,
		491, 499}},
}

func TestGetPrimes(t *testing.T) {
	for _, table := range tables {
		primes := getPrimes(table.in)
		if len(primes) != len(table.out) {
			t.Fatalf("Result size mismatch. Expected %d but got %d", len(table.out), len(primes))
		}
		for i := range table.out {
			if primes[i] != table.out[i] {
				t.Fatalf("Mismatch at index %d; got %d but expected %d", i, primes[i], table.out[i])
			}
		}
	}
}

func TestGetPrimes2(t *testing.T) {
	for _, table := range tables {
		primes := getPrimes2(table.in)
		if len(primes) != len(table.out) {
			t.Fatalf("Result size mismatch. Expected %d but got %d", len(table.out), len(primes))
		}
		for i := range table.out {
			if primes[i] != table.out[i] {
				t.Fatalf("Mismatch at index %d; got %d but expected %d", i, primes[i], table.out[i])
			}
		}
	}
}

func BenchmarkGetPrimes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getPrimes(500)
	}
}

func BenchmarkGetPrimes2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getPrimes2(500)
	}
}
