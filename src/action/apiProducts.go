package action

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"model"
	"net/http"
)

func Ping(){
	e := DB.Ping()
	if e!=nil {
		fmt.Println("loi connect mysql api")
	}
	fmt.Println("ok",DB)
	res, _ := DB.Query("SHOW TABLES")
	defer res.Close()
	var table string
	for res.Next() {
		res.Scan(&table)
		fmt.Println(table)
	}
}
func GetProducts(w http.ResponseWriter, r *http.Request){
	logrus.Info("Request get list Products")
	res, err := DB.Query(`SELECT product_id,product_name,product_price,product_info 
		FROM products`)
	defer res.Close()
	if err != nil {
		logrus.Error("Statement query has mistake")
		return
	}
	var listProducts []model.Product
	for res.Next() {
		var (
			id      int
			name    string
			price   float32
			info 	string
		)
		err = res.Scan(&id, &name, &price, &info)
		if err != nil {
			logrus.Error("Scan result query has mistake")
		}
		product := model.Product{ProductId: id, ProductName: name, ProductPrice: price, ProductInfo: info}
		listProducts = append(listProducts, product)
	}
	err = json.NewEncoder(w).Encode(listProducts)
	if err != nil {
		logrus.Error("respond json list students has mistake")
	}
	logrus.Info("Responded list Students cuccess")
}

func GetDetailProduct(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Request get detail Product")
	para := mux.Vars(r)
	res, err := DB.Query(`SELECT product_id,product_name,product_price,product_info 
		FROM products WHERE product_id = ?`,para["id"])
	defer res.Close()
	if err != nil {
		logrus.Error("Statement get query has mistake")
		return
	}
	var product model.Product
	if res.Next() {
		var (
			id       int
			name     string
			price    float32
			info	 string
		)
		err = res.Scan(&id, &name, &price, &info)
		if err != nil {
			logrus.Error("Scan result query has mistake")
		}
		product = model.Product{ProductId: id,  ProductName: name, ProductPrice: price, ProductInfo: info}
	} else {
		logrus.Info("No found from request")
		_ = json.NewEncoder(w).Encode(http.StatusBadRequest)
		return
	}
	logrus.Info("Responded a Students")
	_ = json.NewEncoder(w).Encode(product)
}
func CreateNewProduct(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Request post Product")
	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err!=nil{
		logrus.Error("Request body has mistake")
		_=json.NewEncoder(w).Encode(http.StatusBadRequest)
	}
	res, err := DB.Query(`INSERT INTO products (product_name,product_price,product_info) 
		VALUES (?,?,?)`,product.ProductName,product.ProductPrice,product.ProductInfo)
	defer res.Close()
	if err != nil {
		logrus.Error("Statement get query has mistake")
		return
	}

	logrus.Info("Responded post a product cuccess")
	_ = json.NewEncoder(w).Encode(http.StatusOK)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Request put detail Product")
	para := mux.Vars(r)
	res, err := DB.Query(`SELECT product_id 
		FROM products WHERE product_id = ?`,para["id"])
	defer res.Close()
	if err != nil {
		logrus.Error("Statement get query has mistake")
		return
	}
	var product model.Product
	if res.Next() {
		err = json.NewDecoder(r.Body).Decode(&product)
		if err != nil||product.ProductName=="" {
			logrus.Error("Request body has mistake")
		}
		res2, err := DB.Query(`UPDATE products SET product_name=?,product_price=?,product_info=? 
		WHERE product_id = ?`,product.ProductName,product.ProductPrice,product.ProductInfo,para["id"])
		defer res2.Close()
		if err != nil {
			logrus.Error("Statement get query has mistake")
			return
		}
	} else {
		logrus.Info("No found product id from request")
		_ = json.NewEncoder(w).Encode(http.StatusBadRequest)
		return
	}
	logrus.Info("Responded update product id = ",para["id"]," cuccess")
	_ = json.NewEncoder(w).Encode(http.StatusOK)
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Request delete Product")
	para := mux.Vars(r)
	res, err := DB.Query(`SELECT product_id FROM products WHERE product_id=?`,para["id"])
	defer res.Close()
	if err != nil {
		logrus.Error("Statement get query has mistake")
		return
	}
	if res.Next(){
		res2, err := DB.Query(`DELETE FROM products WHERE product_id=?`,para["id"])
		defer res2.Close()
		if err != nil {
			logrus.Error("Statement get query has mistake")
			return
		}
		logrus.Info("Responded delete a product id= ", para["id"]," cuccess")
		_ = json.NewEncoder(w).Encode(http.StatusOK)
		return
	}
	logrus.Info("Request not found id")
	_ = json.NewEncoder(w).Encode(http.StatusBadRequest)
}


