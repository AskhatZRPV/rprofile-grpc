package parser

import (
	"github.com/AskhatZRPV/rprofile-grpc/internal/core/usecase"
	"github.com/AskhatZRPV/rprofile-grpc/internal/domain/parser"
	"github.com/AskhatZRPV/rprofile-grpc/internal/domain/rusprofileclient"
)

type Payload struct {
	INN string
}

type Result = parser.ParseResult

type UseCase = usecase.UseCase[*Payload, *Result]

type implementation struct {
	hc             rusprofileclient.HttpClient
	parserProvider parser.Provider
}

func New(
	hc rusprofileclient.HttpClient,
	parserProvider parser.Provider,
) UseCase {
	return &implementation{hc, parserProvider}
}
