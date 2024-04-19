package repository

import (
	"context"
	"fiap-tech-challenge-api/internal/core/domain"
	"github.com/joomcode/errorx"
	"xorm.io/xorm"
)

const tableNameFila string = "fila"

type fila struct {
	session *xorm.Session
}

type FilaRepo interface {
	Insere(ctx context.Context, fila *domain.Fila) error
	AtualizaStatus(ctx context.Context, status string, pedidoId int64) error
	ExistePorPedidoStatus(ctx context.Context, pedidoId int64) (bool, error)
}

func NewFilaRepo(connector DBConnector) FilaRepo {
	session := connector.GetORM().Table(tableNameFila)
	return &fila{
		session: session,
	}
}

func (f *fila) Insere(ctx context.Context, fila *domain.Fila) error {
	_, err := f.session.Context(ctx).Insert(fila)
	if err != nil {
		return err
	}

	return nil
}

func (f *fila) AtualizaStatus(ctx context.Context, status string, pedidoId int64) error {
	_, err := f.session.Context(ctx).Where("pedido_id = ?", pedidoId).Update(&domain.Fila{Status: status})
	if err != nil {
		return errorx.InternalError.New(err.Error())
	}

	return nil
}

func (f *fila) ExistePorPedidoStatus(ctx context.Context, pedidoId int64) (bool, error) {
	existe, err := f.session.Context(ctx).Exist(&domain.Fila{PedidoId: pedidoId})
	if err != nil {
		return false, err
	}

	return existe, nil
}
