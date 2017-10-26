package main

import (
	"better-sports/betcoin"
	"sort"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	var (
		bc                  betcoin.Betcoin
		balance             betcoin.AccountBalance
		sports              []betcoin.Sport
		events, eventsFinal betcoin.Events
		lowMarkets          []betcoin.MarketData
		err                 error
	)

	bc = betcoin.New("", "")

	log.Println("[better-sports] getting all sports")
	if sports, err = bc.GetAllSports(); err != nil {
		log.WithError(err).Errorln("error getting sports listing")
		return
	}

	for {
		// Wait if we come from an error
		if err != nil {
			time.Sleep(30 * time.Minute)
		}

		// Get account balance
		if err = bc.Login(); err != nil {
			log.WithError(err).Errorln("error logging in")
			continue
		}

		log.Println("[better-sports] getting account balance")
		if balance, err = bc.GetBalance(); err != nil {
			log.WithError(err).Errorln("error getting balance")
			continue
		}
		log.Println("[better-sports] account balance is ", balance.Balance)

		// // Check if any ongoing bets
		// if _, err = bc.GetOutstanding(); err != nil {
		// 	log.WithError(err).Errorln("error getting outstanding bets")
		// 	continue
		// }

		log.Println("[better-sports] getting all events")
		for _, s := range sports {
			if events, err = bc.GetAllEvents(s.ID); err != nil {
				log.WithError(err).Errorln("error getting events")
				continue
			}

			eventsFinal = append(eventsFinal, events...)
		}

		//Sort the map by date
		sort.Sort(eventsFinal)

		var queriableEvents betcoin.Events
		for _, e := range eventsFinal {
			var t1 time.Time
			if t1, err = time.Parse(betcoin.DateFormatString, e.DateOfMatch); err != nil {
				log.WithError(err).Errorln("error parsing data")
				continue
			}

			if t1.Sub(time.Now()).Hours() < 2 {
				queriableEvents = append(queriableEvents, e)
			}
		}

		log.Printf("[better-sports] getting markets of %d events", len(queriableEvents))
		for _, e := range queriableEvents {
			var lowMarketsTemp []betcoin.MarketData
			if lowMarketsTemp, err = bc.GetEventMarkets(e.ID); err != nil {
				log.WithError(err).Errorln("error getting markets")
				continue
			}

			for _, ee := range lowMarketsTemp {
				has := false
				for _, eee := range lowMarkets {
					if ee.MatchID == eee.MatchID {
						has = true
					}
				}

				if !has {
					lowMarkets = append(lowMarkets, ee)
				}
			}
		}

		// Sort by odd value
		var lowMarketsSorted betcoin.Markets
		lowMarketsSorted = append(lowMarketsSorted, lowMarkets...)
		sort.Sort(lowMarketsSorted)

		// Place bet
		if err = bc.Login(); err != nil {
			log.WithError(err).Errorln("error logging in")
			continue
		}

		log.Printf("[better-sports] placing bet on %d markets", len(lowMarketsSorted))
		if _, err = bc.PlaceBets(lowMarketsSorted); err != nil {
			log.WithError(err).Errorln("error logging in")
			continue
		}

		// Wait for the next turn
		time.Sleep(2 * time.Hour)
	}
}
