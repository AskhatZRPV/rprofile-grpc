package httpclient

import (
	"io"
	"net/http"

	"github.com/AskhatZRPV/rprofile-grpc/internal/core/config"
	"github.com/AskhatZRPV/rprofile-grpc/internal/domain/rusprofileclient"
)

type httpClient struct {
	baseUrl string
	*http.Client
}

func New(cfg *config.Config) rusprofileclient.HttpClient {
	return &httpClient{
		cfg.BaseUrl,
		&http.Client{
			Timeout: cfg.HttpClient.Timeout,
		},
	}
}

func (hc *httpClient) get(url string) ([]byte, error) {
	req, err := http.NewRequest(
		"GET", url, nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	byteres, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return byteres, nil
}
