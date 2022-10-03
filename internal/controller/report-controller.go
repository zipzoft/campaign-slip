package controller

import (
	"campiagn-slip/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReportController struct {
	repo repository.ReportRepository
}

func NewReportController(repo repository.ReportRepository) *ReportController {
	return &ReportController{repo: repo}
}

func (r ReportController) ReportUserRedeem(c *gin.Context) {

	result, err := r.repo.ReportTransaction(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, result)
}
