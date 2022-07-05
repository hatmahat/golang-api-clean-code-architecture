package controller

import (
	"go-api-with-gin2/delivery/api"
	"go-api-with-gin2/model"
	"go-api-with-gin2/usecase"
	"go-api-with-gin2/utils"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	router        *gin.Engine
	ucProduct     usecase.CreateProductUseCase
	UcProductList usecase.ListProductUseCase
	api.BaseApi
}

// func (p *ProductController) createNewProduct(c *gin.Context) {
// 	var newProduct *model.Product
// 	if err := c.BindJSON(&newProduct); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  "BAD REQUEST",
// 			"message": err.Error(),
// 		})
// 	} else {
// 		err := p.ucProduct.CreateProduct(newProduct)
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 				"status":  "FAILED",
// 				"message": "Error when creating a product",
// 			})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{
// 			"status":  "SUCCESS",
// 			"message": newProduct,
// 		})
// 	}
// }

func (p *ProductController) createNewProduct(c *gin.Context) {
	var newProduct model.Product
	err := p.ParseRequestBody(c, &newProduct)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	err = p.ucProduct.CreateProduct(&newProduct)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, newProduct)
}

// func (p *ProductController) findAllProduct(c *gin.Context) {
// 	product, err := p.UcProductList.Retrive()
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"status":  "FAILED",
// 			"message": "Error when retrive all products",
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  "SUCCESS",
// 		"message": product,
// 	})
// }

func (p *ProductController) findAllProduct(c *gin.Context) {
	products, err := p.UcProductList.Retrive()
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, products)
}

func NewProductController(
	router *gin.Engine,
	ucProduct usecase.CreateProductUseCase,
	UcProductList usecase.ListProductUseCase) *ProductController {
	var controller ProductController = ProductController{
		router:        router,
		ucProduct:     ucProduct,
		UcProductList: UcProductList,
	}

	router.POST("/product", controller.createNewProduct)
	router.GET("/product", controller.findAllProduct)
	return &controller
}
