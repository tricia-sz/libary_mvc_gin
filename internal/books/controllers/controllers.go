package controllers

import "github.com/gin-gonic/gin"

type BooksController struct {
}

func NewBooksController() *BooksController {
	return &BooksController{}
}

func (c *BooksController) RegisterRouter(r *gin.Engine) {
	Books := r.Group("/Bookss")

	{
		Books.POST("", c.CreateBook)
		Books.GET("/:id", c.GetBook)
		Books.GET("", c.GetAllBooks)
		Books.PUT("/:id", c.UpdateBook)
		Books.DELETE("/:id", c.DeleteBook)

	}
}

func (c *BooksController) CreateBook(ctx *gin.Context) {

}

func (c *BooksController) GetBook(ctx *gin.Context) {
	ctx.String(200, "FunUNciou")

}

func (c *BooksController) GetAllBooks(ctx *gin.Context) {

}

func (c *BooksController) UpdateBook(ctx *gin.Context) {

}

func (c *BooksController) DeleteBook(ctx *gin.Context) {

}
