package mapper

import (
	"fiap-tech-challenge-api/internal/core/domain"
	"fmt"
	"strings"
)

type Pedido interface {
	MapReqToDTO(req *domain.PedidoRequest) *domain.PedidoDTO
	MapDTOToModels(req []*domain.PedidoDTO) []*domain.Pedido
	MapDTOToModel(req *domain.PedidoDTO) *domain.Pedido
	MapDTOToResponse(dto *domain.PedidoDTO) *domain.PedidoResponse
}

type pedido struct {
}

func (p pedido) MapDTOToResponse(dto *domain.PedidoDTO) *domain.PedidoResponse {
	return &domain.PedidoResponse{
		Pedido: &domain.Pedido{
			Id:         dto.Id,
			Status:     dto.Status,			
			Observacao: dto.Observacao,
			CreatedAt:  dto.CreatedAt,
			UpdatedAt:  dto.UpdatedAt,
		},
	}
}

func (p pedido) MapReqToDTO(req *domain.PedidoRequest) *domain.PedidoDTO {
	ids := make([]string, len(req.ProdutoIds))
	for i, id := range req.ProdutoIds {
		ids[i] = fmt.Sprint(id)
	}
	return &domain.PedidoDTO{
		ClienteId:  req.ClienteId,
		Observacao: req.Observacao,
		ProdutoIDS: strings.Join(ids, ","),
	}
}

func (p pedido) MapDTOToModel(req *domain.PedidoDTO) *domain.Pedido {
	//TODO implement me
	panic("implement me")
}

func (p pedido) MapDTOToModels(req []*domain.PedidoDTO) []*domain.Pedido {
	pedidos := make([]*domain.Pedido, len(req))
	for i, dto := range req {
		pedidos[i] = &domain.Pedido{
			Id:         dto.Id,
			ClienteId:  dto.ClienteId,
			Status:     dto.Status,
			Observacao: dto.Observacao,
			CreatedAt:  dto.CreatedAt,
			UpdatedAt:  dto.UpdatedAt,
		}
	}

	return pedidos
}

func NewPedidoMapper() Pedido {
	return &pedido{}
}
