package parser

type Provider interface {
	GetInfo(string) (*ParseResult, error)
}
