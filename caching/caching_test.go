// go test -run none -bench . -benchtime 3s

// Tests to show how Data Oriented Design matters.
package caching

import "testing"

// fa 提取为全局变量，是因为go会认为调用的函数没有返回值，编译的时候拿掉这些函数不会对benchmark有影响，
// 所以我们要接收这些函数的返回值赋值给全局变量，这样才能够测试这些函数的benchmark
var fa int

// Capture the time it takes to perform a link list traversal.
func BenchmarkLinkListTraverse(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = LinkedListTraverse()
	}

	fa = a
}

// Capture the time it takes to perform a column traversal.
func BenchmarkColumnTraverse(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = ColumnTraverse()
	}

	fa = a
}

// Capture the time it takes to perform a row traversal.
func BenchmarkRowTraverse(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = RowTraverse()
	}

	fa = a
}
