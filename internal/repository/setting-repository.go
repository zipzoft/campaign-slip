package repository

import (
	"campiagn-slip/models"
	"campiagn-slip/pkg/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

var _ SettingRepository = (*SettingRepo)(nil)

type SettingRepository interface {
	InsertCondition(condition models.Condition) (models.Condition, error)
	UpdateCondition(condition models.Condition, c *gin.Context) error
	FindCondition(c *gin.Context) (result interface{}, err error)
	FindOneCondition(prefix string) (model models.Condition, err error)
}

type SettingRepo struct {
	//
}

func (r SettingRepo) InsertCondition(condition models.Condition) (models.Condition, error) {
	checkCondition := models.Condition{}
	database.FindOne("condition", bson.M{"prefix": condition.Prefix}).Decode(&checkCondition)
	if checkCondition.Prefix == "" {
		condition.ID = primitive.NewObjectID()
		condition.CreatedAt = time.Now()
		_, err := database.InsertOne("condition", condition)
		if err != nil {
			return models.Condition{}, err
		}
		err = database.FindOne("condition", bson.M{"_id": condition.ID}).Decode(&condition)
		return condition, err
	}
	return checkCondition, fmt.Errorf("prefix ถูกใช้งานแล้ว")
}
func (r SettingRepo) UpdateCondition(condition models.Condition, c *gin.Context) error {

	id, err := primitive.ObjectIDFromHex(c.Query("id"))
	if err != nil {
		return err
	}
	condition.ID = id
	_, err = database.UpdateOne("condition", bson.M{"_id": condition.ID}, condition)

	return err
}
func (r SettingRepo) FindCondition(c *gin.Context) (result interface{}, err error) {
	prefix := c.Query("prefix")
	page := c.Query("page")
	limit := c.Query("limit")
	model := make([]models.Condition, 0)
	filter := bson.M{"prefix": strings.ToLower(prefix)}
	if prefix == "" {
		delete(filter, "prefix")
	}
	result, err = database.Pagination("condition", filter, &model, page, limit)

	if err != nil {
		return nil, err
	}
	return result, nil
}
func (r SettingRepo) FindOneCondition(prefix string) (model models.Condition, err error) {
	err = database.FindOne("condition", bson.M{"prefix": strings.ToLower(prefix)}).Decode(&model)
	return model, err
}
