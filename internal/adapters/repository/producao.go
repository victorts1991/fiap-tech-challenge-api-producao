package repository

import (
	"context"
	"fiap-tech-challenge-producao/internal/core/domain"
	"github.com/joomcode/errorx"
	db "github.com/rhuandantas/fiap-tech-challenge-commons/pkg/db/mysql"
	"xorm.io/xorm"
)

const tableNameFila string = "producao"

type producao struct {
	session *xorm.Session
}

type FilaRepo interface {
	Insere(ctx context.Context, fila *domain.Producao) error
	AtualizaStatus(ctx context.Context, status, pedidoId string) error
	PegaPedidoPorID(ctx context.Context, pedidoId string) (*domain.Producao, error)
}

func NewFilaRepo(connector db.DBConnector) FilaRepo {
	err := connector.SyncTables(new(domain.Producao))
	if err != nil {
		panic(err.Error())
	}

	session := connector.GetORM().Table(tableNameFila)
	return &producao{
		session: session,
	}
}

func (f *producao) Insere(ctx context.Context, fila *domain.Producao) error {
	_, err := f.session.Context(ctx).Insert(fila)
	if err != nil {
		return err
	}

	return nil
}

func (f *producao) AtualizaStatus(ctx context.Context, status, pedidoId string) error {
	_, err := f.session.Context(ctx).Where("pedido_id = ?", pedidoId).Update(&domain.Producao{Status: status})
	if err != nil {
		return errorx.InternalError.New(err.Error())
	}

	return nil
}

func (f *producao) PegaPedidoPorID(ctx context.Context, pedidoId string) (*domain.Producao, error) {
	dto := &domain.Producao{PedidoId: pedidoId}
	existe, err := f.session.Context(ctx).Get(dto)
	if err != nil {
		return nil, err
	}

	if existe {
		return dto, nil
	}

	return nil, nil
}
