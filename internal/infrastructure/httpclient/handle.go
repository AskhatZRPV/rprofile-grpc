package httpclient

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

func (i *httpClient) GetIdByInn(inn string) (string, error) {
	res, err := i.get(i.buildUrl(inn))
	if err != nil {
		return "", err
	}

	var rb responseBody
	err = json.Unmarshal(res, &rb)
	if err != nil {
		return "", err
	}

	var aciId string
	if len(rb.Ul) > 0 {
		ul := rb.Ul
		aciId = ul[0].AciID
	} else {
		return "nil", errors.New("not found")
	}

	return aciId, nil
}

func (i *httpClient) buildUrl(inn string) string {
	return fmt.Sprintf("%s/ajax.php?query=%s&action=search", i.baseUrl, inn)
}
