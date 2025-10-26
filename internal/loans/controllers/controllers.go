package controllers

import "github.com/gin-gonic/gin"

type LoansController struct {
}

func NewLoansController() *LoansController {
	return &LoansController{}
}

func (c *LoansController) RegisterRouter(r *gin.Engine) {
	Loans := r.Group("/loans")

	{
		Loans.POST("", c.CreateLoan)
		Loans.GET("/:id", c.GetLoan)
		Loans.GET("", c.GetAllLoans)
		Loans.PUT("/:id", c.UpdateLoan)
		Loans.DELETE("/:id", c.DeleteLoan)

	}
}

func (c *LoansController) CreateLoan(ctx *gin.Context) {

}

func (c *LoansController) GetLoan(ctx *gin.Context) {
	ctx.String(200, "FunUNciou")

}

func (c *LoansController) GetAllLoans(ctx *gin.Context) {

}

func (c *LoansController) UpdateLoan(ctx *gin.Context) {

}

func (c *LoansController) DeleteLoan(ctx *gin.Context) {

}
