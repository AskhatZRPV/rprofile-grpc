package search

import (
	"context"
	"log/slog"

	rprofilev1 "github.com/AskhatZRPV/rprofile-grpc/gen/protos/go/proto/rprofile"
	"github.com/AskhatZRPV/rprofile-grpc/internal/application/rusprofile/parser"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CompanyInfo struct {
	usecase parser.UseCase
	rprofilev1.UnimplementedCompanyInfoServer
	logger *slog.Logger
}

func New(usecase parser.UseCase, logger *slog.Logger) *CompanyInfo {
	return &CompanyInfo{usecase: usecase, logger: logger}
}

func (s *CompanyInfo) SearchInfo(
	ctx context.Context,
	in *rprofilev1.SearchInfoRequest,
) (*rprofilev1.SearchInfoResponse, error) {
	r, err := s.usecase.Execute(ctx, usecasePayloadFromRequest(in.GetInn()))
	if err != nil {
		s.logger.Error(err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := responseFromResult(r)
	return resp, nil
}

func usecasePayloadFromRequest(inn string) *parser.Payload {
	return &parser.Payload{
		INN: inn,
	}
}

func responseFromResult(r *parser.Result) *rprofilev1.SearchInfoResponse {
	return &rprofilev1.SearchInfoResponse{
		Inn:         r.Inn,
		Kpp:         r.Kpp,
		CompanyName: r.CompanyName,
		LeaderName:  r.LeaderName,
	}
}
