package goronext

import (
	"testing"
	"time"
)

func TestIntraday(t *testing.T) {
	trades, err := GetAllIntraDay("FR0010340141", "XPAR", "100", time.Now())
	if err != nil {
		t.Error(err)
	}
	// fmt.Printf("%#v\n", trades[0])
	if len(trades) == 0 {
		t.Error("No trades found")
	}
}
