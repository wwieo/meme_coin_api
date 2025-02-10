package api

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"meme_coin_api/service/controller/memeCoinCtrl"
	"meme_coin_api/service/internal/errorx"
	boMemeCoin "meme_coin_api/service/internal/model/bo/memeCoin"
	"net/http"
)

func NewMemeCoin(pack memeCoinPack) {
	m := &memeCoin{pack: pack}
	group := pack.Root.Group("meme_coin")
	{
		group.POST("", m.createMemeCoin)
	}
}

type memeCoinPack struct {
	dig.In

	Root         *gin.RouterGroup
	MemeCoinCtrl memeCoinCtrl.MemeCoinCtrl
}

type memeCoin struct {
	pack memeCoinPack
}

// createMemeCoin
//
//	@Summary	Create a meme coin
//	@Tags		meme_coin
//	@version	1.0
//	@produce	json
//	@Param		request	body	api.createMemeCoin.body	true	"Request payload for creating a meme coin"
//	@Success	200
//	@Failure	400	{object}	errorx.ErrorResponse
//	@Router		/meme_coin [POST]
func (api *memeCoin) createMemeCoin(ctx *gin.Context) {
	type body struct {
		Name        string `valid:"required" json:"Name"`
		Description string `valid:"-" json:"Description"`
	}
	var reqBody body
	if err := ctx.BindJSON(&reqBody); err != nil {
		errorx.RespondWithError(ctx, http.StatusInternalServerError, err)
		return
	}
	if _, err := govalidator.ValidateStruct(reqBody); err != nil || govalidator.HasWhitespaceOnly(reqBody.Name) {
		errorx.RespondWithError(ctx, http.StatusBadRequest, errorx.ParamKeyRequired)
		return
	}

	args := boMemeCoin.CreateArgs{
		Name:        reqBody.Name,
		Description: reqBody.Description,
	}
	if err := api.pack.MemeCoinCtrl.Create(ctx, args); err != nil {
		errorx.RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.Status(http.StatusOK)
}
