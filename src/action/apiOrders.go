package action

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"model"
	"net/http"
)

func GetOrders(w http.ResponseWriter, r *http.Request){
	logrus.Info("Request get all Orders")
	var listOrders []model.Order
	dbOrder,err := DB.Query(`SELECT order_id,user_id,order_price,order_date,ship_date FROM orders`)
	defer dbOrder.Close()
	if err!=nil{
		logrus.Error("Statement get query order has mistake")
	}
	for dbOrder.Next(){
		var (
			order_id int
			user_id int
			order_price float32
			order_date string
			ship_date string
		)
		err = dbOrder.Scan(&order_id,&user_id,&order_price,&order_date,&ship_date)
		if err!=nil{
			logrus.Error("Scan database order has mistake")
		}
		dbOrderDetail, err := DB.Query(`SELECT product_id,amount FROM order_detail WHERE order_id=?`,order_id)
		defer dbOrderDetail.Close()
		if err!=nil{
			logrus.Error("Statement get query order detail has mistake")
		}
		var listProducts []model.ProductInOrder
		for dbOrderDetail.Next(){
			fmt.Println("dbOrderDetail oke")
			var (
				product_id int
				amount int
			)
			err = dbOrderDetail.Scan(&product_id,&amount)
			if err!=nil{
				logrus.Error("Scan database order detail has mistake")
			}
			dbProduct, err := DB.Query(`SELECT product_name,product_price,product_info 
			FROM products WHERE product_id=?`,product_id)
			defer dbProduct.Close()
			if err!=nil{
				logrus.Error("Statement get query product in order has mistake")
			}
			if dbProduct.Next(){
				var (
					product_name string
					product_price float32
					product_info string
				)

				err = dbProduct.Scan(&product_name,&product_price,&product_info)
				if err!=nil{
					logrus.Error("Scan database product in order has mistake")
				}
				//fmt.Println(product_id,product_name,product_price,product_info,amount)
				pro := model.Product{ProductId: product_id,ProductName: product_name,
					ProductPrice: product_price,ProductInfo: product_info}
				listProducts = append(listProducts,model.ProductInOrder{pro,amount})
			}

		}

		order:= model.Order{OrderId: order_id,OrderDetail: listProducts, OrderPrice: order_price,
			OrderDate: order_date, ShipDate: ship_date}
		listOrders = append(listOrders,order)
	}

	_ = json.NewEncoder(w).Encode(listOrders)
}

func GetOrderDetail(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Request get detail a Order")
	var order model.Order
	para := mux.Vars(r)
	dbOrder, err := DB.Query(`SELECT order_id,user_id,order_price,order_date,ship_date FROM orders
	WHERE order_id=?`,para["id"])
	defer dbOrder.Close()
	if err!=nil{
		logrus.Error("Query find order has mistake")
	}
	if dbOrder.Next(){
		var (
			order_id int
			user_id int
			order_price float32
			order_date string
			ship_date string
		)
		err = dbOrder.Scan(&order_id,&user_id,&order_price,&order_date,&ship_date)
		if err!=nil{
			logrus.Error("Scan database order has mistake")
		}
		dbOrderDetail, err := DB.Query(`SELECT product_id,amount FROM order_detail WHERE order_id=?`,order_id)
		defer dbOrderDetail.Close()
		if err!=nil{
			logrus.Error("Statement get query order detail has mistake")
		}
		var listProducts []model.ProductInOrder
		for dbOrderDetail.Next(){
			fmt.Println("dbOrderDetail oke")
			var (
				product_id int
				amount int
			)
			err = dbOrderDetail.Scan(&product_id,&amount)
			if err!=nil{
				logrus.Error("Scan database order detail has mistake")
			}
			dbProduct, err := DB.Query(`SELECT product_name,product_price,product_info 
			FROM products WHERE product_id=?`,product_id)
			defer dbProduct.Close()
			if err!=nil{
				logrus.Error("Statement get query product in order has mistake")
			}
			if dbProduct.Next(){
				var (
					product_name string
					product_price float32
					product_info string
				)

				err = dbProduct.Scan(&product_name,&product_price,&product_info)
				if err!=nil{
					logrus.Error("Scan database product in order has mistake")
				}
				//fmt.Println(product_id,product_name,product_price,product_info,amount)
				pro := model.Product{ProductId: product_id,ProductName: product_name,
					ProductPrice: product_price,ProductInfo: product_info}
				listProducts = append(listProducts,model.ProductInOrder{pro,amount})
			}

		}

		order= model.Order{OrderId: order_id,OrderDetail: listProducts, OrderPrice: order_price,
			OrderDate: order_date, ShipDate: ship_date}
		_ = json.NewEncoder(w).Encode(order)
		return
	}else{
		_ = json.NewEncoder(w).Encode(http.StatusBadRequest)
	}


}