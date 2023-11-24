package parser

import "github.com/AskhatZRPV/rprofile-grpc/internal/core/domainerr"

var (
	ErrInnDoesNotExist = domainerr.New("inn does not exist")
)
