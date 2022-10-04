package repository

import (
	"campiagn-slip/models"
	"campiagn-slip/pkg/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

var _ ReportRepository = (*ReportRepo)(nil)

type ReportRepository interface {
	ReportTransaction(ctx *gin.Context, Page, PerPage int) (interface{}, error)
}
type ReportRepo struct {
	//
}

func (r ReportRepo) ReportTransaction(ctx *gin.Context, Page, PerPage int) (interface{}, error) {
	dateFrom := ctx.Query("date_from")
	dateTo := ctx.Query("date_to")
	Prefix := ctx.Query("prefix")
	ArrPrefix := strings.Split(Prefix, ",")
	Aggregate := bson.A{
		bson.M{"$match": bson.M{"$and": bson.A{
			bson.M{"date_bank": bson.M{"$gte": dateFrom}},
			bson.M{"date_bank": bson.M{"$lte": dateTo}},
			bson.M{"prefix": bson.M{"$in": ArrPrefix}},
			bson.M{"is_redeem": true},
		}}},
		bson.M{"$group": bson.M{"_id": "$username",
			"data":  bson.M{"$push": "$$ROOT"},
			"count": bson.M{"$count": bson.M{}},
		},
		},
		bson.M{"$project": bson.M{"data.prefix": 1, "data.coin": 1, "data.date_bank": 1, "data.slip_number": 1, "data.is_redeem": 1, "data.created_at": 1, "count": 1}},
		bson.M{"$sort": bson.M{"_id": 1}},
		bson.M{"$skip": (Page - 1) * PerPage},
		bson.M{"$limit": PerPage},
	}
	var ins []models.ReportData
	err := database.Aggregate("user_redeem", Aggregate, &ins)

	return ins, err
}
