package exercises

import (
	"math/rand"
	"time"
)

// assume that 10 goroutines are calling Buggyloop function
// and user can call close function whenever he/she wants
// what are the bugs ?

type coinInfo struct {
	name  string
	price int
}

type coinFetcher struct {
	updates chan coinInfo
	closed  bool
	err     error
}

func (f *coinFetcher) Buggyloop() {
	for {
		if f.closed {
			close(f.updates)
			return
		}
		CoinInfo, nextTimeToRead, err := f.fetch()
		if err != nil {
			f.err = err
			time.Sleep(10 * time.Second)
			continue
		}
		for _, item := range CoinInfo {
			f.updates <- item
		}
		if now := time.Now(); nextTimeToRead.After(now) {
			time.Sleep(nextTimeToRead.Sub(now))
		}
	}
}

func (f *coinFetcher) fetch() ([]coinInfo, time.Time, error) {
	return []coinInfo{
		{name: "Bitcoin", price: 20000 + rand.Intn(1000)},
		{name: "Etherium", price: 1000 + rand.Intn(100)},
		{name: "USDT", price: 1},
	}, time.Now().Add(time.Hour) , nil
}

func (f *coinFetcher) close() error {
	f.closed = true
	return f.err
}
