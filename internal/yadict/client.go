package yadict

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	LangEnRu = "en-ru"
	LangRuEn = "ru-en"
	LangEnEn = "en-en"
)

const lookupURLTemplate = "https://dictionary.yandex.net/api/v1/dicservice.json/lookup?key=%s&lang=%s&text=%s"

var (
	ErrKeyInvalid            = errors.New("invalid API key")
	ErrKeyBlocked            = errors.New("this API key has been blocked")
	ErrDailyReqLimitExceeded = errors.New("exceeded the daily limit on the number of requests")
	ErrTextTooLong           = errors.New("the text size exceeds the maximum")
	ErrLangNotSupported      = errors.New("the specified translation direction is not supported")
)

var errorMapping = map[int]error{
	401: ErrKeyInvalid,
	402: ErrKeyBlocked,
	403: ErrDailyReqLimitExceeded,
	413: ErrTextTooLong,
	501: ErrLangNotSupported,
}

type Client struct {
	httpClient *http.Client
	apiKey     string
}

func NewClient(apiKey string) *Client {
	client := Client{
		httpClient: &http.Client{Timeout: 60 * time.Second},
		apiKey:     apiKey,
	}

	return &client
}

func (c *Client) LookUp(ctx context.Context, lang string, text string) ([]Word, error) {
	targetURL := fmt.Sprintf(lookupURLTemplate, c.apiKey, lang, text)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var errResp errorResponse
		err = json.Unmarshal(data, &errResp)
		if err != nil {
			return nil, err
		}

		if err, ok := errorMapping[errResp.Code]; ok {
			return nil, err
		}

		return nil, errors.New(errResp.Message)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var dictResp response
	err = json.Unmarshal(data, &dictResp)
	if err != nil {
		return nil, err
	}

	logText, _ := json.MarshalIndent(&dictResp, "", "  ")
	log.Println(string(logText))

	return dictResp.toModel(), nil
}

type response struct {
	Def []struct {
		Text string `json:"text"`
		Pos  string `json:"pos"`
		Ts   string `json:"ts"`
		Tr   []struct {
			Text string `json:"text"`
			Pos  string `json:"pos"`
			Syn  []struct {
				Text string `json:"text"`
			} `json:"syn"`
			Mean []struct {
				Text string `json:"text"`
			} `json:"mean"`
			Ex []struct {
				Text string `json:"text"`
				Tr   []struct {
					Text string `json:"text"`
				} `json:"tr"`
			} `json:"ex"`
		} `json:"tr"`
	} `json:"def"`
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (r response) toModel() []Word {
	res := make([]Word, len(r.Def))
	for i, def := range r.Def {
		res[i].Src = def.Text
		res[i].POS = def.Pos
		res[i].Ts = def.Ts
		res[i].Tr = make([]string, 0, len(def.Tr))
		for _, tr := range def.Tr {
			res[i].Tr = append(res[i].Tr, tr.Text)
		}
	}

	return res
}
