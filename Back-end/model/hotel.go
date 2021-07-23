package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hotel/connect"
	"hotel/middlewares"
	"hotel/pkg/cache"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	elastic "github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	Name        string       `gorm:"type:varchar(100);" json:"name,omitempty" `
	Address     string       `gorm:"type:varchar(100);" json:"address,omitempty" `
	Description string       `gorm:"type:varchar(100);" json:"description,omitempty" `
	Image       string       `gorm:"type:varchar(100);" json:"image,omitempty" `
	Longitude   string       `gorm:"type:varchar(100);" json:"longitude,omitempty" `
	Latitude    string       `gorm:"type:varchar(100);" json:"latitude,omitempty" `
	UserID      uint         `json:"userID,omitempty"`
	AverageRate float64      `gorm:"default:0.0;" json:"averagerate,omitempty"`
	NumberRate  float64      `gorm:"default:0;" json:"numberrate,omitempty"`
	ImageHotel  []ImageHotel `json:"imagehotel,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
	Room        []Room       `json:"room,omitempty" gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
	Rate        []Rate       `json:"rate,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
	Bill        []Bill       `json:"bill,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
}

type TopHotel struct {
	ID          uint64 `json:"id"`
	Name        string `gorm:"type:varchar(100);" json:"name" `
	Address     string `gorm:"type:varchar(100);" json:"address" `
	Description string `gorm:"type:varchar(100);" json:"description" `
	Image       string `gorm:"type:varchar(100);" json:"image" `
	Longitude   string `gorm:"type:varchar(100);" json:"longitude" `
	Latitude    string `gorm:"type:varchar(100);" json:"latitude" `
	UserID      uint   `json:"userID"`
	Rate        int    `json:"rate"`
}
type HotelInformation struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	Room        []RoomInformation `json:"room"`
	Longitude   string            `json:"longitude"`
	Latitude    string            `json:"latitude"`
	Description string            `json:"description"`
	Address     string            `json:"address"`
}
type RoomInformation struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Img        []ImageRoom
	Decription string `json:"decription"`
	PriceHrs   string `json:"priceHrs"`
	PriceDay   string `json:"priceDay"`
	ExtraPrice string `json:"extraPrice"`
}

type HotelRate struct {
	HotelID uint `json:"hotelID"`
	UserID  uint `json:"userID"`
	Rate    int  `json:"rate"`
}

func DataHomePage(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	var results []Hotel
	db.Model(&Hotel{}).Distinct().Pluck("address", &results)
	b, _ := json.Marshal(results)
	fmt.Fprintln(w, string(b))
}

func SearchByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := connect.Connect()
	var hotels []Hotel
	db.Debug().Raw("SELECT * FROM hotels WHERE MATCH (name,address,description) AGAINST (? IN NATURAL LANGUAGE MODE)", "Ninh Binh").Scan(&hotels)
	b, _ := json.Marshal(hotels)
	fmt.Fprintln(w, string(b))
}

type ESClient struct {
	*elastic.Client
}
type BookManager struct {
	esClient *ESClient
}

func NewHotelManager(es *ESClient) *BookManager {
	return &BookManager{esClient: es}
}
func (bm *BookManager) SearchHotels(name string) []*Hotel {
	ctx := context.Background()
	if bm.esClient == nil {
		fmt.Println("Nil es client")
		return nil
	}
	// build query to search for title
	query := elastic.NewSearchSource()
	query.Query(elastic.NewMatchQuery("name", name))
	// get search's service
	searchService := bm.esClient.
		Search().
		Index("hotels").
		SearchSource(query)
	// perform search query
	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("Cannot perform search with ES", err)
		return nil
	}
	// get result
	var hotels []*Hotel
	for _, hit := range searchResult.Hits.Hits {
		var hotel Hotel
		err := json.Unmarshal(hit.Source, &hotel)
		if err != nil {
			fmt.Println("Get data error: ", err)
			continue
		}
		hotels = append(hotels, &hotel)
	}
	return hotels
}
func EsSearchByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := string(vars["name"])
	url := "http://localhost:9200"
	esclient, _ := NewESClient(url)
	// search
	bm := NewHotelManager((*ESClient)(esclient))
	productGotten := bm.SearchHotels(name)
	JSONS(w, http.StatusOK, productGotten)
}

func NewESClient(url string) (*ESClient, error) {
	if len(url) == 0 {
		return nil, errors.New("empty url connection")
	}
	client, err := elastic.NewClient(elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return &ESClient{client}, err
}
func JSONS(w http.ResponseWriter, status int, object interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(object)
}
func GetHotelAddress(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	vars := mux.Vars(r)
	rate, _ := strconv.ParseFloat(vars["rate"], 64)
	var hotels []Hotel
	var resulthotels []Hotel
	db.Where("address LIKE ?", "%"+vars["address"]+"%").Find(&hotels)
	for i := 0; i < len(hotels); i++ {
		if int64(hotels[i].AverageRate) == int64(rate) {
			resulthotels = append(resulthotels, hotels[i])
		}
	}
	b, _ := json.Marshal(resulthotels)
	fmt.Fprintln(w, string(b))
}

func GetTopHotel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.ServeJQueryWithCache(w, "tophotel") == "No data in remote cache" {
		db := connect.Connect()
		var hotel []Hotel
		db.Limit(9).Order("average_rate desc").Find(&hotel)
		b, _ := json.Marshal(hotel)
		// cache.InsertData("tophotel", string(b))
		fmt.Fprintf(w, cache.InsertData("tophotel", string(b)))
	} else {
		fmt.Fprintln(w, cache.ServeJQueryWithCache(w, "tophotel"))
	}
}
func SearchHotelAddress(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	vars := mux.Vars(r)
	rate, _ := strconv.ParseFloat(vars["rate"], 64)
	var hotels []Hotel
	var resulthotels []Hotel
	db.Debug().Where("address LIKE ?", "%"+vars["address"]+"%").Find(&hotels)
	for i := 0; i < len(hotels); i++ {
		if int64(hotels[i].AverageRate) == int64(rate) {
			resulthotels = append(resulthotels, hotels[i])
		}
	}
	b, _ := json.Marshal(resulthotels)
	fmt.Fprintln(w, string(b))
}
func Hotelier(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	data := r.Context().Value("data")
	UserID := middlewares.ConvertDataToken(data, "user_id")
	userid, err1 := strconv.ParseUint(UserID, 10, 64)
	if err1 != nil {
		fmt.Println("error:", err1)
	}
	var hotel []Hotel
	db.Where("user_id = ?", userid).Find(&hotel)
	b1, _ := json.Marshal(&hotel)
	fmt.Fprintln(w, string(b1))

}
func GetDetailHotel(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var hotel Hotel
	var rooms []Room
	var imageRoom []ImageRoom

	db.Where("id = ?", id).Find(&hotel)
	// db.Where("hotel_id = ?", id).Find(&rooms)
	// for i := 0; i < len(rooms); i++ {
	// 	hotel.Room = append(hotel.Room, rooms[i])
	// 	var imageRooms []ImageRoom
	// 	db.Where("room_id = ?", rooms[i].ID).Find(&imageRooms)
	// 	for j := 0; j < len(imageRooms); j++ {
	// 		hotel.Room[i].ImageRoom = append(hotel.Room[i].ImageRoom, imageRooms[j])
	// 	}
	// 	var priceroom []Price
	// 	db.Where("room_id = ?", rooms[i].ID).Find(&priceroom)
	// 	for j := 0; j < len(priceroom); j++ {
	// 		hotel.Room[i].Price = append(hotel.Room[i].Price, priceroom[j])
	// 		// var option Option
	// 		// db.Where("id = ?", priceroom[i].OptionID).Find(&option)
	// 		// hotel.Room[i].Price[j].OptionID = append(hotel.Room[i].Price[j].OptionID, option)
	// 	}
	// }
	hotelInformation := HotelInformation{
		ID:      int(hotel.ID),
		Name:    hotel.Name,
		Address: hotel.Address,
		// room       : []RoomInformation
		Longitude:   hotel.Latitude,
		Latitude:    hotel.Longitude,
		Description: hotel.Description,
	}
	db.Where("hotel_id = ?", hotel.ID).Find(&rooms)
	for i := 0; i < len(rooms); i++ {
		var option1 Option
		var option2 Option
		var option3 Option
		var priceHrs Price
		var priceDay Price
		var extraPrice Price
		db.Where("room_id = ?", rooms[i].ID).Find(&imageRoom)
		db.Where("name = ?", "PriceHours").Find(&option1)
		db.Where("room_id = ? AND option_id = ?", rooms[i].ID, option1.ID).Find(&priceHrs)
		db.Where("name = ?", "PriceDays").Find(&option2)
		db.Where("room_id = ? AND option_id = ?", rooms[i].ID, option2.ID).Find(&priceDay)
		db.Where("name = ?", "ExtraPrice").Find(&option3)
		db.Where("room_id = ? AND option_id = ?", rooms[i].ID, option3.ID).Find(&extraPrice)

		roomInformation := RoomInformation{
			ID:         int(rooms[i].ID),
			Name:       rooms[i].Name,
			Img:        imageRoom,
			Decription: rooms[i].Description,
			PriceHrs:   priceHrs.Price,
			PriceDay:   priceDay.Price,
			ExtraPrice: extraPrice.Price,
		}
		hotelInformation.Room = append(hotelInformation.Room, roomInformation)
	}
	b1, _ := json.Marshal(hotelInformation)
	fmt.Fprintln(w, string(b1))
}

func Rating(w http.ResponseWriter, r *http.Request) {
	// cache.DeleteRemoteCache(w, "tophotel")
	// cache.DeleteLocalCache(w, "tophotel")
	data := r.Context().Value("data")
	UserID := middlewares.ConvertDataToken(data, "user_id")
	userid, err1 := strconv.ParseUint(UserID, 10, 64)
	if err1 != nil {
		fmt.Println("error:", err1)
	}
	//RoleID := middlewares.ConvertDataToken(data, "roles_id")
	//rolesid, _ := strconv.ParseUint(RoleID, 10, 64)
	var hotelrate HotelRate
	var rate Rate
	var hotel Hotel
	err := json.NewDecoder(r.Body).Decode(&hotelrate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := connect.Connect()
	db.Where("user_id = ? AND hotel_id = ?", userid, hotelrate.HotelID).Find(&rate)
	b1, _ := json.Marshal(&rate.HotelID)
	ID, _ := strconv.ParseUint(string(b1), 10, 32)
	b, _ := json.Marshal(&rate.Rate)
	var Rate = Rate{
		UserID:  uint(userid),
		HotelID: hotelrate.HotelID,
		Rate:    hotelrate.Rate,
	}
	if string(b) == "0" {
		result := db.Create(&Rate)
		if result.Error != nil {
			fmt.Fprintln(w, "Rating error: ", result.Error)
			return
		} else {
			db.Where("id = ?", hotelrate.HotelID).Find(&hotel)
			b, _ := json.Marshal(&hotel.NumberRate)
			NumberRate, _ := strconv.ParseFloat(string(b), 64)
			b1, _ := json.Marshal(&hotel.AverageRate)
			b3, _ := json.Marshal(&hotelrate.Rate)
			Rate, _ := strconv.ParseUint(string(b3), 10, 64)
			AverageRate, _ := strconv.ParseFloat(string(b1), 64)
			AverageRate = (AverageRate*NumberRate + float64(Rate)) / (NumberRate + 1)
			db.Model(Hotel{}).Where("id = ?", hotelrate.HotelID).Updates(Hotel{NumberRate: NumberRate + 1.0, AverageRate: AverageRate})
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "Rating successfull")
		}
	} else {
		b3, _ := json.Marshal(&rate.Rate)
		result := db.Model(&rate).Where("user_id = ? AND hotel_id = ?", userid, hotelrate.HotelID).Update("rate", hotelrate.Rate)
		db.Where("id = ?", ID).Find(&hotel)
		b, _ := json.Marshal(&hotel.NumberRate)
		NumberRate, _ := strconv.ParseFloat(string(b), 64)
		b1, _ := json.Marshal(&hotel.AverageRate)
		b2, _ := json.Marshal(&hotelrate.Rate)
		NewRate, _ := strconv.ParseUint(string(b2), 10, 64)
		OldRate, _ := strconv.ParseUint(string(b3), 10, 64)
		//fmt.Fprintln(w, "NewRate = ", NewRate, "\nOldRate = ", OldRate)
		AverageRate, _ := strconv.ParseFloat(string(b1), 64)
		AverageRate = (AverageRate*NumberRate - float64(OldRate) + float64(NewRate)) / NumberRate
		AverageRate = float64(int(AverageRate*10)) / 10 //chuyển thành số thập phân có 2 chữ số
		//fmt.Fprintln(w, "AverageRate = ", AverageRate)
		db.Model(Hotel{}).Where("id = ?", hotelrate.HotelID).Updates(Hotel{AverageRate: AverageRate})
		if result.Error != nil {
			fmt.Fprintln(w, "Update rating error: ", result.Error)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "Update rating successfull")
		}
	}
}
func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

// func Converts(time string) int {
// 	a := strings.Split(time, "T")
// 	x1 := strings.Split(a[1], ":")
// 	x2 := strings.Split(a[0], "-")
// 	var s1, s2 []int
// 	for i := 0; i < len(x1); i++ {
// 		x1, _ := strconv.Atoi(x1[i])
// 		s1 = append(s1, x1)
// 	}
// 	for j := 0; j < len(x2); j++ {
// 		x2, _ := strconv.Atoi(x2[j])
// 		s2 = append(s2, x2)
// 	}
// 	hashtime := (s1[0]-00)*3600 + (s1[1]-00)*60
// 	hashdate := (s2[1]-7)*30 + (s2[1]-01)*1
// 	hash := hashtime + hashdate
// 	return hash
// }

func Checkroomstatus(w http.ResponseWriter, r *http.Request) {

	var checkTime Times
	err := json.NewDecoder(r.Body).Decode(&checkTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := connect.Connect()
	var times []Times
	result := db.Where("room_id = ?", checkTime.RoomID).Find(&times)
	if result.Error != nil {
		fmt.Fprintln(w, "Can not find room: ", result.Error)
		return
	}
	available := 1
	for i := 0; i < len(times); i++ {
		if inTimeSpan(times[i].StartTime, times[i].EndTime, checkTime.StartTime) {
			available--
			break
		} else if inTimeSpan(times[i].StartTime, times[i].EndTime, checkTime.EndTime) {
			available--
			break
		}
	}
	if available != 1 {
		fmt.Fprintln(w, "Room has been booked")
	} else {
		fmt.Fprintln(w, "Room avaliable")
	}
}
func CreateHotel(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	data := r.Context().Value("data")
	DataUser := middlewares.ConvertDataToken(data, "user_id")
	userID, err1 := strconv.ParseUint(DataUser, 10, 64)
	if err1 != nil {
		fmt.Println("error:", err1)
	}
	var Data Hotel
	err := json.NewDecoder(r.Body).Decode(&Data)
	if err != nil {
		fmt.Println(w, err)
	}
	hotel := Hotel{
		Name:        Data.Name,
		Address:     Data.Address,
		Description: Data.Description,
		Image:       Data.Image,
		Longitude:   Data.Longitude,
		Latitude:    Data.Latitude,
		UserID:      uint(userID),
	}
	result := db.Create(&hotel)
	if result.Error != nil {
		fmt.Fprint(w, result.Error)
	}
	// create each room
	for i := 0; i < len(Data.Room); i++ {
		room := Room{
			Name:        Data.Room[i].Name,
			Description: Data.Room[i].Description,
			HotelID:     hotel.ID,
		}
		result := db.Create(&room)
		if result.Error != nil {
			fmt.Print(result.Error)
			return
		}
		// create each image for each room
		for j := 0; j < len(Data.Room[i].ImageRoom); j++ {
			imagerRoom := ImageRoom{
				Image:  Data.Room[i].ImageRoom[j].Image,
				RoomID: room.ID,
			}
			result := db.Create(&imagerRoom)
			if result.Error != nil {
				fmt.Print(result.Error)
				return
			}
		}
		for k := 0; k < len(Data.Room[i].Price); k++ {
			price := Price{
				Price:    Data.Room[i].Price[k].Price,
				OptionID: Data.Room[i].Price[k].OptionID,
				RoomID:   room.ID,
			}
			result := db.Create(&price)
			if result.Error != nil {
				fmt.Print(result.Error)
				return
			}

		}

	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Create successfull")
}
func UpdateHotel(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	datatoken := r.Context().Value("data")
	DataUser := middlewares.ConvertDataToken(datatoken, "user_id")
	userID, err1 := strconv.ParseUint(DataUser, 10, 64)
	if err1 != nil {
		fmt.Println("error:", err1)
	}
	var data HotelInformation
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("errr", err)
	}
	// fmt.Println(string(b))
	// take older rate
	var option []Option
	db.Find(&option)
	var hotelOdler Hotel
	var rooms []Room
	var imageroom []ImageRoom
	for _, room := range data.Room {
		for _, image := range room.Img {
			imageroom = append(imageroom, ImageRoom{
				Model: gorm.Model{
					ID: image.ID,
				},
				Image:  image.Image,
				RoomID: uint(room.ID),
			})
		}
		db.Debug().Model(&Price{}).Where("room_id = ? AND option_id = ?", room.ID, option[0].ID).Update("price", room.PriceDay)
		db.Debug().Model(&Price{}).Where("room_id = ? AND option_id = ?", room.ID, option[1].ID).Update("price", room.PriceHrs)
		db.Debug().Model(&Price{}).Where("room_id = ? AND option_id = ?", room.ID, option[2].ID).Update("price", room.ExtraPrice)
		rooms = append(rooms, Room{
			Model: gorm.Model{
				ID: uint(room.ID),
			},
			Name:        room.Name,
			Description: room.Decription,
			HotelID:     uint(room.ID),
			ImageRoom:   imageroom,
		})
	}
	db.Where("id =?", data.ID).Find(&hotelOdler)
	hotel := Hotel{
		Model: gorm.Model{
			ID: uint(data.ID),
		},
		Name:        data.Name,
		Address:     data.Address,
		Description: data.Description,
		Image:       "",
		Longitude:   data.Longitude,
		Latitude:    data.Latitude,
		UserID:      uint(userID),
		AverageRate: hotelOdler.AverageRate,
		NumberRate:  hotelOdler.NumberRate,
		Room:        rooms,
	}
	// b, _ := json.Marshal(&hotel)
	// fmt.Println(string(b))
	// fmt.Println(hotel)
	result := db.Save(&hotel)
	if result.Error != nil {
		fmt.Fprintln(w, "Rating error: ", result.Error)
	}
	result1 := db.Save(&rooms)
	if result1.Error != nil {
		fmt.Fprintln(w, "Rating error: ", result1.Error)
	}
	result2 := db.Save(&imageroom)
	if result2.Error != nil {
		fmt.Fprintln(w, "Rating error: ", result2.Error)
	}
}
func DeleteHotel(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	data := r.Context().Value("data")
	UserID := middlewares.ConvertDataToken(data, "user_id")
	userid, err1 := strconv.ParseUint(UserID, 10, 64)
	if err1 != nil {
		fmt.Println("error:", err1)
	}
	var hotel Hotel
	hotelid, _ := strconv.Atoi(mux.Vars(r)["id"])
	db.Where("id = ? AND user_id", hotelid, userid).Delete(&hotel)
	var hotelInformation []Hotel
	db.Debug().Where("user_id = ?", 6).Find(&hotelInformation)
	b1, _ := json.Marshal(&hotelInformation)
	fmt.Fprintln(w, string(b1))
}
