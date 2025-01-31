package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func raizQuadrada(c *fiber.Ctx) error {
	opStr := c.Params("op")
	op, err := strconv.ParseFloat(opStr, 64)

	// Checando parâmetro.
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(fmt.Sprintf("Parâmetro Inválido:\":%s\"", opStr))
	}

	// Calculando raiz quadrada.
	raiz := math.Sqrt(op)

	// formatando retorno para ter 2 casas decimais e enviando para o cliente.
	return c.SendString(fmt.Sprintf("%.2f", raiz))
}
