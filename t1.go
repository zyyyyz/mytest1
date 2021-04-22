package git_te2

import (
	"fmt"
)

func testInner1() {
	fmt.Println("内部包1")
}

func testInner2() {
	fmt.Println("内部包2")
}

func TestOut1() {
	fmt.Println("内部包2")
}

func TestOut2() {
	testInner2()
	fmt.Println("外部包2")
}