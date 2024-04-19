package domain

import (
	"time"
)

const (
	StatusRecebido            string = "recebido"
	StatusEmPreparacao        string = "em_preparacao"
	StatusPronto              string = "pronto"
	StatusFinalizado          string = "finalizado"
	StatusAguardandoPagamento string = "aguardando_pagamento"
	StatusPagamentoAprovado   string = "pagamento_aprovado"
	StatusPagamentoRecusado   string = "pagamento_recusado"
)

type PedidoRequest struct {
	ClienteId  int64   `json:"cliente_id" validate:"required"`
	ProdutoIds []int64 `json:"produtos" validate:"required"`
	Observacao string  `json:"observacao"`
}

type PedidoDTO struct {
	Id         int64      `xorm:"pk autoincr 'pedido_id'"`
	ClienteId  int64      `xorm:"'cliente_id'"`	
	ProdutoIDS string     `xorm:"'produtos'"`
	Status     string     `xorm:"'status'"`
	Observacao string
	CreatedAt  time.Time `xorm:"created"`
	UpdatedAt  time.Time `xorm:"updated"`
}

type PedidosDTO []*PedidoDTO

func (a PedidosDTO) Len() int      { return len(a) }
func (a PedidosDTO) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a PedidosDTO) Less(i, j int) bool {
	if a[i].Status != a[j].Status {
		return a[i].Status == StatusPronto || (a[i].Status == StatusEmPreparacao && a[j].Status != StatusPronto)
	}

	return a[i].CreatedAt.Before(a[j].CreatedAt)

}

type PedidoProduto struct {
	Id        int64
	PedidoId  int64     `xorm:"index"`
	ProdutoId int64     `xorm:"index"`
	CreatedAt time.Time `json:"created_at" xorm:"created"`
}

func (dto *PedidoDTO) TableName() string {
	return "pedido"
}

type PedidoResponse struct {
	*Pedido
}

type Pedido struct {
	Id         int64      `json:"id"`
	ClienteId  int64      `json:"cliente_id"`
	Status     string     `json:"status"`
	Observacao string     `json:"observacao"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type Fila struct {
	Id         int64
	PedidoId   int64  `xorm:"index unique"`
	Status     string `xorm:"status"`
	Observacao string
	CreatedAt  time.Time `xorm:"created"`
	UpdatedAt  time.Time `xorm:"updated"`
}

type StatusRequest struct {
	Status string `json:"status" validate:"required"`
}
