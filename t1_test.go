package mytest1_test

import (
	"testing"
	tool "github.com/zyyyyz/mytest1"
)


// go test  单元测试
// go test -v 测试所有的测试项
// go test -run 函数 -v  函数如TestFun1, go test -run TestFun1 -v

func TestFun1(t *testing.T) {
	// tool.testInner1()
	t.Log("测试一")
	tool.TestOut1()

}

func TestFun2(t *testing.T) {
	// tool.testInner1()

	t.Log("测试二")
	tool.TestOut2()

}