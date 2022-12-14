package goronext

import (
	"bytes"
	"io"
	"net/http"

	"golang.org/x/net/html"
)

type Stock struct {
	ID       string
	Symbol   string
	Name     string
	Exchange string
	Market   string
}

func FindAllStock() (stocks []*Stock, err error) {
	goOn := true
	page := 0
	for goOn {
		url := searchStock("FR001", page)
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		d, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		node, err := html.Parse(bytes.NewReader(d))
		if err != nil {
			return nil, err
		}
		// Find the table with id 'awl-lookup-instruments-directory-table'
		var table *html.Node
		var crawler func(*html.Node)
		crawler = func(node *html.Node) {
			if node.Type == html.ElementNode && node.Data == "tbody" {
				table = node
				return
			}
			for child := node.FirstChild; child != nil; child = child.NextSibling {
				crawler(child)
			}
		}
		crawler(node)
		if table == nil {
			goOn = false
			continue
		}
		// <tr class="odd">
		// 	<td class="priority-medium lookup_instrument_logo"></td>
		// 	<td class="priority-medium">ALMIL</td>
		// 	<td><a href="/en/product/equities/fr0010285965-alxp/1000mercis/almil" target="" data-title-hover="1000MERCIS">1000MERCIS</a></td>
		// 	<td class="priority-medium">FR0010285965</td>
		// 	<td>Euronext Growth Paris</td>
		// 	<td class="priority-low">ALXP</td>
		// 	<td class="priority-low">Stock</td>
		// </tr>
		for child := table.FirstChild.NextSibling; child != nil; child = child.NextSibling.NextSibling {
			stock := &Stock{}
			if child.Type == html.ElementNode && child.Data == "tr" {
				td := child.FirstChild.NextSibling
				if child.FirstChild.NextSibling.Attr[0] == (html.Attribute{Key: "class", Val: "priority-medium lookup_instrument_logo"}) {
					td = td.NextSibling.NextSibling
				}
				stock.Symbol = td.FirstChild.Data
				td = td.NextSibling.NextSibling
				stock.Name = td.FirstChild.FirstChild.Data
				td = td.NextSibling.NextSibling
				stock.ID = td.FirstChild.Data
				td = td.NextSibling.NextSibling
				stock.Exchange = td.FirstChild.Data
				td = td.NextSibling.NextSibling
				stock.Market = td.FirstChild.Data
			}
			for _, registeredStock := range stocks {
				if registeredStock.ID == stock.ID {
					goOn = false
					break
				}
			}
			if !goOn {
				break
			}
			stocks = append(stocks, stock)
		}
		page++
	}
	return stocks, nil
}
