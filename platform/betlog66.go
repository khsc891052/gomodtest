package platform

import (
	"encoding/json"
	"github.com/mikemintang/go-curl"
	"time"
)

type betlog66 struct {
	Status    string `json:"status"`
	ErrorCode string `json:"errorCode"`
	Data      struct {
		Data []struct {
			WagersID       int    `json:"WagersID"`
			UserID         int    `json:"UserID"`
			BetAmount      string `json:"BetAmount"`
			Currency       string `json:"Currency"`
			Result         int    `json:"Result"`
			GameType       int    `json:"GameType"`
			RoundSerial    int    `json:"RoundSerial"`
			Platform       int    `json:"Platform"`
			WagersType     int    `json:"WagersType"`
			WagersDate     string `json:"WagersDate"`
			Commissionable string `json:"Commissionable"`
			ModifiedDate   string `json:"ModifiedDate"`
			Payoff         string `json:"Payoff"`
			Revenue        string `json:"Revenue"`
			RoundDate      string `json:"RoundDate"`
		} `json:"data"`
		Amount struct {
		} `json:"amount"`
		SubNumber   int `json:"SubNumber"`
		TotalNumber int `json:"TotalNumber"`
	} `json:"data"`
	Version string `json:"version"`
	Guid    string `json:"guid"`
}

type Betlog66 struct {
	startTime  time.Time
	betlog     betlog66
	BetlogJSON string
}

func (betlog Betlog66) GetBetlog() (string, error) {

	url := "https://api.bbbattleapi.com/API/M/User/Wagers"

	headers := map[string]string{
		"Ekey":  "melttabbb",
		"Token": "4082342799270551",
	}

	queries := map[string]string{
		"startTime":  betlog.startTime.Format("2006-01-02 15:04:05"),
		"endTime":    betlog.startTime.Format("2006-01-02 15:04:") + "59",
		"limitStart": "0",
		"limitEnd":   "10000",
	}

	req := curl.NewRequest()

	resp, err := req.
		SetUrl(url).
		SetHeaders(headers).
		SetQueries(queries).
		Get()

	if err != nil {
		return "", err
	} else {
		json.Unmarshal([]byte(resp.Body), &betlog.betlog)

		if betlog.betlog.ErrorCode != "00" || betlog.betlog.Status != "000" {
			panic("Operation failed")
		}
	}

	return resp.Body, nil
}
