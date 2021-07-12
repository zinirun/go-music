package dblayer

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/zinirun/go-music/backend/src/models"
	"golang.org/x/crypto/bcrypt"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbname, con string) (*DBORM, error) {
	db, err := gorm.Open(dbname, con)
	return &DBORM{
		DB: db,
	}, err
}

func (db *DBORM) GetAllProducts() (products []models.Product, err error) {
	return products, db.Find(&products).Error
}

func (db *DBORM) GetPromos() (products []models.Product, err error) {
	return products, db.Where("promotion IS NOT NULL").Find(&products).Error
}

func (db *DBORM) GetCustomerByName(firstName string, lastName string) (customer models.Customer, err error) {
	return customer, db.Where(&models.Customer{FirstName: firstName, LastName: lastName}).Find(&customer).Error
}

func (db *DBORM) GetCustomerByID(id int) (customer models.Customer, err error) {
	return customer, db.First(&customer, id).Error
}

func (db *DBORM) GetProduct(id int) (product models.Product, err error) {
	return product, db.First(&product, id).Error
}

func (db *DBORM) AddUser(customer models.Customer) (models.Customer, error) {
	hashPassword(&customer.Password)
	customer.LoggedIn = true
	return customer, db.Create(&customer).Error
}

func (db *DBORM) SignInUser(email, password string) (customer models.Customer, err error) {
	if !checkPassword(password) {
		return customer, errors.New("Invalid Password")
	}
	// Specify table to run operations
	result := db.Table("Customers").Where(&models.Customer{Email: email})
	// run operation (update)
	if err = result.Update("loggedin", 1).Error; err != nil {
		return customer, err
	}

	return customer, result.Find(&customer).Error
}

func (db *DBORM) SignOutUserByID(id int) error {
	customer := models.Customer{Model: gorm.Model{
		ID: uint(id)}}
	return db.Table("Customers").Where(&customer).Update("loggedin", 0).Error
}

func hashPassword(s *string) error {
	if s == nil {
		return errors.New("reference provided for hashing password is nil")
	}
	sBytes := []byte(*s)
	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	*s = string(hashedBytes[:])
	return nil
}

// func checkPassword(password string) bool {
// 	err := hashPassword(&password)
// 	err = bcrypt.CompareHashAndPassword(password, password)
// }
