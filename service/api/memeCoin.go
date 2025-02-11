package api

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"meme_coin_api/service/controller/memeCoinCtrl"
	"meme_coin_api/service/internal/errorx"
	boMemeCoin "meme_coin_api/service/internal/model/bo/memeCoin"
	"net/http"
	"strconv"
)

func NewMemeCoin(pack memeCoinPack) {
	m := &memeCoin{pack: pack}
	group := pack.Root.Group("meme_coin")
	{
		group.GET(":id", m.getMemeCoin)
		group.POST("", m.createMemeCoin)
		group.POST(":id/poke", m.pokeMemeCoin)
		group.PATCH(":id", m.updateMemeCoin)
		group.DELETE(":id", m.deleteMemeCoin)
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

// getMemeCoin
//
//	@Summary	Retrieve data for a meme coin
//	@Tags		meme_coin
//	@version	1.0
//	@produce	json
//	@Param		id				path		int	true	"id = ?"
//	@Success	200				{object}	boMemeCoin.GetReply
//	@Failure	400				{object}	errorx.ErrorResponse
//	@Router		/meme_coin/{id}	[GET]
func (api *memeCoin) getMemeCoin(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		errorx.RespondWithError(ctx, http.StatusBadRequest, errorx.ParamKeyRequired)
		return
	}

	args := boMemeCoin.GetArgs{
		ID: int64(id),
	}
	result, err := api.pack.MemeCoinCtrl.Get(ctx, args)
	if err != nil {
		errorx.RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// createMemeCoin
//
//	@Summary	Create a meme coin
//	@Tags		meme_coin
//	@version	1.0
//	@produce	json
//	@Param		request	body		api.createMemeCoin.body	true	"Request payload for creating a meme coin"
//	@Success	200		{object}	boMemeCoin.CreateReply
//	@Failure	400		{object}	errorx.ErrorResponse
//	@Router		/meme_coin [POST]
func (api *memeCoin) createMemeCoin(ctx *gin.Context) {
	type body struct {
		Name        string `valid:"required" json:"name"`
		Description string `valid:"-" json:"description"`
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
	reply, err := api.pack.MemeCoinCtrl.Create(ctx, args)
	if err != nil {
		errorx.RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, reply)
}

// updateMemeCoin
//
//	@Summary	Update a meme coin
//	@Tags		meme_coin
//	@version	1.0
//	@produce	json
//	@Param		id		path	int						true	"id = ?"
//	@Param		request	body	api.updateMemeCoin.body	true	"Payload to update the description of a meme coin"
//	@Success	200
//	@Failure	400	{object}	errorx.ErrorResponse
//	@Router		/meme_coin/{id} [PATCH]
func (api *memeCoin) updateMemeCoin(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		errorx.RespondWithError(ctx, http.StatusBadRequest, errorx.ParamKeyRequired)
		return
	}
	type body struct {
		Description string `valid:"-" json:"description"`
	}
	var reqBody body
	if err = ctx.BindJSON(&reqBody); err != nil {
		errorx.RespondWithError(ctx, http.StatusInternalServerError, err)
		return
	}
	if _, err = govalidator.ValidateStruct(reqBody); err != nil {
		errorx.RespondWithError(ctx, http.StatusBadRequest, errorx.ParamKeyRequired)
		return
	}

	args := boMemeCoin.UpdateArgs{
		ID:          int64(id),
		Description: reqBody.Description,
	}
	if err = api.pack.MemeCoinCtrl.Update(ctx, args); err != nil {
		errorx.RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.Status(http.StatusOK)
}

// deleteMemeCoin
//
//	@Summary	Delete a meme coin
//	@Tags		meme_coin
//	@version	1.0
//	@produce	json
//	@Param		id	path	int	true	"id = ?"
//	@Success	200
//	@Failure	400				{object}	errorx.ErrorResponse
//	@Router		/meme_coin/{id}	[DELETE]
func (api *memeCoin) deleteMemeCoin(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		errorx.RespondWithError(ctx, http.StatusBadRequest, errorx.ParamKeyRequired)
		return
	}

	args := boMemeCoin.DeleteArgs{
		ID: int64(id),
	}
	if err = api.pack.MemeCoinCtrl.Delete(ctx, args); err != nil {
		errorx.RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.Status(http.StatusOK)
}

// pokeMemeCoin
//
//	@Summary	Poke a meme coin
//	@Tags		meme_coin
//	@version	1.0
//	@produce	json
//	@Param		id	path	int	true	"id = ?"
//	@Success	200
//	@Failure	400						{object}	errorx.ErrorResponse
//	@Router		/meme_coin/{id}/poke	[POST]
func (api *memeCoin) pokeMemeCoin(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		errorx.RespondWithError(ctx, http.StatusBadRequest, errorx.ParamKeyRequired)
		return
	}

	args := boMemeCoin.IncreasePopularityScoreArgs{
		ID: int64(id),
	}
	if err = api.pack.MemeCoinCtrl.IncreasePopularityScore(ctx, args); err != nil {
		errorx.RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.Status(http.StatusOK)
}
