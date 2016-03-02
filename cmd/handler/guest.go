package handler

import (
	"net/http"

	"github.com/huyjack178/izi-guest-manager-ver0.1/cmd/store"
	"github.com/huyjack178/izi-guest-manager-ver0.1/cmd/util/logs"
	"github.com/huyjack178/izi-guest-manager-ver0.1/cmd/util/xhttp"
	"github.com/huyjack178/izi-guest-manager-ver0.1/cmd/model"
	"github.com/huyjack178/izi-guest-manager-ver0.1/cmd/util/message"
)

type GuestCtrl struct {
	guestStore *store.GuestStore
}

var log = logs.New("guestCtrl")

func NewGuestCtrl(guestStore *store.GuestStore) *GuestCtrl {
	return &GuestCtrl{
		guestStore: guestStore,
	}
}

func (this *GuestCtrl) List(w http.ResponseWriter, r *http.Request) {
	//guests, err := this.guestStore.List()
	//
	//if err != nil {
	//	xhttp.ResponseJson(w, xhttp.NewResponse(http.StatusInternalServerError, err.Error()))
	//	return
	//}

	xhttp.ResponseJson(w, xhttp.NewResponse(http.StatusCreated, "OK"))
}

func (this *GuestCtrl) Get(w http.ResponseWriter, r *http.Request) {
	ctx := xhttp.GetContext(r)
	id := ctx.Params.ByName("id")

	guest, err := this.guestStore.Get(id)
	if err != nil {
		 xhttp.ResponseJson(w, xhttp.NewResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	xhttp.ResponseJson(w, xhttp.NewResponse(http.StatusOK, guest))
}

func (this *GuestCtrl) Create(w http.ResponseWriter, r *http.Request) {
	var guest model.Guest

	if err := xhttp.BindJSON(r, &guest); err != nil {
		xhttp.ResponseJson(w, xhttp.NewResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	if err := this.guestStore.Create(&guest); err != nil {
		xhttp.ResponseJson(w, xhttp.NewResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	xhttp.ResponseJson(w, xhttp.NewResponse(http.StatusCreated, message.CREATE_SUCCESS))
}

func (this *GuestCtrl) Update(w http.ResponseWriter, r *http.Request) {
	ctx := xhttp.GetContext(r)
	id := ctx.Params.ByName("id")

	var guest model.Guest

	if err := xhttp.BindJSON(r, &guest); err != nil {
		xhttp.ResponseJson(w, xhttp.NewResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	if err := this.guestStore.Update(id, &guest); err != nil {
		xhttp.ResponseJson(w, xhttp.NewResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	xhttp.ResponseJson(w, xhttp.NewResponse(http.StatusCreated, message.UPDATE_SUCCESS))
}

func (this *GuestCtrl) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := xhttp.GetContext(r)
	id := ctx.Params.ByName("id")

	if err := this.guestStore.Delete(id); err != nil {
		xhttp.ResponseJson(w, xhttp.NewResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	xhttp.ResponseJson(w, xhttp.NewResponse(http.StatusOK, message.DELETE_SUCCESS))
}