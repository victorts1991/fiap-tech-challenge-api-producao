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

var _ = Describe("adiciona pedido a producao use case testes", func() {
	var (
		ctx            = context.Background()
		repo           *mock_repo.MockFilaRepo
		cadastraFilaUC CadastrarFila
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		repo = mock_repo.NewMockFilaRepo(mockCtrl)
		cadastraFilaUC = NewCadastraFila(repo)
	})

	Context("cadastra na fila", func() {
		objID := primitive.NewObjectID()
		fila := &domain.Producao{
			Id:         1,
			PedidoId:   objID.Hex(),
			Status:     "pronto",
			Observacao: "teste",
		}
		It("cadastra com sucesso", func() {
			repo.EXPECT().Insere(ctx, fila).Return(nil)
			err := cadastraFilaUC.Cadastra(ctx, fila)

			gomega.Expect(err).To(gomega.BeNil())
		})
		It("falha ao cadastrar", func() {
			repo.EXPECT().Insere(ctx, fila).Return(errors.New("mock error"))
			err := cadastraFilaUC.Cadastra(ctx, fila)

			gomega.Expect(err.Error()).To(gomega.Equal("mock error"))
		})
	})
})
