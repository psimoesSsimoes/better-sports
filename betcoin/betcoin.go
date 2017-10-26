package betcoin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	baseURL         = "https://sports.betcoin.ag/api"
	authEndpoint    = "/authorization/post"
	balanceEndpoint = "/account"
	historyEndpoint = "/profile/history/GetSportHistory"
	sportsEndpoint  = "/sporttype/getallactive"
	catEndpoint     = "/sportcategory/GetBySportType"
	getEndpoint     = "/sportmatch/Get"
	marketEndpoint  = "/sportmarket/getbymatchandmarkettype"
	oddEndpoint     = "/sportodd/GetByMarkets"
)

// Betcoin main struct
type Betcoin struct {
	User   string
	Pass   string
	Stake  float64
	Client *http.Client
}

// New returns a new betcoin
func New(u, p string) Betcoin {
	// proxyURL, _ := url.Parse("http://localhost:8080")
	return Betcoin{
		User:  u,
		Pass:  p,
		Stake: 0.8,
		Client: &http.Client{
			Timeout: time.Duration(30 * time.Second),
			Transport: &http.Transport{
				// Proxy:               http.ProxyURL(proxyURL),
				DisableKeepAlives:   true,
				MaxIdleConns:        2,
				MaxIdleConnsPerHost: 2,
			},
		},
	}
}

// GetBalance user balance
func (b *Betcoin) GetBalance() (AccountBalance, error) {
	var (
		bodyBytes []byte
		data      AccountBalance
		err       error
	)

	link := fmt.Sprintf("%s%s", baseURL, balanceEndpoint)

	if bodyBytes, err = b.Get(link); err != nil {
		return AccountBalance{}, err
	}

	if err = json.Unmarshal(bodyBytes, &data); err != nil {
		return AccountBalance{}, err
	}

	return data, nil
}

// GetOutstanding returns the ids of uncleared bets
func (b *Betcoin) GetOutstanding() ([]int, error) {
	var (
		bodyBytes []byte
		data      []OutstandingBet
		dataOut   []int
		err       error
	)

	// Format the link
	now := time.Now().Format(DateFormatString)
	then := time.Now().AddDate(0, -1, 0).Format(DateFormatString)
	link := fmt.Sprintf("%s%s?from=%s&to=%s&type=UnclearedBets", baseURL, historyEndpoint, then, now)

	if bodyBytes, err = b.Post(link, "", false); err != nil {
		return []int{}, err
	}

	fmt.Println(string(bodyBytes))

	if err = json.Unmarshal(bodyBytes, &data); err != nil {
		return []int{}, err
	}

	for _, d := range data {
		dataOut = append(dataOut, d.ID)
	}

	return dataOut, nil
}

// GetAllSports all sports
func (b *Betcoin) GetAllSports() ([]Sport, error) {
	var (
		bodyBytes []byte
		data      []Sport
		err       error
	)

	link := fmt.Sprintf("%s%s", baseURL, sportsEndpoint)

	if bodyBytes, err = b.Get(link); err != nil {
		return []Sport{}, err
	}

	if err = json.Unmarshal(bodyBytes, &data); err != nil {
		return []Sport{}, err
	}

	return data, nil
}

// GetAllEvents retrieves all events
func (b *Betcoin) GetAllEvents(id int) ([]Event, error) {
	var (
		bodyBytes            []byte
		dataCategory         []Category
		dataEvent, dataFinal []Event
		err                  error
	)

	linkCat := fmt.Sprintf("%s%s?sportTypeID=%d", baseURL, catEndpoint, id)

	// Create request
	if bodyBytes, err = b.Get(linkCat); err != nil {
		return []Event{}, err
	}

	if err = json.Unmarshal(bodyBytes, &dataCategory); err != nil {
		return []Event{}, err
	}

	for _, c := range dataCategory {
		linkEvt := fmt.Sprintf("%s%s?categoryID=%d&sportID=%d", baseURL, getEndpoint, c.ID, id)

		// Create request
		if bodyBytes, err = b.Get(linkEvt); err != nil {
			log.WithError(err).Errorln("error executing request")
			continue
		}

		if err = json.Unmarshal(bodyBytes, &dataEvent); err != nil {
			log.WithError(err).Errorln("error unmashalling data")
			continue
		}

		dataFinal = append(dataFinal, dataEvent...)
	}

	return dataFinal, nil
}

// GetEventMarkets returns all events for a market
func (b *Betcoin) GetEventMarkets(id int) ([]MarketData, error) {
	var (
		bodyBytes  []byte
		data       []Market
		dataMarket []MarketData
		err        error
	)

	link := fmt.Sprintf("%s%s?id=%d&isLive=false&sportMarketTypeID=1", baseURL, marketEndpoint, id)

	// Create request
	if bodyBytes, err = b.Get(link); err != nil {
		return []MarketData{}, err
	}

	if err = json.Unmarshal(bodyBytes, &data); err != nil {
		return []MarketData{}, err
	}

	var idSlice []string
	for _, d := range data {
		if d.IsActive {
			idSlice = append(idSlice, strconv.Itoa(d.ID))
		}
	}
	linkOdd := fmt.Sprintf("%s%s?isLive=false&marketIDs=%%5B%s%%5D", baseURL, oddEndpoint, strings.Join(idSlice, ","))

	// Create request
	if bodyBytes, err = b.Post(linkOdd, "", true); err != nil {
		return []MarketData{}, err
	}

	if err = json.Unmarshal(bodyBytes, &dataMarket); err != nil {
		return []MarketData{}, err
	}

	// TODO o cliente não decide que valores a retornar ou não
	var markets []MarketData
	for _, d := range dataMarket {
		if d.Value < 1.09 {
			markets = append(markets, d)
		}
	}

	return markets, nil
}

// Login logs in
func (b *Betcoin) Login() error {
	var (
		login              LoginRequest
		loginResponse      APIResponse
		bodyBytes, loginBy []byte
		err                error
	)

	// Set variables
	login.Login = b.User
	login.Password = b.Pass
	login.RememberMe = false
	login.BrowserFingerPrint = int(time.Now().UnixNano())
	link := fmt.Sprintf("%s%s", baseURL, authEndpoint)

	if loginBy, err = json.Marshal(login); err != nil {
		log.WithError(err).Errorln("error encoding login message")
		return err
	}

	// Create request
	if bodyBytes, err = b.Post(link, string(loginBy), true); err != nil {
		log.WithError(err).Errorln("error creating http request")
		return err
	}

	if err = json.Unmarshal(bodyBytes, &loginResponse); err != nil {
		log.WithError(err).Errorln("error unmarshalling body")
		return err
	}

	if loginResponse.HasErrors {
		log.Errorln("error from API ", loginResponse.ErrorMessage)
		return err
	}

	return err
}

// PlaceBets place the actual bets
func (b *Betcoin) PlaceBets(markets []MarketData) ([]int, error) {
	var (
		bodyBytes []byte
		items     []SportsBetItem
		response  APIResponse
		err       error
	)

	i := 0
	bulk := b.Stake
	for _, m := range markets {
		var item SportsBetItem

		bulk = bulk * m.Value

		item.ID = m.ID
		item.IsBanker = false
		items = append(items, item)

		if i == 8 {
			var betslip SportBetSlip
			var selection []SportBetSelection

			selection = make([]SportBetSelection, 1)
			selection[0].IsLay = false
			selection[0].Stake = fmt.Sprintf("%.2f", b.Stake)
			selection[0].Return = fmt.Sprintf("%.2f", bulk)
			selection[0].Items = items

			betslip.AcceptBetterOdds = true
			betslip.Selections = selection

			var betslipBy []byte
			if betslipBy, err = json.Marshal(betslip); err != nil {
				return []int{}, err
			}

			// Create request
			if bodyBytes, err = b.Post("http://sports.betcoin.ag/api/betslip/place", string(betslipBy), false); err != nil {
				return []int{}, err
			}

			if err = json.Unmarshal(bodyBytes, &response); err != nil {
				return []int{}, err
			}

			log.WithFields(log.Fields{
				"HasErrors": response.HasErrors,
				"Message":   response.ErrorMessage,
				"raw":       string(bodyBytes),
			}).Println("[better-sports] bet placed successfully")

			// Clear stuff
			items = make([]SportsBetItem, 0)
			i = 0
		}
		i++
	}

	return []int{}, err
}

// Get helper method to perform http gets
func (b *Betcoin) Get(link string) ([]byte, error) {
	var (
		req       *http.Request
		resp      *http.Response
		bodyBytes []byte
		err       error
	)

	// Create request
	if req, err = http.NewRequest("GET", link, nil); err != nil {
		return []byte{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:55.0) Gecko/20100101 Firefox/55.0")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "close")

	// Execute request
	if resp, err = b.Client.Do(req); err != nil {
		return []byte{}, err
	}

	if bodyBytes, err = ioutil.ReadAll(resp.Body); err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	return bodyBytes, err
}

// Post helper method to perform http posts
func (b *Betcoin) Post(link, data string, clear bool) ([]byte, error) {
	var (
		req       *http.Request
		resp      *http.Response
		jar       *cookiejar.Jar
		bodyBytes []byte
		err       error
	)

	if clear {
		// Clear the cookies
		if jar, err = cookiejar.New(nil); err != nil {
			return []byte{}, err
		}
		b.Client.Jar = jar
	}

	if req, err = http.NewRequest("POST", link, strings.NewReader(data)); err != nil {
		return []byte{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:55.0) Gecko/20100101 Firefox/55.0")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "close")

	// Execute request
	if resp, err = b.Client.Do(req); err != nil {
		return []byte{}, err
	}

	if bodyBytes, err = ioutil.ReadAll(resp.Body); err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	return bodyBytes, err
}
