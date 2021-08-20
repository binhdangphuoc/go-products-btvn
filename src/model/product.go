package model
/* create table product (
product_id int NOT NULL AUTO_INCREMENT,
product_name varchar(50) NOT NULL,
product_price double NOT NULL,
product_info varchar(100),
PRIMARY KEY (product_id)
	);*/
type Product struct {
	ProductId int 	`json:"product_id"`
	ProductName string 	`json:"product_name"`
	ProductPrice float32 `json:"product_price"`
	ProductInfo string `json:"product_info"`
}
/* create table orders (
order_id int NOT NULL AUTO_INCREMENT,
order_price double,
user_id int NOT NULL,
order_date Date NOT NULL,
ship_date Date,
PRIMARY KEY (order_id),
FOREIGN KEY (user_id) REFERENCES user(user_id)
	);*/

/*create table order_detail (
order_id int NOT NULL,
product_id int NOT NULL,
amount int NOT NUll,
price_products double,
FOREIGN KEY (product_id) REFERENCES products(product_id),
PRIMARY KEY (order_id)
); */
type Order struct {
	OrderId int `json:"order_id"`
	OrderDetail []ProductInOrder `json:"order_detail"`
	OrderPrice float32 `json:"order_price"`
	OrderDate string `json:"order_date"`
	ShipDate string `json:"ship_date"`
}
type ProductInOrder struct {
	ProductItem Product
	Amount int
}
/*create table payment_detail (
payment_id int NOT NULL auto_increment,
order_id int NOT NULL,
user_id int NOT NULL,
total double NOT NUll,
payment_date Date,
PRIMARY KEY (payment_id),
FOREIGN KEY (order_id) REFERENCES orders(order_id),
FOREIGN KEY (user_id) REFERENCES users(user_id)
); */
type PaymentDetail struct {
	UserId int `json:"user_id"`
	OrderId int `json:"order_id"`
	Total float32 `json:"total"`
	PaymentDate string `json:"payment_date"`
}
