package handlers

import (
	"fiap-tech-challenge-producao/internal/core/commons"
	"fiap-tech-challenge-producao/internal/core/domain"
	"fiap-tech-challenge-producao/internal/core/usecase"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
	serverErr "github.com/rhuandantas/fiap-tech-challenge-commons/pkg/errors"
	"github.com/rhuandantas/fiap-tech-challenge-commons/pkg/middlewares/auth"
	"github.com/rhuandantas/fiap-tech-challenge-commons/pkg/util"
	"net/http"
)

type Producao struct {
	validator        util.Validator
	listaPorStatusUC usecase.PegaPedidoPorID
	atualizaStatusUC usecase.AtualizaStatusProducao
	fila             usecase.CadastrarFila
	tokenJwt         auth.Token
}

func NewProducao(validator util.Validator,
	listaPorStatusUC usecase.PegaPedidoPorID,
	atualizaStatusUC usecase.AtualizaStatusProducao,
	fila usecase.CadastrarFila,
	tokenJwt auth.Token,
) *Producao {
	return &Producao{
		validator:        validator,
		listaPorStatusUC: listaPorStatusUC,
		atualizaStatusUC: atualizaStatusUC,
		fila:             fila,
		tokenJwt:         tokenJwt,
	}
}

func (h *Producao) RegistraRotasFila(server *echo.Echo) {
	server.GET("/producao/:pedido_id", h.pegaPorPedidoId)
	server.POST("/internal/producao", h.adicionaProducao)
	server.PATCH("/producao/:pedido_id", h.atualizaStatus)
}

// pegaPorPedidoId godoc
// @Summary lista pedido em producao
// @Tags Producao
// @Produce json
// @Param        pedidoID   path      string  true  "id do  pedidos em producao"
// @Success 200 {object} domain.Producao
// @Router /producao/{pedidoID} [get]
func (h *Producao) pegaPorPedidoId(ctx echo.Context) error {
	pedidoID := ctx.Param("pedido_id")
	pedidos, err := h.listaPorStatusUC.PesquisaPorID(ctx.Request().Context(), pedidoID)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}
	return ctx.JSON(http.StatusOK, pedidos)
}

// atualizaStatus godoc
// @Summary atualiza o status do pedido na producao
// @Tags Producao
// @Accept json
// @Param        id   path      integer  true  "id do pedido"
// @Param        id   body      domain.StatusRequest  true  "status"
// @Produce json
// @Router /producao/{oedidoID} [patch]
func (h *Producao) atualizaStatus(ctx echo.Context) error {
	var (
		status struct {
			Status string `json:"status"`
		}
		err error
	)

	if err = ctx.Bind(&status); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(err.Error()))
	}

	id := ctx.Param("pedido_id")

	err = h.atualizaStatusUC.Atualiza(ctx.Request().Context(), status.Status, id)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}

	return ctx.JSON(http.StatusOK, status)
}

// adicionaProducao godoc
// @Summary adiciona pedido a produção
// @Tags Producao
// @Produce json
// @Router /producao [post]
func (h *Producao) adicionaProducao(ctx echo.Context) error {
	var (
		producao domain.Producao
		err      error
	)
	if err = ctx.Bind(&producao); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(err.Error()))
	}

	err = h.validator.ValidateStruct(producao)
	if err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(err.Error()))
	}

	err = h.fila.Cadastra(ctx.Request().Context(), &producao)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}
	return ctx.JSON(http.StatusOK, producao)
}
