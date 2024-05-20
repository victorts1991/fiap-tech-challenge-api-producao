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

var _ = Describe("pesquisa por pedido id use case testes", func() {
	var (
		ctx             = context.Background()
		repo            *mock_repo.MockFilaRepo
		pegaPorPedidoUC PegaPedidoPorID
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		repo = mock_repo.NewMockFilaRepo(mockCtrl)
		pegaPorPedidoUC = NewPegaPedidoPorID(repo)
	})

	Context("pesquisa pedido na fila", func() {
		objID := primitive.NewObjectID()
		It("pesquisa com sucesso", func() {
			repo.EXPECT().PegaPedidoPorID(ctx, objID.Hex()).Return(&domain.Producao{}, nil)
			ped, err := pegaPorPedidoUC.PesquisaPorID(ctx, objID.Hex())

			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(ped).ToNot(gomega.BeNil())
		})
		It("falha ao pesquisar por pedido id", func() {
			repo.EXPECT().PegaPedidoPorID(ctx, objID.Hex()).Return(nil, errors.New("mock error"))
			ped, err := pegaPorPedidoUC.PesquisaPorID(ctx, objID.Hex())

			gomega.Expect(err.Error()).To(gomega.Equal("mock error"))
			gomega.Expect(ped).To(gomega.BeNil())
		})
		It("falha ao pesquisar por pedido id retorna nil", func() {
			repo.EXPECT().PegaPedidoPorID(ctx, objID.Hex()).Return(nil, nil)
			ped, err := pegaPorPedidoUC.PesquisaPorID(ctx, objID.Hex())

			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(ped).To(gomega.BeNil())
		})
	})
})
