package commission

type Handler struct {
	ID              *uuid.UUID
	AppID           *uuid.UUID
	UserID          *uuid.UUID
	GoodID          *uuid.UUID
	SettleType      *types.SettleType
	SettleMode      *types.SettleMode
	SettleInterval  *types.SettleInterval
	AmountOrPercent *decimal.Decimal
	StartAt         *uint32
	Threshold       *decimal.Decimal
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = &_id
		return nil
	}
}

func WithAppID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AppID = &_id
		return nil
	}
}

func WithUserID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.UserID = &_id
		return nil
	}
}

func WithGoodID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.GoodID = &_id
		return nil
	}
}

func WithSettleType(settleType *types.SettleType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if settleType == nil {
			return nil
		}
		switch *settleType {
		case types.SettleType_GoodOrderPercent:
		case types.SettleType_GoodOrderValuePercent:
		case types.SettleType_TechniqueFeePercent:
		default:
			return fmt.Errorf("invalid settletype")
		}
		h.SettleType = settleType
		return nil
	}
}
