package UnitTest

import (
	"fmt"
	"testing"
)

//注意测试go test +filename 不仅要添加测试文件名，还有被测文件及其相关引用到的文件！！！！！

//测试前
func setup() {
	fmt.Println("Before all tests")
}
//测试后
func teardown() {
	fmt.Println("After all tests")
}

//普通测试 go test -v(显示过程) -cover(显示覆盖率)
func TestAdd(t *testing.T) {
	if ans:=Add(1,2);ans!=3{
		t.Errorf("1+2 execpted be 3,but %d got",ans)
	}
	if ans:=Add(-1,-2);ans!=-3{
		t.Errorf("-1+-2 execpted be -3,but %d got",ans)
	}
}

//子测试,用t.run来测试
func TestMul(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		if Mul(2, 3) != 6 {
			//t.Fatal和t.Error的区别是前者会终止测试，而后者不会
			t.Fatal("fail")
		}

	})
	t.Run("neg", func(t *testing.T) {
		if Mul(2, -3) != -6 {
			t.Fatal("fail")
		}
	})
}
//用case规范测试，增加可读性
func TestMul2(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"pos", 2, 3, 6},
		{"neg", 2, -3, -6},
		{"zero", 2, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := Mul(c.A, c.B); ans != c.Expected {
				t.Fatalf("%d * %d expected %d, but %d got",
					c.A, c.B, c.Expected, ans)
			}
		})
	}
}
//用helper函数，对一些重复的逻辑，抽取出来作为公共的帮助函数(helpers)，可以增加测试代码的可读性和可维护性
type calcCase struct{ A, B, Expected int }

func createMulTestCase(t *testing.T, c *calcCase) {
	//帮助定位调用函数的位置
	 t.Helper()
	if ans := Mul(c.A, c.B); ans != c.Expected {
		t.Fatalf("%d * %d expected %d, but %d got",
			c.A, c.B, c.Expected, ans)
	}
}

func TestMul3(t *testing.T) {
	createMulTestCase(t, &calcCase{2, 3, 6})
	createMulTestCase(t, &calcCase{2, -3, -6})
	createMulTestCase(t, &calcCase{2, 0, 0}) // wrong case
}
