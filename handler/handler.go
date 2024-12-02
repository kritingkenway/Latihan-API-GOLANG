package handler

import (
	"coba/model"
	"coba/utils"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!!!!")
}

func HelloName(c *fiber.Ctx) error {
	return c.SendString("Halo " + c.FormValue("name") + " dan umurmu adalah " + c.FormValue("umur"))
}


// bikin CRUD user API

func GetUsers(c *fiber.Ctx) error {

	var users []model.User

	model.DB.Find(&users)
	return c.JSON(fiber.Map{
		"users" : users,
	})
}

func GetUserById(c *fiber.Ctx) error{
	id := c.Params("id")

	var user model.User

	model.DB.Where("id = ?", id).First(&user)

	return c.JSON(fiber.Map{
		"user" : user,
	})
}

func CreateUser(c *fiber.Ctx) error {
	nama, email, password := c.FormValue("nama"), c.FormValue("email"), c.FormValue("password")
	umur, _ := strconv.Atoi(c.FormValue("umur"))

	// model.User{
	// 	Nama 	: nama,
	// 	Umur 	: umur,
	// 	Email 	: &email,
	// }

	model.DB.Create(&model.User{
		Nama 	: nama,
		Umur 	: umur,
		Email 	: &email,
		Password: password,
	})

	return c.JSON(fiber.Map{
		"message" : "data berhasil dibuat",
	})


}

func UpdateUser(c *fiber.Ctx) error {
	id ,_ := c.ParamsInt("id")
	nama, email := c.FormValue("nama"), c.FormValue("email")
	umur, _ := strconv.Atoi(c.FormValue("umur"))

	model.DB.Save(&model.User{
		ID		: id,
		Nama	: nama,
		Umur	: umur,
		Email	: &email ,
	})

	return c.JSON(fiber.Map{
		"message" : "Data berhasil di Update",
	})
}

func Login(c *fiber.Ctx) error {
	
	nama, password := c.FormValue("nama"), c.FormValue("password") 

	var user model.User

	

	if err := model.DB.Where("nama = ?", nama).First(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "email atau password salah",
		})
	}

	if user.Password != password {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "email atau password salah",
		})
	}

	var userToken model.UserToken

	if err := model.DB.Where("user_id = ?", user.ID).First(&userToken).Error; err == nil {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"message" : true,
			"token": userToken.Token,
		})
	}


	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status" : false,
			"error" : err.Error(),
		})
	}

	userToken = model.UserToken{
		UserID: user.ID,
		Token : token,
	}

	if err := model.DB.Create(&userToken).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status" : false,
			"message" : err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Token" :token,
		
	})
}

func Logout(c *fiber.Ctx) error {

	// fungsi log out dengan menghapus token user dari database
	userId := c.Locals("user").(model.User)
	
	var userToken model.UserToken


	if err := model.DB.Where("user_id =?", userId.ID).First(&userToken).Error; err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status" : false,
			"message" : err.Error(),
			
			
		})
	} else {

		model.DB.Delete(&userToken)
	
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": true,
			"message" : "User has been Logged Out",
			
		})
	}


}


// API PRODUCT

func GetProducts(c *fiber.Ctx) error {
	var product model.Product

	if err := model.DB.Find(&product).Error; err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"message" : "Tidak ada Konten Produk",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status" : true,
		"Product" : product,
	})
}


func GetProductById(c *fiber.Ctx) error {
	var product model.Product
	productId, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"message": err.Error(),
		})
	}

	if err := model.DB.Where("id =?",productId).First(&product).Error; err !=nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status" : false,
			"message" : err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status" : true,
		"Product": product,
	})
}

func CreateProduct(c *fiber.Ctx) error {

	product := new(model.Product)

	
	
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"message" : err.Error(),
		})
	} 
		
	
	

	model.DB.Create(&product)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status" : true,
		"message" : "produk telah dibuat",
	})
}

func UpdateProduct(c *fiber.Ctx) error {

	productId, err := c.ParamsInt("id")
	product := new(model.Product)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"message": err.Error(),
		})
	}


	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status" : false,
			"message" : err.Error(),
		})
	}

	updatedProduct := model.Product{
		ID : productId,
		Name : product.Name,
		Harga : product.Harga,
		Stock : product.Stock,
	}
	
	if err := model.DB.Where("id =?",productId).First(&product).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status" : false,
			"message" : err.Error(),
		})

	}

	model.DB.Save(&updatedProduct)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": true,
		"message": "Produk telah diperbarui",
	})
}
func DeleteProduct(c *fiber.Ctx) error {

	productId, err := c.ParamsInt("id")
	var product model.Product

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"message": err.Error(),
		})
	}

	
	
	if err := model.DB.Where("id =?",productId).First(&product).Delete(&product).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status" : false,
			"message" : err.Error(),
		})

	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": true,
		"message": "Produk telah dihapus",
	})
}



// Cart Api susah GK ngerti soal preload

func GetCartItems(c *fiber.Ctx) error {
	cartID := c.Locals("cartID")
	cart := model.Cart{}
	model.DB.Preload("CartItem").Preload("CartItem.Product").Where("id = ?", cartID).Find(&cart)

	fmt.Print(cartID)
	return c.JSON(fiber.Map{
		"result" : cart,
	})
}

func AddItemCart(c *fiber.Ctx) error {
	cartID := c.Locals("cartID").(uint)
	cartItem := model.CartItem{}
	productID, _ := c.ParamsInt("id")


	if err := model.DB.Where("cart_id = ?",cartID).Where("product_id = ?",productID).First(&cartItem).Error; err != nil {
		
		cartItem = model.CartItem{
			Qty: 1,
			ProductID: uint(productID),
			CartID: cartID,
		}
		
		model.DB.Create(&cartItem)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status" : true,
			"message" : "produk telah berhasil ditambahkan ke keranjang",
		})
	}
	
	

	// updateItems := model.CartItem{
	// 	ID: cartItem.ID,
	// 	Qty: cartItem.Qty + 1,
	// 	ProductID: uint(productID),
	// 	CartID: cartID,
	// }

	// model.DB.Save(&updateItems)

	model.DB.Model(&cartItem).Update("qty", cartItem.Qty + 1)
	
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status" : true,
		"message" : "Kuantitas Produk telah ditambahkan",
	})
}

func SubtractItemCart(c *fiber.Ctx) error {
	cartID := c.Locals("cartID").(uint)
	cartItem := model.CartItem{}
	productID, _ := c.ParamsInt("id")


	if err := model.DB.Where("cart_id = ?",cartID).Where("product_id = ?",productID).First(&cartItem).Error; err != nil {
		
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status" : false,
			"message" : err.Error(),
		})
	}
	
	if cartItem.Qty > 1 {


		// updateItems := model.CartItem{
		// 	ID: cartItem.ID,
		// 	Qty: cartItem.Qty - 1,
		// 	ProductID: uint(productID),
		// 	CartID: cartID,
		// }
	
		// model.DB.Save(&updateItems)
		
		model.DB.Model(&cartItem).Update("qty", cartItem.Qty - 1)

		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"status" : true,
			"message" : "Kuantitas Produk telah dikurangi",
		})
	}
	model.DB.Delete(&cartItem)
	
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status" : true,
		"message" : "Produk telah dihapus dari daftar keranjang",
	})

}