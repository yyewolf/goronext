package goronext

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Trade struct {
	Price     int
	Time      *time.Time
	TradeId   string
	TradeType string
	Volume    float64
}

type RawData struct {
	Rows []struct {
		Tprice    string `json:"price"`
		Ttime     string `json:"time"`
		Ttrade_id string `json:"tradeId"`
		Ttype_    string `json:"type"`
		Tvolume   string `json:"volume"`
	} `json:"rows"`
}

func ParsePrice(price string) (int, error) {
	price = strings.ReplaceAll(price, ",", "")
	p := strings.Split(price, ".")
	whole := p[0]
	decimal := p[1]

	w, err := strconv.Atoi(whole)
	if err != nil {
		return 0, err
	}
	d, err := strconv.Atoi(decimal)
	if err != nil {
		return 0, err
	}

	// 1.2345 -> 12345, cut at 4 digits after the comma
	for i := len(decimal); i < 4; i++ {
		d *= 10
	}
	for i := len(decimal); i > 4; i-- {
		d /= 10
	}
	return w*10000 + d, nil
}

func GetAllIntraDay(id, market, amount string, when time.Time) ([]*Trade, error) {
	url := intradayEndpoint(id, market)

	today := when.Format("2006-02-01")
	body := strings.NewReader(fmt.Sprintf("nbitems=%s&date=%s&timezone=CET", amount, today))

	resp, err := http.Post(url, "application/x-www-form-urlencoded", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	d, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var rd RawData
	err = json.Unmarshal(d, &rd)
	if err != nil {
		return nil, err
	}
	var trades []*Trade
	for _, row := range rd.Rows {
		t, err := time.Parse("15:04:05", row.Ttime)
		if err != nil {
			return nil, err
		}
		n := time.Now()
		t = t.AddDate(n.Year(), int(n.Month())-1, n.Day()) // add 1 day to get the correct date

		row.Tvolume = strings.ReplaceAll(row.Tvolume, ",", "")
		vol, err := strconv.ParseFloat(row.Tvolume, 64)
		if err != nil {
			return nil, err
		}

		price, err := ParsePrice(row.Tprice)
		if err != nil {
			return nil, err
		}

		trades = append(trades, &Trade{
			Price:     price,
			Time:      &t,
			TradeId:   row.Ttrade_id,
			TradeType: row.Ttype_,
			Volume:    vol,
		})
	}
	return trades, nil
}
