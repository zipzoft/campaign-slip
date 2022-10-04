package controller

import (
	"campiagn-slip/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ReportController struct {
	repo repository.ReportRepository
}

func NewReportController(repo repository.ReportRepository) *ReportController {
	return &ReportController{repo: repo}
}

func (r ReportController) ReportUserRedeem(c *gin.Context) {

	Page, _ := strconv.Atoi(c.Query("page"))
	PerPage, _ := strconv.Atoi(c.Query("limit"))
	result, err := r.repo.ReportTransaction(c, Page, PerPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}
