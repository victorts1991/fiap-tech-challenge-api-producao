package domain

import (
	"time"
)

func (dto *Producao) TableName() string {
	return "producao"
}

type Producao struct {
	Id         int64
	PedidoId   string    `json:"pedido_id" xorm:"index unique"`
	Status     string    `json:"status" xorm:"status"`
	Observacao string    `json:"observacao" xorm:"observacao"`
	CreatedAt  time.Time `xorm:"created"`
	UpdatedAt  time.Time `xorm:"updated"`
}

type StatusRequest struct {
	Status string `json:"status" validate:"required"`
}
