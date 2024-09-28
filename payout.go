package cryptomus

import (
	"encoding/json"
	"errors"
	"time"
)

const (
	createPayoutEndpoint       = "/payout"
	payoutInfoEndpoint         = "/payout/info"
	payoutHistoryEndpoint      = "/payout/list"
	payoutServicesListEndpoint = "/payout/services"
)

type PayoutRequest struct {
	Amount     string `json:"amount"`
	Currency   string `json:"currency"`
	OrderId    string `json:"order_id"`
	Address    string `json:"address"`
	IsSubtract bool   `json:"is_subtract"`
	Network    string `json:"network"`
}

type PayoutRequestOptions struct {
	UrlCallback  string `json:"url_callback,omitempty"`
	ToCurrency   string `json:"to_currency,omitempty"`
	CourseSource string `json:"course_source,omitempty"`
	FromCurrency string `json:"from_currency,omitempty"`
	Priority     string `json:"priority,omitempty"`
	Memo         string `json:"memo,omitempty"`
}

type Payout struct {
	UUID          string `json:"uuid"`
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	Network       string `json:"network"`
	Address       string `json:"address"`
	TxId          string `json:"txid"`
	Status        string `json:"status"`
	IsFinal       bool   `json:"is_final"`
	Balance       string `json:"balance"`
	PayerCurrency string `json:"payer_currency"`
	PayerAmount   string `json:"payer_amount"`
}

type payoutRawResponse struct {
	Result *Payout
	State  int8
}

type PayoutInfoRequest struct {
	PayoutUUID string `json:"uuid,omitempty"`
	OrderId    string `json:"order_id,omitempty"`
}

type PayoutHistoryResponse struct {
	Payouts  []*Payout
	Paginate *PayoutHistoryPaginate
}

type PayoutHistoryPaginate struct {
	Count          int16  `json:"count"`
	HasPages       bool   `json:"hasPages"`
	NextCursor     string `json:"nextCursor,omitempty"`
	PreviousCursor string `json:"previousCursor,omitempty"`
	PerPage        int16  `json:"perPage"`
}

type payoutHistoryRawResponse struct {
	State    int8                   `json:"state"`
	Result   []*Payout              `json:"result"`
	Paginate *PayoutHistoryPaginate `json:"paginate"`
}

type PayoutService struct {
	Network     string                  `json:"network"`
	Currency    string                  `json:"currency"`
	IsAvailable bool                    `json:"isAvailable"`
	Limit       *PayoutServiceLimit     `json:"limit"`
	Commision   *PayoutServiceCommision `json:"commision"`
}

type PayoutServiceLimit struct {
	MinAmount string `json:"minAmount"`
	MaxAmount string `json:"maxAmount"`
}

type PayoutServiceCommision struct {
	FeeAmount string `json:"feeAmount"`
	Percent   string `json:"percent"`
}

type payoutServiceListRawResponse struct {
	Result []*PayoutService `json:"result"`
	State  int8             `json:"state"`
}

func (c *Cryptomus) CreatePayout(payoutReq *PayoutRequest) (*Payout, error) {
	res, err := c.fetch("POST", createPayoutEndpoint, payoutReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response := &payoutRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response.Result, nil
}

func (c *Cryptomus) GetPayoutInfo(payoutInfoReq *PayoutInfoRequest) (*Payout, error) {
	if payoutInfoReq.PayoutUUID == "" || payoutInfoReq.OrderId == "" {
		return nil, errors.New("you should pass one of required values [PayoutUUID, OrderId]")
	}

	res, err := c.fetch("POST", payoutInfoEndpoint, payoutInfoReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response := &payoutRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response.Result, nil
}

func (c *Cryptomus) GetPayoutHistory(dateFrom, dateTo time.Time) (*PayoutHistoryResponse, error) {
	payload := map[string]any{"date_from": dateFrom, "date_to": dateTo}
	res, err := c.fetch("POST", payoutHistoryEndpoint, payload)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response := &payoutHistoryRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	payoutHistory := &PayoutHistoryResponse{
		Payouts:  response.Result,
		Paginate: response.Paginate,
	}

	return payoutHistory, nil
}

func (c *Cryptomus) GetPayoutServicesList() ([]*PayoutService, error) {
	payload := make(map[string]any)
	res, err := c.fetch("POST", payoutServicesListEndpoint, payload)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response := &payoutServiceListRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response.Result, nil
}
