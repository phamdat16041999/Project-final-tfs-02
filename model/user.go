package model

import (
	"encoding/json"
	"fmt"
	"hotel/auth"
	"hotel/connect"
	"hotel/gmail"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// var (
// 	router = gin.Default()
// )

type User struct {
	gorm.Model
	FirstName          string           `gorm:"type:varchar(100);" json:"firstName"`
	LastName           string           `gorm:"type:varchar(100);" json:"lastName"`
	DOB                time.Time        `json:"dob"`
	Address            string           `gorm:"type:varchar(100);" json:"address"`
	Phone              int              `json:"phone"`
	Email              string           `gorm:"type:varchar(100);unique;" json:"email"`
	CodeAuthentication string           `gorm:"type:varchar(20);unique;" json:"codeAuthentication"`
	UserName           string           `gorm:"type:varchar(100);unique;" json:"userName"`
	Password           string           `gorm:"type:varchar(100); default: 123;" json:"password"`
	Active             *bool            `gorm:"default: false;" json:"active"`
	RoleID             uint             `gorm:"default: 0;" json:"roleid"`
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
	var User = User{FirstName: user.FirstName, LastName: user.LastName, DOB: user.DOB, Address: user.Address, Phone: user.Phone, Email: user.Email, CodeAuthentication: randomCode, UserName: user.UserName, Password: hash, Active: user.Active, RoleID: user.RoleID}
	result := db.Create(&User)
	if result.Error != nil {
		fmt.Fprint(w, result.Error)
		return
	}
	fmt.Fprint(w, "Successfully")
}

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
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// func ActiveAccount(w http.ResponseWriter, r *http.Request) {
// 	var user User
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	db := connect.Connect()
// 	var query User
// 	db.Where("active = ?", user.Active).Find(&query)
// 	b, _ := json.Marshal(query.Active)
// 	fmt.Fprintf(w, string(b))
// 	return

// }

// func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
// 	hmacSecretString := // Value
// 	hmacSecret := []byte(hmacSecretString)
// 	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
// 			// check token signing method etc
// 			return hmacSecret, nil
// 	})

// 	if err != nil {
// 		return nil, false
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		return claims, true
// 	} else {
// 		log.Printf("Invalid JWT Token")
// 		return nil, false
// 	}
// }

func LoginAcount(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := connect.Connect()
	var query User
	db.Where("user_name = ?", user.UserName).Find(&query)
	db.Where("role_id = ?", user.RoleID).Find(&query)
	b, _ := json.Marshal(query.UserName)
	b1, _ := json.Marshal(query.RoleID)
	username := string(b)
	role1 := string(b1)
	fmt.Fprint(w, username)
	fmt.Fprint(w, role1)
	if len(username) == 2 {
		fmt.Fprint(w, "Username not created yet!")
	} else {
		b, _ := json.Marshal(query.Password)
		password := strings.Split(string(b), "\"")
		bb, _ := json.Marshal(query.Active)
		x := string(bb)

		if CheckPasswordHash(user.Password, password[1]) {
			b, _ := json.Marshal(query.ID)
			b1, _ := json.Marshal(query.ID)
			userid, _ := strconv.ParseUint(string(b), 10, 64)
			roleid, _ := strconv.ParseUint(string(b1), 10, 64)
			if x == "true" {
				token, err := auth.CreateToken(userid, roleid)
				if err != nil {
					fmt.Fprint(w, err.Error())
					return
				}
				auth.JSON(w, http.StatusOK, token)
			} else {
				fmt.Fprint(w, "Not Active!")
			}
		} else {
			fmt.Fprint(w, "Wrong Password!")
		}
	}
}
func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := connect.Connect()
	var query User
	db.Where("email = ?", user.Email).Find(&query)
	b, _ := json.Marshal(query.Email)
	email := strings.Split(string(b), "\"")
	if email[1] == "" {
		fmt.Fprint(w, "Email does not exist")
	} else {
		fmt.Fprint(w, "Trả về đường dẫn để nhập mã code và gửi code vào email")
		RandomCode := codeAuthentication()
		gmail.SendEmail(user.Email, RandomCode)
	}
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := connect.Connect()

	hash, _ := HashPassword(user.Password)

	var query User
	db.First(&query, user.ID)
	b, _ := json.Marshal(query.CodeAuthentication)
	code := strings.Split(string(b), "\"")
	// b1, _ := json.Marshal(query.Email)
	// email := strings.Split(string(b1), "\"")
	// randomCode := codeAuthentication()
	// gmail.SendEmail(user.Email, randomCode) // gửi mã code về mail người dùng
	// db.Model(&query).Where("ID = ?", user.ID).Update("code_authentication", randomCode)

	if code[1] == user.CodeAuthentication { //Mã code người dùng nhập đúng thì update mật khẩu
		result := db.Model(&query).Where("ID = ?", user.ID).Update("password", hash)
		if result.Error != nil {
			fmt.Fprint(w, "Error change")
		} else {
			fmt.Fprint(w, "Change Password Successfully")
		}
	} else {
		fmt.Fprint(w, "Wrong Code Authentication")
	}
}

// active account
func ActiveAccount(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := connect.Connect()
	var query User
	db.First(&query, user.ID)
	b, _ := json.Marshal(query.CodeAuthentication)
	code := strings.Split(string(b), "\"")
	if code[1] == user.CodeAuthentication {
		result := db.Model(&query).Where("ID = ?", user.ID).Update("active", true)
		if result.Error != nil {
			fmt.Fprint(w, "Account activation failed")
		} else {
			fmt.Fprint(w, "Successfully")
		}

	} else {
		fmt.Fprint(w, "Wrong code")
	}
}
