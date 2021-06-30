package model

import (
	"encoding/json"
	"fmt"
	"hotel/connect"
	"hotel/gmail"
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName          string           `gorm:"type:varchar(100); json:"firstName"`
	LastName           string           `gorm:"type:varchar(100); json:"lastName"`
	DOB                time.Time        `json:"dob"`
	Address            string           `gorm:"type:varchar(100); json:"address"`
	Phone              int              `json:"phone"`
	Email              string           `gorm:"type:varchar(100); json:"email"`
	CodeAuthentication string           `gorm:"type:varchar(20); json:"codeAuthentication"`
	UserName           string           `gorm:"type:varchar(100); json:"userName"`
	Password           string           `gorm:"type:varchar(100); json:"password"`
	Active             *bool            `json:"active"`
	Authentication     []Authentication `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:UserID;associationForeignKey:ID"`
	Conversation1      []Conversation   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:User1ID; associationForeignKey:ID"`
	Conversation2      []Conversation   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:User2ID; associationForeignKey:ID"`
	Messenger          []Messenger      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:UserID;associationForeignKey:ID"`
	Hotel              []Hotel          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:UserID;associationForeignKey:ID"`
	Rate               []Rate           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:UserID;associationForeignKey:ID"`
	Bill               []Bill           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:UserID;associationForeignKey:ID"`
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := connect.Connect()
	// Encrypt password
	hash, _ := HashPassword(user.Password)
	// Create password
	randomCode := codeAuthentication()
	gmail.SendEmail(user.Email, randomCode)
	var User = User{FirstName: user.FirstName, LastName: user.LastName, DOB: user.DOB, Address: user.Address, Phone: user.Phone, Email: user.Email, CodeAuthentication: randomCode, UserName: user.UserName, Password: hash, Active: user.Active}
	result := db.Create(&User)
	if result.Error != nil {
		fmt.Fprint(w, result.Error)
		return
	}
	fmt.Fprint(w, "Successfully")
}

// func ActiveAccount(w http.ResponseWriter, r *http.Request){

// }
// Random code
func codeAuthentication() string {
	rand.Seed(time.Now().UTC().UnixNano())
	return randomString(12)
}
func randomString(len int) string {

	bytes := make([]byte, len)

	for i := 0; i < len; i++ {
		bytes[i] = byte(randInt(97, 122))
	}

	return string(bytes)
}

func randInt(min int, max int) int {

	return min + rand.Intn(max-min)
}

// Random code
// Encrypt password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Hàm để kiểm tra password được mã hóa có trùng với Password nhập vào hay không
// Password đầu vào là mật khẩu người dùng nhập vào
// Hash là mật khẩu lấy từ database
// match := CheckPasswordHash(password, hash)
// fmt.Println("Match:   ", match)
// func CheckPasswordHash(password, hash string) bool {
//     err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//     return err == nil
// }
