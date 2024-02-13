package handlers

import (
	"github.com/evoaway/erc20-transfers-storage-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"log"
	"net/http"
)

func GetTransfersByAddress(w http.ResponseWriter, r *http.Request) {
	address := requests.NewGetAddress(r)
	err, trs := DB(r).SelectTransfersByAddress(address)
	log.Println(trs)
	if err != nil {
		Log(r).WithError(err).Error("error processing get_transaction request")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	w.Write([]byte(address))
}
