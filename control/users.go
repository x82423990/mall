package control

import (
	"context"
	"dumall/common"
	"dumall/models"
	"dumall/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoCli *mongo.Database

func init() {
	cli := utils.MongoCli()
	cli.Database("dumall")
	MongoCli = cli.Database("dumall")
}

func ListUsers(c *gin.Context) {
	col := MongoCli.Collection("users")
	filter := bson.M{"userId": "100000077"}
	var user models.Users
	err := col.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		common.ResErrSrv(c, err)
	}
	common.ResJSON(c, 200, user)
}
