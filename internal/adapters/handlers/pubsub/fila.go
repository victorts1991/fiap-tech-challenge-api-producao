package pubsub

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fiap-tech-challenge-producao/internal/core/domain"
	"fiap-tech-challenge-producao/internal/core/usecase"
	"github.com/rhuandantas/fiap-tech-challenge-commons/pkg/messaging"
	"log"
	"strings"
)

type ProducaoHandler struct {
	fila usecase.CadastrarFila
}

func NewProducaoHandler(fila usecase.CadastrarFila) *ProducaoHandler {
	return &ProducaoHandler{fila: fila}
}

func (h *ProducaoHandler) AdicionaProducaoHandler(ctx context.Context, message string, attr messaging.MessageAttr) error {
	var producao domain.Producao
	message = strings.Trim(message, `"`)
	decodedMessage, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return err
	}
	// Unmarshal the message into the domain object
	if err := json.Unmarshal(decodedMessage, &producao); err != nil {
		log.Printf("failed to unmarshal message: %s, error: %v", message, err)
		return err
	}

	log.Println("Message received ", producao)

	// Call the usecase to handle the business logic
	if err := h.fila.Cadastra(ctx, &producao); err != nil {
		log.Printf("error happened while processing message: %v", err)
		return err
	}
	return nil
}
