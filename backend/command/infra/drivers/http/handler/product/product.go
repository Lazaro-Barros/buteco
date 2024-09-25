package handler

import (
	"net/http"

	product_service "github.com/Lazaro-Barros/buteco/command/application/product"
	"github.com/Lazaro-Barros/buteco/command/infra/drivers/http/handler"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService *product_service.ProductService
}

// NewProductHandler cria um novo ProductHandler.
func NewProductHandler(service *product_service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: service,
	}
}

// CreateProductHandler lida com a requisição de criação de um novo produto.
func (h *ProductHandler) CreateProductHandler(c handler.Context) {
	var createProductDTO product_service.CreateProductDTO

	// Valida os dados de entrada do JSON.
	if err := c.ShouldBindJSON(&createProductDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chama o serviço para criar o produto.
	uuid, err := h.productService.CreateProduct(createProductDTO.Name, createProductDTO.Description, createProductDTO.Value, createProductDTO.Visible)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"uuid": uuid})
}

// ListProductsHandler lida com a requisição para listar todos os produtos.
func (h *ProductHandler) ListProductsHandler(c handler.Context) {
	
	// Chama o serviço para listar os produtos.
	products, err := h.productService.ListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

// UpdateProductHandler lida com a requisição de atualização de um produto existente.
func (h *ProductHandler) UpdateProductHandler(c handler.Context) {
	var updateProductDTO product_service.UpdateProductDTO

	// Valida os dados de entrada do JSON.
	if err := c.ShouldBindJSON(&updateProductDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chama o serviço para atualizar o produto.
	if err := h.productService.UpdateProduct(c.Param("uuid"), updateProductDTO.Name, updateProductDTO.Description, updateProductDTO.Value, updateProductDTO.Visible); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// DeleteProductHandler lida com a exclusão de um produto.
func (h *ProductHandler) DeleteProductHandler(c handler.Context) {
	// Chama o serviço para excluir o produto.
	if err := h.productService.DeleteProduct(c.Param("uuid")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
