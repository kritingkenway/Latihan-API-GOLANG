package model

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB

type User struct {
	ID	int	`gorm:"primaryKey;autoIncrement"`

	Nama string `gorm:"not null"`

	Umur int `gorm:"not null"`

	Email *string

	Password string `gorm:"not null"`

	Cart Cart

	Order []Order
}

type UserToken struct {
	Token string `gorm:"unique"`
	UserID int `gorm:"primaryKey;unique"`
	User User
}

type Product struct {
	ID int `form:"id" gorm:"primaryKey;autoIncrement"`
	Name string `form:"nama" gorm:"not null"`
	Harga int `form:"harga"`
	Stock int `form:"stock"`
	

}


type Cart struct {
	ID uint `form:"id" gorm:"primaryKey"`
	UserID uint `form:"user_id" gorm:"not null"`
	CreatedAt time.Time
	CartItem []CartItem `gorm:"constraint:OnDelete:CASCADE;"`
	
}


type CartItem struct {
	ID uint `form:"id" gorm:"primaryKey;autoIncrement"`
	Qty uint `form:"qty" gorm:"not null"`
	ProductID uint `form:"product_id" gorm:"not null"`	
	Product Product
	CartID uint `form:"cart_id" gorm:"not null"`

	
}

type Order struct {
	ID uint `form:"id" gorm:"primaryKey"`
	UserID uint `form:"user_id" gorm:"not null"`
	CreatedAt time.Time
	OrderItem []OrderItem `gorm:"constraint:OnDelete:CASCADE;"`
	
}


type OrderItem struct {
	ID uint `form:"id" gorm:"primaryKey;autoIncrement"`
	Qty uint `form:"qty" gorm:"not null"`
	ProductID uint `form:"product_id" gorm:"not null"`	
	Product Product
	OrderID uint `form:"order_id" gorm:"not null"`

	
}

func DatabaseInit() {

	var err error

	dsn := "host=localhost user=postgres password=toor dbname=go_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannot Connect to DB")
	}

	AutoMigration(DB)

	fmt.Println("Connected to DB")
}

func AutoMigration (db *gorm.DB) {
	db.Debug().AutoMigrate(
		&User{},
		&UserToken{},
		&Product{},
		&Cart{},
		&CartItem{},
		&Order{},
		&OrderItem{},
	)
}
