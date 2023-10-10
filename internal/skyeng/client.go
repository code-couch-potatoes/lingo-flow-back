package skyeng

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	baseURL = "https://dictionary.skyeng.ru/api/public/v1/"

	meaningsURLTemplate            = baseURL + "meanings?ids=%s"
	wordsSearchURLTemplate         = baseURL + "words/search?search=%s"
	pageableWordsSearchURLTemplate = wordsSearchURLTemplate + "&page=%d&pageSize=%d"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{httpClient: &http.Client{Timeout: 60 * time.Second}}
}

func (c *Client) Meanings(ctx context.Context, ids ...int) ([]Meaning, error) {
	sIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		sIDs = append(sIDs, strconv.Itoa(id))
	}

	targetURL := fmt.Sprintf(meaningsURLTemplate, strings.Join(sIDs, ","))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetURL, nil)
	if err != nil {
		return nil, err
	}

	data, err := c.doRequest(req, targetURL)
	if err != nil {
		return nil, err
	}

	var meanings []Meaning
	err = json.Unmarshal(data, &meanings)
	if err != nil {
		return nil, err
	}

	return meanings, nil
}

func (c *Client) WordsSearch(ctx context.Context, word string) ([]Word, error) {
	targetURL := fmt.Sprintf(wordsSearchURLTemplate, word)

	return c.search(ctx, targetURL)
}

func (c *Client) WordsSearchByPage(ctx context.Context, word string, pageNum, pageSize int) ([]Word, error) {
	targetURL := fmt.Sprintf(pageableWordsSearchURLTemplate, word, pageNum, pageSize)

	return c.search(ctx, targetURL)
}

func (c *Client) search(ctx context.Context, targetURL string) ([]Word, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetURL, nil)
	if err != nil {
		return nil, err
	}

	data, err := c.doRequest(req, targetURL)
	if err != nil {
		return nil, err
	}

	var words []Word
	err = json.Unmarshal(data, &words)
	if err != nil {
		return nil, err
	}

	return words, nil
}

func (c *Client) doRequest(req *http.Request, targetURL string) ([]byte, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, parseError(resp.Body, targetURL)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func parseError(body io.Reader, targetURL string) error {
	data, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	var errResp errorResponse
	err = json.Unmarshal(data, &errResp)
	if err != nil {
		return err
	}

	return fmt.Errorf("can't search word url=%s, status=%d, reason=%s:%s", targetURL, errResp.Status, errResp.Title, errResp.Detail)
}

type errorResponse struct {
	Status int    `json:"status"`
	Type   string `json:"type"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}
