package goronext

import "fmt"

var (
	// Endpoint for the API
	companyInfoEndpoint = func(id, market string) string {
		return fmt.Sprintf("https://live.euronext.com/en/ajax/getFactsheetInfoBlock/STOCK/%s-%s/fs_tradinginfo_pea_block", id, market)
	}

	searchStock = func(query string, page int) string {
		return fmt.Sprintf("https://live.euronext.com/en/search_instruments/%s?type=Stock&restMic=&idRest=all&page=%d", query, page)
	}

	intradayEndpoint = func(id, market string) string {
		return fmt.Sprintf("https://live.euronext.com/en/ajax/getIntradayPriceFilteredData/%s-%s", id, market)
	}
)
