package model

import (
	"encoding/json"
	// "errors"
	"fmt"
	"hotel/auth"
	"hotel/connect"
	"hotel/gmail"
	"hotel/middlewares"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	// "github.com/badoux/checkmail"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// var (
// 	router = gin.Default()
// )

type User struct {
	gorm.Model
	FirstName          string            `gorm:"type:varchar(100);" json:"firstName,omitempty"`
	LastName           string            `gorm:"type:varchar(100);" json:"lastName,omitempty"`
	Address            string            `gorm:"type:varchar(100);" json:"address,omitempty"`
	DOB                string            `json:"dob,omitempty"`
	Phone              string            `json:"phone,omitempty"`
	Email              string            `gorm:"type:varchar(100);unique;" json:"email,omitempty"`
	CodeAuthentication string            `gorm:"type:varchar(20);unique;" json:"codeAuthentication,omitempty"`
	UserName           string            `gorm:"type:varchar(100);unique;" json:"userName,omitempty"`
	Password           string            `gorm:"type:varchar(100); default: 123;" json:"password,omitempty"`
	Active             *bool             `gorm:"default: false;" json:"active,omitempty"`
	Authentication     []*Authentication `json:"authentication,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:UserID;associationForeignKey:ID"`
	Conversation1      []Conversation    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:User1ID; associationForeignKey:ID"`
	Conversation2      []Conversation    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:User2ID; associationForeignKey:ID"`
	Messenger          []Messenger       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:UserID;associationForeignKey:ID"`
	Hotel              []Hotel           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:UserID;associationForeignKey:ID"`
	Rate               []Rate            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:UserID;associationForeignKey:ID"`
	Bill               []Bill            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:UserID;associationForeignKey:ID"`
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
	errmail := gmail.SendEmail(user.Email, randomCode)
	var User = User{
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		DOB:                user.DOB,
		Address:            user.Address,
		Phone:              user.Phone,
		Email:              user.Email,
		CodeAuthentication: randomCode,
		UserName:           user.UserName,
		Password:           hash,
		Active:             user.Active,
	}
	if errmail != "" {
		fmt.Fprintln(w, "Email is incorrect!")
		return
	} else {
		result := db.Create(&User)
		if result.Error != nil {
			fmt.Fprintln(w, "Account already in use, please change username and email !")
			return
		} else {
			var newUser = User
			type Response struct {
				ID        int
				Messenger string
			}
			db.Last(&newUser)
			var response Response
			response.ID = int(newUser.ID)
			response.Messenger = "Create successfull"
			w.WriteHeader(http.StatusOK)
		}
		var Auth = Authentication{
			UserID: User.ID,
			RoleID: 1,
		}
		auth := db.Create(&Auth)
		if auth.Error != nil {
			fmt.Fprintln(w, "Error!")
			return
		}
	}

}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data")
	UserID := middlewares.ConvertDataToken(data, "user_id")
	userid, err1 := strconv.ParseUint(UserID, 10, 64)
	if err1 != nil {
		fmt.Println("error:", err1)
	}
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := connect.Connect()
	hash, _ := HashPassword(user.Password)
	randomCode := codeAuthentication()
	var query User
	var query1 User
	db.Where("id = ?", userid).Find(&query1)
	b, _ := json.Marshal(query1.ID)
	if string(b) != "0" {
		db.Model(query).Where("id = ?", userid).Updates(User{
			FirstName:          user.FirstName,
			LastName:           user.LastName,
			DOB:                user.DOB,
			Address:            user.Address,
			Phone:              user.Phone,
			Email:              user.Email,
			CodeAuthentication: randomCode,
			UserName:           user.UserName,
			Password:           hash,
			Active:             user.Active,
		})
		fmt.Fprintln(w, "Update successfull!")
	} else {
		fmt.Fprintln(w, "Can not find ID")
	}

}
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data")
	UserID := middlewares.ConvertDataToken(data, "user_id")
	userid, err1 := strconv.ParseUint(UserID, 10, 64)
	if err1 != nil {
		fmt.Println("error:", err1)
	}
	db := connect.Connect()
	var query User
	db.Where("id = ?", userid).Find(&query)
	b, _ := json.Marshal(query.ID)
	if string(b) != "0" {
		db.Where("id = ?", userid).Delete(&query)
		fmt.Fprintln(w, "Delete successfull!")
	} else {
		fmt.Fprintln(w, "Can not find ID")
	}
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

// H??m ????? ki???m tra password ???????c m?? h??a c?? tr??ng v???i Password nh???p v??o hay kh??ng
// Password ?????u v??o l?? m???t kh???u ng?????i d??ng nh???p v??o
// Hash l?? m???t kh???u l???y t??? database
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

func LoginAcount(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := connect.Connect()
	var query User
	var queryAuth Authentication
	var queryRole Role

	db.Where("user_name = ?", user.UserName).Find(&query)
	db.Where("user_id = ?", query.ID).Find(&queryAuth)
	db.Where("id = ?", queryAuth.RoleID).Find(&queryRole)

	b, _ := json.Marshal(query.UserName)
	// b1, _ := json.Marshal(query.ID)
	// b2, _ := json.Marshal(queryAuth.RoleID)
	b3, _ := json.Marshal(queryRole.Name)
	userName := string(b)
	// userId := string(b1)
	// roleId := string(b2)
	roleName := string(b3)

	// fmt.Fprintln(w, userId)
	// fmt.Fprintln(w, roleId)
	fmt.Fprintln(w, roleName)

	if len(userName) == 2 {
		fmt.Fprint(w, "Username not created yet!")
	} else {
		b, _ := json.Marshal(query.Password)
		password := strings.Split(string(b), "\"")
		b1, _ := json.Marshal(query.Active)
		x := string(b1)

		if CheckPasswordHash(user.Password, password[1]) {
			b, _ := json.Marshal(query.ID)
			b1, _ := json.Marshal(queryAuth.RoleID)
			userId, _ := strconv.ParseUint(string(b), 10, 64)
			roleId, _ := strconv.ParseUint(string(b1), 10, 64)
			if x == "true" {
				token, err := auth.CreateToken(userId, roleId)
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
		fmt.Fprint(w, "Tr??? v??? ???????ng d???n ????? nh???p m?? code v?? g???i code v??o email")
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
	// gmail.SendEmail(user.Email, randomCode) // g???i m?? code v??? mail ng?????i d??ng
	// db.Model(&query).Where("ID = ?", user.ID).Update("code_authentication", randomCode)

	if code[1] == user.CodeAuthentication { //M?? code ng?????i d??ng nh???p ????ng th?? update m???t kh???u
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
	db.Debug().Where("email=?", user.Email).First(&query)
	b, _ := json.Marshal(query.CodeAuthentication)
	code := strings.Split(string(b), "\"")
	if code[1] == user.CodeAuthentication {
		result := db.Model(&query).Where("email = ?", user.Email).Update("active", true)
		if result.Error != nil {
			fmt.Fprint(w, "Account activation failed")
		} else {
			fmt.Fprint(w, "Successfully")
		}

	} else {
		fmt.Fprint(w, "Wrong code")

		w.WriteHeader(http.StatusInternalServerError)

	}
}
