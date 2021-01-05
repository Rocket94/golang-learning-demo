package UnitTest

import (
	"bytes"
	"testing"
)
//Benchmark测试通常用于对代码片段的性能测试和对第三方库的性能进行评估测试。
//Benchmark性能测试用例也是测试代码，所以它应该写在以_test结尾（xxx_test.go）的测试用例代码文件中。
func BenchmarkConcatStringByAdd(b *testing.B) {

	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer

		for _, elem := range elems {
			buf.WriteString(elem)

		}
	}
	b.StopTimer()

}