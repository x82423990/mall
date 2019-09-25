package control

import (
	"context"
	"dumall/common"
	"dumall/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
)

func ListGoods(c *gin.Context) {
	//var sortFlag string
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	PageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	sort, _ := strconv.Atoi(c.DefaultQuery("sort", "-1"))
	min, _ := strconv.Atoi(c.DefaultQuery("searchPriceMin", "0"))
	max, _ := strconv.Atoi(c.DefaultQuery("searchPriceMax", "0"))
	//if sort == 1 {
	//	sortFlag = "salePrice"
	//} else {
	//	sortFlag = "productName"
	//}
	var filter primitive.D

	if max != 0 {
		//filter = bson.D{
		//	//{"salePrice",bson.D{{"$lt", min}}},
		//	{"status", "A"},
		//	{"qty", bson.D{{"$lt", 30}}},
		//
		//}
		filter = bson.D{
			{"salePrice", bson.D{{"$lt", max}}},
			{"salePrice", bson.D{{"$gt", min}}},
		}
	} else {
		filter = bson.D{}

	}
	fmt.Println(filter)
	col := MongoCli.Collection("goods")
	var goods []models.Good

	var opts options.FindOptions
	skip := int64((page - 1) * PageSize)
	limit := int64(PageSize)
	sortField := make(map[string]interface{})
	// -1 是升序 , 1 是降序.
	sortField["salePrice"] = sort
	opts.Skip = &skip
	opts.Sort = sortField
	opts.Limit = &limit
	cur, err := col.Find(context.Background(), filter, &opts)
	count, _ := col.CountDocuments(context.Background(), filter)
	//col.CountDocuments()
	if err != nil {
		common.ResErrSrv(c, err)
	}
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var good models.Good
		err := cur.Decode(&good)
		if err != nil {
			log.Fatal(err)
		}
		goods = append(goods, good)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	common.ResSuccessPage(c, int(count), goods)
}
