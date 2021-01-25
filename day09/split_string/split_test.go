package split_string

import (
	"reflect"
	"testing"
)

// 在终端中 go test  -v 即可测试
// 查看测试覆盖率 go test -cover
// 讲覆盖率相关记录输出到一个文件中, go test -cover -coverprofile=cover.out
// 然后执行go tool cover -html=cover.out,
// 测试组

// func TestSplit(t *testing.T) {
// 	// 定义一个测试结构体
// 	type testCase struct {
// 		str  string
// 		sep  string
// 		want []string
// 	}
// 	// 把测试内容存入到结构体

// 	testGroup := []testCase{
// 		testCase{"babcbef", "b", []string{"", "a", "c", "ef"}},
// 		testCase{"cbdef", "b", []string{"c", "def"}},
// 		testCase{"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
// 	}

// 	for _, tc := range testGroup {
// 		got := Spilt(tc.str, tc.sep)
// 		if !reflect.DeepEqual(got, tc.want) {
// 			// 测试用例失败
// 			t.Errorf("want :%v but got : %v\n", tc.want, got)
// 		}
// 	}
// }

// 子测试
func TestSplit(t *testing.T) {
	// 定义一个测试结构体
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	// 把测试内容存入map中
	testGroup := map[string]testCase{
		"case1": testCase{"babcbef", "b", []string{"", "a", "c", "ef"}},
		"case2": testCase{"cbdef", "b", []string{"c", "def"}},
		"case3": testCase{"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				// 测试用例失败
				t.Fatalf("want :%v but got : %v\n", tc.want, got)
			}
		})
	}
}

// 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}

// 性能比较测试
// fib_test.go

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B) {
	benchmarkFib(b, 1)
}

func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib30(b *testing.B) { benchmarkFib(b, 10) }
