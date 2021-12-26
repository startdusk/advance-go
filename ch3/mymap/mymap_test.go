package mymap

import "testing"

func BenchmarkMayMapLoad(b *testing.B) {
	n := b.N
	m := NewMyMap()
	m.Store(n, n)
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			m.Load(n)
		}
	})
}

func BenchmarkMayMapStore(b *testing.B) {
	m := NewMyMap()
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			m.Store(b.N, b.N)
		}
	})
}

func BenchmarkMayMapDelete(b *testing.B) {
	m := NewMyMap()
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			m.Delete(b.N)
		}
	})
}

func BenchmarkMayMapLoadOrStore(b *testing.B) {
	m := NewMyMap()
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			m.LoadOrStore(b.N, b.N)
		}
	})
}

func BenchmarkMayMapLoadAndDelete(b *testing.B) {
	m := NewMyMap()
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			m.LoadAndDelete(b.N)
		}
	})
}
