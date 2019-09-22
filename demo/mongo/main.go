package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"reflect"
	"time"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}
type Users struct {
	UserID       string        `bson:"userId"`
	UserName     string        `json:"userName"`
	UserPwd      string        `json:"userPwd"`
	OrderLists   []OderList    `bson:"orderList"`
	CartLists    []CartList    `json:"cartList"`
	AddressLists []AddressList `json:"addressList"`
}
type OderList struct {
	OrderID     string      `bson:"orderId"`
	OrderTotal  interface{} `json:", int"`
	AddressInfo `json:"addressInfo,omitempty"`
	GoodsLists  []GoodList  `json:"goodsList"`
	OrderStatus interface{} `json:"orderStatus"`
	CreateDate  string      `json:"createDate"`
}
type AddressList struct {
	AddressID  string `json:"addressId"`
	UserName   string `json:"userName"`
	StreetName string `json:"streetName"`
	PostCode   string `json:"postCode"`
	Tel        string `json:"tel"`
	IsDefault  bool   `json:"isDefault"`
}
type CartList struct {
	ProductImage string `json:"productImage"`
	SalePrice    string `json:"salePrice"`
	ProductName  string `json:"productName"`
	ProductID    string `json:"productId"`
	ProductNum   string `json:"productNum"`
	Checked      string `json:"checked,omitempty"`
}
type AddressInfo struct {
	AddressID  string      `json:"addressId"`
	UserName   string      `json:"userName"`
	StreetName string      `json:"streetName"`
	PostCode   interface{} `json:"postCode"`
	Tel        float64     `json:"tel"`
	IsDefault  bool        `json:"isDefault"`
}
type GoodList struct {
	ProductImage string `json:"productImage"`
	SalePrice    string `json:"salePrice"`
	ProductName  string `json:"productName"`
	ProductID    string `json:"productId"`
	ProductNum   string `json:"productNum"`
	Checked      string `json:"checked"`
}

func main() {
	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.44.128:27017"))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// 超时设置
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.44.128:27017"))
	if err != nil {
		return
	}
	collection := client.Database("dumall").Collection("users")
	filter := bson.M{"userId": "100000077"}
	var user Users
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(reflect.TypeOf(user.OrderLists[0].OrderTotal))
}
