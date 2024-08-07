package handler

import (
	"github.com/gofiber/fiber/v2"
	"projectSwagger/internal/model"
	"projectSwagger/internal/pkg/service"
)

type ProductHandler struct {
	ProductService service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: service,
	}
}

// GetAllProductsHandler GetProduct getting product
//
//	@Summary		Getting products
//	@Description	Getting products with given request
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string
//	@Failure		400	{string}	string	"Bad Request"
//	@Router			/products/all [get]
func (p *ProductHandler) GetAllProductsHandler(c *fiber.Ctx) error {
	products, err := p.ProductService.GetAllProducts()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": products,
	})
}

// BuyProductHandler BuyProduct buying product
//
//	@Summary		Buying product
//	@Description	Buying product with given request
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.BuyProductBody	true	"Request of Buying Product Object"
//	@Success		200		{string}	string
//	@Failure		400		{string}	string	"Bad Request"
//	@Router			/products/buy [post]
func (p *ProductHandler) BuyProductHandler(c *fiber.Ctx) error {
	payload := new(model.BuyProductBody)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Payload error",
			"error":   err.Error(),
		})
	}

	response, err := p.ProductService.BuyProduct(payload)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Can't buy product",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": response,
	})
}
