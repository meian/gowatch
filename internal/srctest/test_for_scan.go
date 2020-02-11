package srctest

//golint:ignore

import (
	"fmt"
	"testing"
)

// their are captured

func TestSimpleTarget1(t *testing.T) {
}

func TestSimpleTarget2(t *testing.T) {
}

func TestSimpleTarget3(t *testing.T) {
}

func TestHasAnounimous(t *testing.T) {
	// this is not captured
	f := func(t *testing.T) int {
		fmt.Print("anounimous func")
		return 1
	}
	fmt.Print(f(t))
}

// their are not captured

// -- no test method
func NoTestMethod(t *testing.T) {
}

// -- private is not test method
func testSimpleTarget4(t *testing.T) {
}

// -- invalid signiture
func TestInvalidSigniture() {

}

type SampleType int

// -- receiver is out of target
func (s SampleType) TestInvalidReceiver(t *testing.T) {

}
