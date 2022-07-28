package exercises

import (
	"math/rand"
	"time"
)

// 1 _ f.err and f.closed have data race
// now we changed it and we use req/resp mechanism to nitoce the error safely
// 2 _ loop may blocks forever when reading from updates (we cannot ensure that user is subcribing)
// when no one 
// 3 _ the close may be called when goroutine is sleeping on time.sleep(nextTimeToRead)
// its in the select statment and will be called whenever the user wants it
// here is another approach we work with stack of each goroutine and we share it whenever we need
// so we re using the advantages of GOLANG CSP MODEL(search about it. its interesting !).
var maxPendings = 10

type fixedCoinInfo struct {
	name  string
	price int
}

type fixedcoinFetcher struct {
	updates chan fixedCoinInfo
	closing chan chan error
}

func (f *fixedcoinFetcher) FixedLoop() {
	for {
		var Coins []fixedCoinInfo
		var pending []fixedCoinInfo
		var first fixedCoinInfo
		var updates chan fixedCoinInfo
		var NextFetch time.Time
		var err error
		for {
			var StartFetch <-chan time.Time 
			// if we have too many pendings then StartFetch is nil 
			// and it will never get selected in the select statement
			// so we wont fetch more and just update the feed
			if len(pending) < maxPendings { 				
				StartFetch = time.After(time.Until(NextFetch))
			}
			StartFetch = time.After(time.Until(NextFetch))
			if len(pending) > 0 {
				first = pending[0]
				updates = f.updates
			}
			select {
			case updates <- first:
				pending = pending[1:]
			case errChan := <- f.closing:
				errChan <- err
				close(f.updates)
				return
			case <-StartFetch:
				Coins, NextFetch, err = f.fetch()
				if err != nil {
					NextFetch = time.Now().Add(10 * time.Second)
					break
				}
				
				pending = append(pending, Coins...)

			}
		}
	}
}

func (f *fixedcoinFetcher) fetch() ([]fixedCoinInfo, time.Time, error) {
	return []fixedCoinInfo{
		{name: "Bitcoin", price: 20000 + rand.Intn(1000)},
		{name: "Etherium", price: 1000 + rand.Intn(100)},
		{name: "USDT", price: 1},
	}, time.Now().Add(time.Hour), nil
}

func (f *fixedcoinFetcher) FixedClose() error {
	errch := make(chan error)
	f.closing <- errch
	return <-errch
}
