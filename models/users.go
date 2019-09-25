package models

type Users struct {
	UserID       string    `bson:"userId"`
	UserName     string    `json:"userName"`
	UserPwd      string    `json:"userPwd"`
	OrderLists   []Oder    `bson:"orderList"`
	CartLists    []Cart    `json:"cartList"`
	AddressLists []Address `json:"addressList"`
}
type Oder struct {
	OrderID     string      `bson:"orderId"`
	OrderTotal  interface{} `json:", int"`
	AddressInfo `json:"addressInfo,omitempty"`
	GoodsLists  []Good      `json:"goodsList"`
	OrderStatus interface{} `json:"orderStatus"`
	CreateDate  string      `json:"createDate"`
}
type Address struct {
	AddressID  string `json:"addressId"`
	UserName   string `json:"userName"`
	StreetName string `json:"streetName"`
	PostCode   string `json:"postCode"`
	Tel        string `json:"tel"`
	IsDefault  bool   `json:"isDefault"`
}
type Cart struct {
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
type Good struct {
	ProductImage string      `bson:"productImage"`
	SalePrice    interface{} `json:"salePrice"`
	ProductName  string      `json:"productName"`
	ProductID    string      `json:"productId"`
	ProductNum   string      `json:"productNum"`
	Checked      string      `json:"checked"`
}
