package goronext

import (
	"io"
	"net/http"
	"strings"
)

type CompanyInfo struct {
	SRD  bool
	Loan bool
}

func GetCompanyInfo(id string, market string) (*CompanyInfo, error) {
	url := companyInfoEndpoint(id, market)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	d, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := string(d)

	info := &CompanyInfo{}
	if strings.Contains(data, "eligible for the SRD") {
		info.SRD = true
	}
	if strings.Contains(data, "and for the Loan and Lending Market") {
		info.Loan = true
	}
	return info, nil
}
