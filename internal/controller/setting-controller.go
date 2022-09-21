package controller

import (
	"campiagn-slip/internal/repository"
	"campiagn-slip/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type SettingController struct {
	repo repository.SettingRepository
}

func NewSettingController(repo repository.SettingRepository) *SettingController {
	return &SettingController{repo: repo}
}

func (ctrl *SettingController) InsertAndUpdateCondition(c *gin.Context) {
	condition := models.Condition{}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	json.Unmarshal(body, &condition)
	if c.Request.Method == "POST" {

		condition, err = ctrl.repo.InsertCondition(condition)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		result := map[string]interface{}{
			"data": condition,
			"msg":  "success",
		}
		c.JSON(http.StatusOK, result)

	}
	if c.Request.Method == "PATCH" {

		err = ctrl.repo.UpdateCondition(condition, c)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, "success")
	}
}
func (ctrl *SettingController) Condition(c *gin.Context) {
	prefix := c.Query("prefix")
	condition, err := ctrl.repo.FindCondition(prefix)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, condition)

}
