package goronext

import (
	"testing"
)

func TestCompanyInfo(t *testing.T) {
	info, err := getCompanyInfo("FR0010340141", "XPAR")
	if err != nil {
		t.Error(err)
	}
	if !info.SRD {
		t.Error("SRD should be true")
	}
	if !info.Loan {
		t.Error("Loan should be true")
	}
}
