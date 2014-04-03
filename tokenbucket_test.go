// tokenbucket_test.go (c) 2014 David Rook - all rights reserved

package tokenbucket

import (
	"fmt"
	"testing"
	"time"
)

func Test_01(t *testing.T) {
	tb := New(time.Millisecond, 100)
	tb.Dump()
	delay := tb.Take(15)
	fmt.Printf("Take(15) rv = %v\n", delay)
	tb.Dump()
	delay = tb.Take(5)
	fmt.Printf("Take(5) rv = %v\n", delay)
	fmt.Printf("Waiting - should be 17 sec\n")
	fmt.Printf("Wait(7) + 10 from previous takes\n")
	tb.Wait(7)
	if false {
		fmt.Printf("Test_01\n")
		t.Fatalf("failed")
	}

}
