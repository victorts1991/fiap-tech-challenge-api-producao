package usecase

import (
	"context"
	"errors"
	"fiap-tech-challenge-producao/internal/core/domain"
	mock_repo "fiap-tech-challenge-producao/test/mock/repository"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

var _ = Describe("atualiza status do pedido na producao use case testes", func() {
	var (
		ctx              = context.Background()
		repo             *mock_repo.MockFilaRepo
		atualizaStatusUC AtualizaStatusProducao
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		repo = mock_repo.NewMockFilaRepo(mockCtrl)
		atualizaStatusUC = NewAtualizaStatusProducaoUC(repo)
	})

	Context("atualiza pedido na fila", func() {
		objID := primitive.NewObjectID()
		It("atualiza com sucesso", func() {
			repo.EXPECT().PegaPedidoPorID(ctx, objID.Hex()).Return(&domain.Producao{}, nil)
			repo.EXPECT().AtualizaStatus(ctx, "pronto", objID.Hex()).Return(nil)
			err := atualizaStatusUC.Atualiza(ctx, "pronto", objID.Hex())

			gomega.Expect(err).To(gomega.BeNil())
		})
		It("falha ao atualiza", func() {
			repo.EXPECT().AtualizaStatus(ctx, "pronto", objID.Hex()).Return(errors.New("mock error"))
			repo.EXPECT().PegaPedidoPorID(ctx, objID.Hex()).Return(&domain.Producao{}, nil)
			err := atualizaStatusUC.Atualiza(ctx, "pronto", objID.Hex())

			gomega.Expect(err.Error()).To(gomega.Equal("mock error"))
		})
		It("falha ao pesquisar por pedido id", func() {
			repo.EXPECT().PegaPedidoPorID(ctx, objID.Hex()).Return(nil, errors.New("mock error"))
			err := atualizaStatusUC.Atualiza(ctx, "pronto", objID.Hex())

			gomega.Expect(err.Error()).To(gomega.Equal("mock error"))
		})
		It("falha ao pesquisar por pedido id retorna nil", func() {
			repo.EXPECT().PegaPedidoPorID(ctx, objID.Hex()).Return(nil, nil)
			err := atualizaStatusUC.Atualiza(ctx, "pronto", objID.Hex())

			gomega.Expect(err).To(gomega.Not(gomega.BeNil()))
		})
	})
})
