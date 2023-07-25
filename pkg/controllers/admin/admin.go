package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"project/initializers"
	"project/pkg/models"
	"project/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!user manangement!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func ListUser(g *gin.Context) {
	var user []models.Userdata
	initializers.DB.Find(&user)
	response.SurcessMessage(g, "displaying users", user)
}

func BlockUser(g *gin.Context) {

	params := g.Param("id")
	fmt.Println(params)
	page, err := strconv.Atoi(params)
	if err != nil {
		response.ErrorMessage(g, "", err.Error(), http.StatusBadRequest)
	}
	var users models.Userdata
	initializers.DB.Raw("update Userdata SET Status=false WHERE id=?", page).Scan(&users)
	g.JSON(http.StatusOK, gin.H{"": "user boceked surcessfully", "id": page})
}
func Unblock(g *gin.Context) {
	params := g.Param("id")
	fmt.Printf("%T", params)
	page, _ := strconv.Atoi(params)
	var users models.Userdata
	initializers.DB.Raw("update Userdata SET Status=true WHERE id=?", page).Scan(&users)
	response.SurcessMessage(g, "user unblocked surcessfully", users)
}

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!Product manangement!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

func AddProduct(g *gin.Context) {
	var productimages []string
	if err := g.Request.ParseMultipartForm(32 << 20); err != nil {
		response.ErrorMessage(g, "error while parsing", err.Error(), 400)
	}

	file := g.Request.MultipartForm.File["image"]
	for _, img := range file {
		exrention := filepath.Ext(img.Filename)
		image := "product" + uuid.New().String() + exrention
		productimages = append(productimages, image)
		g.SaveUploadedFile(img, "./public/"+image)
	}

	fmt.Println(productimages)

	// data storage of user from body
	var product models.Product
	var homepic string
	body := g.PostForm("product")
	err := json.Unmarshal([]byte(body), &product)
	homeimg, err1 := g.FormFile("homeimage")
	if err1 == nil {
		exrention := filepath.Ext(homeimg.Filename)
		image := "product" + uuid.New().String() + exrention
		homepic = image
		g.SaveUploadedFile(homeimg, "./public/"+image)
	}

	if err != nil {
		response.ErrorMessage(g, "error on unmarshalling body", err.Error(), 400)
	}
	product.Image = homepic
	product_result := initializers.DB.Create(&product)

	if product_result.Error != nil {
		response.ErrorMessage(g, "failed to add product... try again", product_result.Error.Error(), 502)
		return
	}
	// respond
	response.SurcessMessage(g, "product added surcessfully", product)
}

func EditProduct(g *gin.Context) {
	var PP struct {
		ProductName string
		Category    int
		Brand       string
		Price       int
		Quantity    int
		Image       string
		Description string
	}

	params := g.Query("id")
	page, _ := strconv.Atoi(params)
	var product models.Product

	if err := g.Bind(&PP); err != nil {
		response.ErrorMessage(g, "failed to bind", err.Error(), http.StatusBadRequest)
		return
	}

	initializers.DB.First(&product, page)
	if product.ID < 1 {
		response.ErrorMessage(g, "failed to get product", " product is not on database", 502)
		return
	}

	//set changed total
	offer := product.ProductOffer + product.CategoryOffer
	totalprice := PP.Price - offer

	initializers.DB.Model(&product).Updates(models.Product{
		ProductName: PP.ProductName,
		Category:    PP.Category,
		Brand:       PP.Brand,
		Price:       PP.Price,
		Quantity:    PP.Quantity,
		Image:       PP.Image,
		Total:       totalprice,
		Description: PP.Description,
	})
	response.SurcessMessage(g, "product edited surcessfully", product)
}
func DeleteProduct(g *gin.Context) {
	params := g.Query("id")
	page, _ := strconv.Atoi(params)
	var product models.Product
	var res models.Product
	initializers.DB.Raw(" SELECT *FROM products WHERE id=?", page).Scan(&res)
	initializers.DB.Raw(" DELETE FROM products WHERE id=?", page).Scan(&product)

	response.SurcessMessage(g, "Product is deleted", res)

}

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!category manangement!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

func AddCategory(g *gin.Context) {

	var CC struct{ Categoryname string }

	if v := g.Bind(&CC); v != nil {
		response.ErrorMessage(g, "failed to bind", v.Error(), http.StatusBadRequest)
		return
	}

	AddCategory := models.Category{Categoryname: CC.Categoryname}
	category_result := initializers.DB.Create(&AddCategory)

	if category_result.Error != nil {
		response.ErrorMessage(g, "failed to add category... try again", category_result.Error.Error(), 502)
		return
	}
	// respond
	response.SurcessMessage(g, "category added surcessfully", AddCategory)
}

func EditCategory(g *gin.Context) {
	params := g.Query("id")
	fmt.Printf("%T", params)
	page, _ := strconv.Atoi(params)

	var CC struct {
		Categoryname string
	}
	if v := g.Bind(&CC); v != nil {
		response.ErrorMessage(g, "failed to bind", v.Error(), http.StatusBadRequest)
		return
	}
	// checking category
	var check models.Category
	initializers.DB.First(&check, "id= ?", page)
	if check.ID < 1 {
		response.ErrorMessage(g, "checking not found", "invalid category or category not found on database", 400)
		return
	}

	var EditCategory models.Category
	if CC.Categoryname != "" {
		initializers.DB.Raw("update categories SET Categoryname=? WHERE id=?", CC.Categoryname, page).Scan(&EditCategory)
		response.SurcessMessage(g, "category name edited surcessfully", EditCategory)
	}

}
func DeleteCategory(g *gin.Context) {
	params := g.Query("id")
	fmt.Printf("%T", params)
	page, _ := strconv.Atoi(params)
	var category models.Category
	initializers.DB.Raw(" DELETE FROM categories WHERE id=?", page).Scan(&category)
	response.SurcessMessage(g, "Category is deleted", category)
}
