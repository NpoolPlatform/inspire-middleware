package accounting

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/accounting"

	accounting1 "github.com/NpoolPlatform/inspire-middleware/pkg/accounting"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//nolint:gocyclo
func (s *Server) Accounting(ctx context.Context, in *npool.AccountingRequest) (*npool.AccountingResponse, error) {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		return &npool.AccountingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		return &npool.AccountingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetGoodID()); err != nil {
		return &npool.AccountingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetOrderID()); err != nil {
		return &npool.AccountingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetPaymentID()); err != nil {
		return &npool.AccountingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetCoinTypeID()); err != nil {
		return &npool.AccountingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetPaymentCoinTypeID()); err != nil {
		return &npool.AccountingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	paymentCoinUSDCurrency, err := decimal.NewFromString(in.GetPaymentCoinUSDCurrency())
	if err != nil {
		return &npool.AccountingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if paymentCoinUSDCurrency.Cmp(decimal.NewFromInt(0)) <= 0 {
		return &npool.AccountingResponse{}, status.Error(codes.InvalidArgument, "PaymentCoinUSDCurrency <= 0")
	}
	untis, err := decimal.NewFromString(in.GetUnits())
	if err != nil {
		return &npool.AccountingResponse{}, status.Error(codes.InvalidArgument, "Units == 0")
	}
	if untis.Cmp(decimal.NewFromInt(0)) == 0 {
		return &npool.AccountingResponse{}, status.Error(codes.InvalidArgument, "Units == 0")
	}

	paymentAmount, err := decimal.NewFromString(in.GetPaymentAmount())
	if err != nil {
		return &npool.AccountingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	goodValue, err := decimal.NewFromString(in.GetGoodValue())
	if err != nil {
		return &npool.AccountingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if goodValue.Cmp(decimal.NewFromInt(0)) <= 0 {
		return &npool.AccountingResponse{}, status.Error(codes.InvalidArgument, "GoodValue is invalid")
	}

	infos, err := accounting1.Accounting(
		ctx,
		in.GetAppID(),
		in.GetUserID(),
		in.GetGoodID(),
		in.GetOrderID(),
		in.GetPaymentID(),
		in.GetCoinTypeID(),
		in.GetPaymentCoinTypeID(),
		paymentCoinUSDCurrency,
		in.GetUnits(),
		in.GetSettleType(),
		paymentAmount,
		goodValue,
	)
	if err != nil {
		return &npool.AccountingResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AccountingResponse{
		Infos: infos,
	}, nil
}
