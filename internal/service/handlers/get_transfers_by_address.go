package handlers

import (
	"github.com/evoaway/erc20-transfers-storage-svc/internal/data"
	"github.com/evoaway/erc20-transfers-storage-svc/internal/service/requests"
	"github.com/evoaway/erc20-transfers-storage-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"strconv"
)

func GetTransfersByAddress(w http.ResponseWriter, r *http.Request) {
	address := requests.NewGetAddress(r)
	err, trs := DB(r).SelectTransfersByAddress(address)
	if err != nil {
		Log(r).WithError(err).Error("error processing get_transaction request")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	ape.Render(w, resources.TransferListResponse{
		Data: getArray(trs),
	})
}

func getArray(transfers []data.Transfer) []resources.Transfer {
	var resTransfers = make([]resources.Transfer, len(transfers))
	for i, t := range transfers {
		resTransfers[i].Key = resources.Key{ID: strconv.Itoa(i), Type: resources.TRANSFER}
		resTransfers[i].Attributes = resources.TransferAttributes{
			Id:    t.Id,
			From:  t.FromAddress,
			To:    t.ToAddress,
			Value: t.Value,
		}
	}
	return resTransfers
}
