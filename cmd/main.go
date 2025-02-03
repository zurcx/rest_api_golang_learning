package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zurcx/rest_api_golang_learning/controller"
	"github.com/zurcx/rest_api_golang_learning/usecase"
)
func main() {
	// Inicializa o servidor Gin
	server := gin.Default()

    // Camada de UseCase
    ProductUseCase := usecase.NewProductUseCase()
	// Cria uma instância do controlador de produtos
	ProductController := controller.NewProductController(ProductUseCase)

	// Rota simples de teste
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Pong!",
		})
	})

	// Certifique-se de que a função GetProducts está corretamente implementada
	server.GET("/products", ProductController.GetProducts)

	// Inicia o servidor na porta 8000
	server.Run(":8000")
}
