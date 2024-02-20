package handlers

import (
	"context"
	"github.com/evoaway/erc20-transfers-storage-svc/internal/data"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	transfer
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxTransfer(t data.ITransfer) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, transfer, t)
	}
}

func Transfer(r *http.Request) data.ITransfer {
	return r.Context().Value(transfer).(data.ITransfer)
}
