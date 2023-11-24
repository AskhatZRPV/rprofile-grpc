package parser

import (
	"context"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) (*Result, error) {
	respId, err := i.hc.GetIdByInn(p.INN)
	if err != nil {
		return nil, err
	}

	res, err := i.parserProvider.GetInfo(respId)
	if err != nil {
		return nil, err
	}

	return res, nil

}
