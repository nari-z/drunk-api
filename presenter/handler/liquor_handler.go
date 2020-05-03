package handler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/runtime/middleware"

	"github.com/nari-z/drunk-api/usecase"
	"github.com/nari-z/drunk-api/domain/model"

	"github.com/nari-z/drunk-api/generate/models"
	"github.com/nari-z/drunk-api/generate/restapi/operations"
)

// LiquorHandler is request handler for liquor.
type LiquorHandler interface {
	GetLiquorList(params operations.GetLiquorsParams) middleware.Responder
	RegistLiquor(params operations.AddLiquorParams) middleware.Responder
}

type liquorHandler struct {
	LiquorUseCase usecase.LiquorUseCase
}

// NewLiquorHandler return LiquorHandler。
func NewLiquorHandler(u usecase.LiquorUseCase) LiquorHandler {
	return &liquorHandler{u}
}

func (h *liquorHandler) GetLiquorList(params operations.GetLiquorsParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	fmt.Println("LiquorHandler.GetLiquorList.")

	liquorList, err := h.LiquorUseCase.GetLiquorList(ctx)
	if err != nil {
		// TODO: send error message
		return operations.NewAddLiquorBadRequest()
	}

	// convert to response param
	resParams := make([]*models.GetLiquorResponseParams, len(liquorList))
	for i, liquor := range liquorList {
		resParam := h.toGetLiquorResponseParams(liquor)
		if resParam == nil {
			// TODO: send error message
			// TODO: 500 error
			return operations.NewAddLiquorBadRequest()
		}

		resParams[i] = resParam
	}

	return operations.NewGetLiquorsOK().WithPayload(resParams)
}

func (h *liquorHandler) RegistLiquor(params operations.AddLiquorParams) middleware.Responder {
	fmt.Println("In RegistLiquor.")
	ctx := params.HTTPRequest.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	_, err := h.LiquorUseCase.RegistLiquor(ctx, params.Body.Name, params.Body.FileName, params.Body.Image)
	if err != nil {
		// TODO: send error message
		return operations.NewAddLiquorBadRequest()
	}

	return operations.NewAddLiquorCreated()
}

func (h *liquorHandler) toGetLiquorResponseParams(src *model.Liquor) *models.GetLiquorResponseParams {
	if src == nil {
		// TODO: nil safe
		return nil
	}

	return &models.GetLiquorResponseParams {
		ID: strconv.FormatUint(src.ID, 10),// TODO: not use db id
		ImageFilePath: src.ImageFilePath,
		Name: src.Name,
		UpdatedAt: strfmt.Date(src.UpdatedAt),
	}
}
