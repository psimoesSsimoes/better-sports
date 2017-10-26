package betcoin

import (
	"time"

	log "github.com/sirupsen/logrus"
)

// AccountBalance users account balance
type AccountBalance struct {
	Balance                    float64     `json:"Balance"`
	BlockchainAddress          interface{} `json:"BlockchainAddress"`
	CoineverWithdrawalAddress  interface{} `json:"CoineverWithdrawalAddress"`
	Currency                   string      `json:"Currency"`
	CurrencyCode               string      `json:"CurrencyCode"`
	CurrencySystemAbbreviature interface{} `json:"CurrencySystemAbbreviature"`
	CurrencyWholeNumbersOnly   bool        `json:"CurrencyWholeNumbersOnly"`
	HasActiveBonus             bool        `json:"HasActiveBonus"`
	IsAgent                    bool        `json:"IsAgent"`
	IsNewBitcoinAddress        interface{} `json:"IsNewBitcoinAddress"`
	UserLoyaltyModel           interface{} `json:"UserLoyaltyModel"`
	UserName                   string      `json:"UserName"`
}

// Sport description
type Sport struct {
	HeaderID     int    `json:"HeaderID"`
	ID           int    `json:"ID"`
	IconID       int    `json:"IconID"`
	Key          int    `json:"Key"`
	MatchesCount int    `json:"MatchesCount"`
	Name         string `json:"Name"`
	OriginalName string `json:"OriginalName"`
	Rank         int    `json:"Rank"`
}

// Category sport category
type Category struct {
	HeaderID     interface{} `json:"HeaderID"`
	ID           int         `json:"ID"`
	MatchesCount int         `json:"MatchesCount"`
	Name         string      `json:"Name"`
	Rank         int         `json:"Rank"`
}

// Event sport event
type Event struct {
	AwayPlayerName       interface{} `json:"AwayPlayerName"`
	AwayTeamLogo         interface{} `json:"AwayTeamLogo"`
	AwayTeamName         string      `json:"AwayTeamName"`
	Category             Category    `json:"Category"`
	DateOfMatch          string      `json:"DateOfMatch"`
	DateOfMatchLocalized struct {
		Date                   string `json:"Date"`
		DayName                string `json:"DayName"`
		DayNameShort           string `json:"DayNameShort"`
		MonthNameShort         string `json:"MonthNameShort"`
		NextMonthNameShort     string `json:"NextMonthNameShort"`
		PreviuosMonthNameShort string `json:"PreviuosMonthNameShort"`
		Time                   string `json:"Time"`
		Value                  string `json:"Value"`
	} `json:"DateOfMatchLocalized"`
	HomePlayerName         interface{} `json:"HomePlayerName"`
	HomeTeamLogo           interface{} `json:"HomeTeamLogo"`
	HomeTeamName           string      `json:"HomeTeamName"`
	ID                     int         `json:"ID"`
	IsClockStarted         bool        `json:"IsClockStarted"`
	IsCountDown            bool        `json:"IsCountDown"`
	IsLive                 bool        `json:"IsLive"`
	IsManuallyStarted      bool        `json:"IsManuallyStarted"`
	MarketsCount           int         `json:"MarketsCount"`
	MatchMinute            interface{} `json:"MatchMinute"`
	Name                   string      `json:"Name"`
	OriginalAwayPlayerName interface{} `json:"OriginalAwayPlayerName"`
	OriginalAwayTeamLogo   interface{} `json:"OriginalAwayTeamLogo"`
	OriginalAwayTeamName   string      `json:"OriginalAwayTeamName"`
	OriginalHomePlayerName interface{} `json:"OriginalHomePlayerName"`
	OriginalHomeTeamLogo   interface{} `json:"OriginalHomeTeamLogo"`
	OriginalHomeTeamName   string      `json:"OriginalHomeTeamName"`
	PreviewOdds            []struct {
		AltID           interface{} `json:"AltID"`
		Delta           string      `json:"Delta"`
		ID              int         `json:"ID"`
		IsFeedSuspended bool        `json:"IsFeedSuspended"`
		IsLive          bool        `json:"IsLive"`
		IsSuspended     bool        `json:"IsSuspended"`
		IsVisible       bool        `json:"IsVisible"`
		MarketID        int         `json:"MarketID"`
		MatchID         int         `json:"MatchID"`
		SpecialBetValue interface{} `json:"SpecialBetValue"`
		Status          int         `json:"Status"`
		Title           string      `json:"Title"`
		TitleSuffix     interface{} `json:"TitleSuffix"`
		Value           float64     `json:"Value"`
	} `json:"PreviewOdds"`
	Score              string   `json:"Score"`
	SetScore           string   `json:"SetScore"`
	SportType          Category `json:"SportType"`
	TeamNamesSeparator string   `json:"TeamNamesSeparator"`
	Tournament         struct {
		ID              int    `json:"ID"`
		IsTeamsReversed bool   `json:"IsTeamsReversed"`
		MatchesCount    int    `json:"MatchesCount"`
		Name            string `json:"Name"`
		Rank            int    `json:"Rank"`
		SportCategoryID int    `json:"SportCategoryID"`
	} `json:"Tournament"`
	TvChannels interface{} `json:"TvChannels"`
}

// Events slice of Event
type Events []Event

// Market sports market
type Market struct {
	EventID             int         `json:"EventID"`
	HasAlternativeLines bool        `json:"HasAlternativeLines"`
	ID                  int         `json:"ID"`
	IsActive            bool        `json:"IsActive"`
	IsCashoutEnabled    bool        `json:"IsCashoutEnabled"`
	IsLive              bool        `json:"IsLive"`
	IsOutcome           bool        `json:"IsOutcome"`
	IsSuspended         bool        `json:"IsSuspended"`
	MarketTypeID        int         `json:"MarketTypeID"`
	Name                string      `json:"Name"`
	Rank                int         `json:"Rank"`
	SubType             interface{} `json:"SubType"`
	TypeID              int         `json:"TypeId"`
	ViewType            string      `json:"ViewType"`
}

// MarketData data for market
type MarketData struct {
	AltID           interface{} `json:"AltID"`
	Delta           string      `json:"Delta"`
	ID              int         `json:"ID"`
	IsFeedSuspended bool        `json:"IsFeedSuspended"`
	IsLive          bool        `json:"IsLive"`
	IsSuspended     bool        `json:"IsSuspended"`
	IsVisible       bool        `json:"IsVisible"`
	MarketID        int         `json:"MarketID"`
	MatchID         int         `json:"MatchID"`
	SpecialBetValue interface{} `json:"SpecialBetValue"`
	Status          int         `json:"Status"`
	Title           string      `json:"Title"`
	TitleSuffix     interface{} `json:"TitleSuffix"`
	Value           float64     `json:"Value"`
}

// Markets slice of MarketData
type Markets []MarketData

// SportBetSlip bet slip
type SportBetSlip struct {
	AcceptBetterOdds bool                `json:"AcceptBetterOdds"`
	Selections       []SportBetSelection `json:"Selections"`
}

// SportsBetItem item
type SportsBetItem struct {
	ID       int  `json:"ID"`
	IsBanker bool `json:"IsBanker"`
}

// SportBetSelection bet selection
type SportBetSelection struct {
	IsLay  bool            `json:"IsLay"`
	Items  []SportsBetItem `json:"Items"`
	Return string          `json:"Return"`
	Stake  string          `json:"Stake"`
}

// OutstandingBet struct for uncleared bets
type OutstandingBet struct {
	ID    int `json:"ID"`
	Items []struct {
		AdminID         interface{} `json:"AdminID"`
		BetType         interface{} `json:"BetType"`
		CashoutAmount   int         `json:"CashoutAmount"`
		ID              int         `json:"ID"`
		IsCombo         bool        `json:"IsCombo"`
		IsLay           bool        `json:"IsLay"`
		IsSplitStake    interface{} `json:"IsSplitStake"`
		JackpotID       int         `json:"JackpotID"`
		Message         string      `json:"Message"`
		OfferOddValue   interface{} `json:"OfferOddValue"`
		OfferStakeValue interface{} `json:"OfferStakeValue"`
		OriginalMessage interface{} `json:"OriginalMessage"`
		PossibleReturn  float64     `json:"PossibleReturn"`
		ReturnedAmount  int         `json:"ReturnedAmount"`
		Selections      []struct {
			EndTime string `json:"EndTime"`
			Event   struct {
				AwayTeamName     string      `json:"AwayTeamName"`
				CategoryHeaderID interface{} `json:"CategoryHeaderId"`
				CategoryName     string      `json:"CategoryName"`
				HomeTeamName     string      `json:"HomeTeamName"`
				ID               int         `json:"ID"`
				Name             string      `json:"Name"`
				SportType        struct {
					BetISNKey      interface{} `json:"BetISNKey"`
					DonbestKey     interface{} `json:"DonbestKey"`
					HeaderID       int         `json:"HeaderID"`
					ID             int         `json:"ID"`
					IconID         int         `json:"IconID"`
					IsActive       bool        `json:"IsActive"`
					IsVirtual      bool        `json:"IsVirtual"`
					Key            int         `json:"Key"`
					Name           string      `json:"Name"`
					OriginalName   string      `json:"OriginalName"`
					PinnacleKey    int         `json:"PinnacleKey"`
					Rank           int         `json:"Rank"`
					RunningBallKey interface{} `json:"RunningBallKey"`
					SynchType      int         `json:"SynchType"`
					TXOddsKey      interface{} `json:"TXOddsKey"`
					Timestamp      string      `json:"Timestamp"`
				} `json:"SportType"`
				StartDate      string `json:"StartDate"`
				TournamentName string `json:"TournamentName"`
			} `json:"Event"`
			ID     int `json:"ID"`
			Market struct {
				EventID             int         `json:"EventID"`
				HasAlternativeLines bool        `json:"HasAlternativeLines"`
				ID                  int         `json:"ID"`
				IsActive            bool        `json:"IsActive"`
				IsCashoutEnabled    bool        `json:"IsCashoutEnabled"`
				IsLive              bool        `json:"IsLive"`
				IsOutcome           bool        `json:"IsOutcome"`
				IsSuspended         bool        `json:"IsSuspended"`
				MarketTypeID        int         `json:"MarketTypeID"`
				Name                string      `json:"Name"`
				Rank                int         `json:"Rank"`
				SubType             interface{} `json:"SubType"`
				TypeID              int         `json:"TypeId"`
				ViewType            string      `json:"ViewType"`
			} `json:"Market"`
			Odd struct {
				AltID           interface{} `json:"AltID"`
				Delta           string      `json:"Delta"`
				ID              int         `json:"ID"`
				IsFeedSuspended bool        `json:"IsFeedSuspended"`
				IsLive          bool        `json:"IsLive"`
				IsSuspended     bool        `json:"IsSuspended"`
				IsVisible       bool        `json:"IsVisible"`
				MarketID        int         `json:"MarketID"`
				MatchID         int         `json:"MatchID"`
				SpecialBetValue string      `json:"SpecialBetValue"`
				Status          int         `json:"Status"`
				Title           string      `json:"Title"`
				TitleSuffix     string      `json:"TitleSuffix"`
				Value           float64     `json:"Value"`
			} `json:"Odd"`
			OddValue float64 `json:"OddValue"`
			Status   string  `json:"Status"`
		} `json:"Selections"`
		Stake         float64 `json:"Stake"`
		Status        string  `json:"Status"`
		Timestamp     string  `json:"Timestamp"`
		TotalOddValue float64 `json:"TotalOddValue"`
		TraderStatus  string  `json:"TraderStatus"`
	} `json:"Items"`
	PlacedDate          string  `json:"PlacedDate"`
	Timestamp           string  `json:"Timestamp"`
	TotalPossibleReturn float64 `json:"TotalPossibleReturn"`
	TotalReturnedAmount int     `json:"TotalReturnedAmount"`
	TotalStake          float64 `json:"TotalStake"`
}

// LoginRequest betcoin login request
type LoginRequest struct {
	BrowserFingerPrint int         `json:"BrowserFingerPrint"`
	CaptchaResponse    interface{} `json:"CaptchaResponse"`
	Login              string      `json:"Login"`
	Password           string      `json:"Password"`
	Pin                string      `json:"Pin"`
	RememberMe         bool        `json:"RememberMe"`
	UserName           string      `json:"UserName"`
}

// APIResponse betcoin generic API response
type APIResponse struct {
	CaptchaID       string      `json:"CaptchaId"`
	CaptchaPicture  interface{} `json:"CaptchaPicture"`
	ErrorMessage    string      `json:"ErrorMessage"`
	HasErrors       bool        `json:"HasErrors"`
	MustAcceptTerms bool        `json:"MustAcceptTerms"`
}

// DateFormatString format string
const DateFormatString = "2006-01-02T15:04:05"

func (p Events) Len() int {
	return len(p)
}

func (p Events) Less(i, j int) bool {
	var (
		t1, t2 time.Time
		err    error
	)

	if t1, err = time.Parse(DateFormatString, p[i].DateOfMatch); err != nil {
		log.WithError(err).Errorln("error parsing data")
		return false
	}

	if t2, err = time.Parse(DateFormatString, p[j].DateOfMatch); err != nil {
		log.WithError(err).Errorln("error parsing data")
		return false
	}

	return t1.Before(t2)
}

func (p Events) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Markets) Len() int {
	return len(p)
}

func (p Markets) Less(i, j int) bool {
	return p[i].Value > p[j].Value
}

func (p Markets) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
