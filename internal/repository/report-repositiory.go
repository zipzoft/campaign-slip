package repository

import (
	"campiagn-slip/pkg/database"
	"encoding/json"
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
			"data":       bson.M{"$push": "$$ROOT"},
			"total_coin": bson.M{"$sum": "$coin"},
		},
		},
		bson.M{"$project": bson.M{"data.prefix": 1, "data.coin": 1, "data.date_bank": 1, "data.slip_number": 1, "data.is_redeem": 1, "data.created_at": 1, "total_coin": 1}},
		bson.M{"$facet": bson.M{
			"metadata": []bson.M{
				{
					"$count": "total",
				},
				{
					"$addFields": bson.M{"page": Page},
				},
			},
			"data": []bson.M{
				{
					"$sort": bson.M{"_id": 1},
				},
				{
					"$skip": (Page - 1) * PerPage,
				},
				{
					"$limit": PerPage,
				},
				{
					"$addFields": bson.M{"username": "$_id"},
				},
			},
		}},
		bson.M{"$project": bson.M{
			"_id": false,
			"total": bson.M{
				"$arrayElemAt": []interface{}{"$metadata.total", 0},
			},
			"page": bson.M{
				"$arrayElemAt": []interface{}{"$metadata.page", 0},
			},
			"data": 1,
		}},
	}
	//var report []models.ReportData
	ins, err := database.AggregatePagination("user_redeem", Aggregate)
	response := bson.M{}
	b, _ := json.Marshal(ins)
	s := string(b)
	s = strings.TrimSpace(s)
	s = s[1 : len(s)-1]
	_ = json.Unmarshal([]byte(s), &response)

	return response, err
}
