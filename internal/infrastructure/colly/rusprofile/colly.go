package colly

import (
	"fmt"
	"strings"

	"github.com/AskhatZRPV/rprofile-grpc/internal/core/config"
	"github.com/AskhatZRPV/rprofile-grpc/internal/domain/parser"
	"github.com/gocolly/colly/v2"
	"github.com/pkg/errors"
)

type implementation struct {
	baseURL string
}

func New(config *config.Config) parser.Provider {
	return &implementation{config.HttpClient.BaseUrl}
}

func (i *implementation) GetInfo(id string) (*parser.ParseResult, error) {
	res, err := i.getInfo(i.buildUrl(id))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (i *implementation) getInfo(url string) (*parser.ParseResult, error) {
	var res parser.ParseResult

	c := colly.NewCollector()

	c.OnHTML("h1[itemprop]", func(e *colly.HTMLElement) {
		if e.Attr("itemprop") == "name" {
			text := strings.TrimSpace(e.Text)
			res.CompanyName = text
		}
	})

	c.OnHTML("#clip_inn", func(e *colly.HTMLElement) {
		res.Inn = e.Text
	})

	c.OnHTML("#clip_kpp", func(e *colly.HTMLElement) {
		res.Kpp = e.Text
	})

	c.OnHTML("div.company-row.hidden-parent", func(e *colly.HTMLElement) {
		res.LeaderName = e.ChildText(".company-info__text")
	})

	err := c.Visit(url)
	// TODO:
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return &res, nil
}

func (i *implementation) buildUrl(id string) string {
	return fmt.Sprintf("%s/id/%s", i.baseURL, id)
}
