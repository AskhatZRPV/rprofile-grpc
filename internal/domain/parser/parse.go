package parser

type ParseResult struct {
	Inn, Kpp, CompanyName, LeaderName string
}

func New(inn, kpp, companyName, leaderName string) *ParseResult {
	return &ParseResult{
		Inn:         inn,
		Kpp:         kpp,
		CompanyName: companyName,
		LeaderName:  leaderName,
	}
}
