package model

import (
	"encoding/json"
	"fmt"
	"hotel/connect"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	Name        string       `gorm:"type:varchar(100);" json:"name" `
	Address     string       `gorm:"type:varchar(100);" json:"address" `
	Description string       `gorm:"type:varchar(100);" json:"description" `
	Image       string       `gorm:"type:varchar(100);" json:"image" `
	Longitude   string       `gorm:"type:varchar(100);" json:"longitude" `
	Latitude    string       `gorm:"type:varchar(100);" json:"latitude" `
	UserID      uint         `json:"userID"`
	AverageRate float64      `gorm:"default:0.0;" json:"averagerate"`
	NumberRate  float64      `gorm:"default:0;" json:"numberrate"`
	ImageHotel  []ImageHotel `json:"authentication omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
	Room        []Room       `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
	Rate        []Rate       `json:"authentication omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
	Bill        []Bill       `json:"authentication omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
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
}
type RoomInformation struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Img        []ImageRoom
	Time       []Times
	PriceHrs   int `json:"priceHrs"`
	PriceDay   int `json:"priceDay"`
	ExtraPrice int `json:"extraPrice"`
}

type HotelRate struct {
	HotelID uint `json:"hotelID"`
	UserID  uint `json:"userID"`
	Rate    int  `json:"rate"`
}

func DataHomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := connect.Connect()
	var results []Hotel
	db.Model(&Hotel{}).Distinct().Pluck("address", &results)
	b, err := json.Marshal(results)
	if err != nil {
		fmt.Fprintln(w, "Error:", err)
	}
	fmt.Fprintln(w, string(b))

}
func SeachHotelAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

func GetTopHotel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := connect.Connect()
	var hotel []Hotel
	db.Limit(2).Order("average_rate desc").Find(&hotel)
	b, _ := json.Marshal(hotel)
	fmt.Fprintln(w, string(b))
}

func GetDetailHotel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
		ID:          int(hotel.ID),
		Name:        hotel.Name,
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
		var time []Times
		db.Where("room_id = ?", rooms[i].ID).Find(&imageRoom)
		db.Where("name = ?", "PriceHours").Find(&option1)
		db.Where("room_id = ? AND option_id = ?", rooms[i].ID, option1.ID).Find(&priceHrs)
		db.Where("name = ?", "PriceDays").Find(&option2)
		db.Where("room_id = ? AND option_id = ?", rooms[i].ID, option2.ID).Find(&priceDay)
		db.Where("name = ?", "ExtraPrice").Find(&option3)
		db.Where("room_id = ? AND option_id = ?", rooms[i].ID, option3.ID).Find(&extraPrice)
		db.Debug().Where("room_id = ? AND active = ?", rooms[i].ID, true).Find(&time)

		roomInformation := RoomInformation{
			ID:         int(rooms[i].ID),
			Name:       rooms[i].Name,
			Img:        imageRoom,
			PriceHrs:   priceHrs.Price,
			PriceDay:   priceDay.Price,
			ExtraPrice: extraPrice.Price,
			Time:       time,
		}
		hotelInformation.Room = append(hotelInformation.Room, roomInformation)
	}
	b1, _ := json.Marshal(hotelInformation)
	fmt.Fprintln(w, string(b1))
}

func Rating(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var hotelrate HotelRate
	var rate Rate
	var hotel Hotel
	err := json.NewDecoder(r.Body).Decode(&hotelrate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := connect.Connect()
	db.Where("user_id = ? AND hotel_id = ?", hotelrate.UserID, hotelrate.HotelID).Find(&rate)
	b1, _ := json.Marshal(&rate.HotelID)
	ID, _ := strconv.ParseUint(string(b1), 10, 32)
	b, _ := json.Marshal(&rate.Rate)
	var Rate = Rate{
		UserID:  hotelrate.UserID,
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
		result := db.Model(&rate).Where("user_id = ? AND hotel_id = ?", hotelrate.UserID, hotelrate.HotelID).Update("rate", hotelrate.Rate)
		db.Where("id = ?", ID).Find(&hotel)
		b, _ := json.Marshal(&hotel.NumberRate)
		NumberRate, _ := strconv.ParseFloat(string(b), 64)
		b1, _ := json.Marshal(&hotel.AverageRate)
		b2, _ := json.Marshal(&hotelrate.Rate)
		NewRate, _ := strconv.ParseUint(string(b2), 10, 64)
		OldRate, _ := strconv.ParseUint(string(b3), 10, 64)
		fmt.Fprintln(w, NewRate, OldRate)
		AverageRate, _ := strconv.ParseFloat(string(b1), 64)
		AverageRate = (AverageRate*NumberRate - float64(OldRate) + float64(NewRate)) / NumberRate
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
func ConvertsTime(time string) []int {
	a := strings.Split(time, "T")
	x := strings.Split(a[1], ":")
	var s []int
	for i := 0; i < len(x); i++ {
		x, _ := strconv.Atoi(x[i])
		s = append(s, x)
	}
	return s
}
func ConvertsDate(time string) []int {
	a := strings.Split(time, "T")
	x := strings.Split(a[0], "-")
	var s []int
	for i := 0; i < len(x); i++ {
		x, _ := strconv.Atoi(x[i])
		s = append(s, x)
	}
	return s
}
func Checkroomstatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var checkTime Times
	err := json.NewDecoder(r.Body).Decode(&checkTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// b3, _ := json.Marshal(a)
	// fmt.Fprintln(w, string(b3))
	db := connect.Connect()
	var times []Times
	result := db.Where("room_id = ?", checkTime.RoomID).Find(&times)
	if result.Error != nil {
		fmt.Fprintln(w, "Can not find room: ", result.Error)
		return
	}
	available := 1
	for i := 0; i < len(times); i++ {
		b, _ := json.Marshal(&checkTime.StartTime)
		checkStartTime := ConvertsTime(string(b))
		checkStartDate := ConvertsDate(string(b))
		b3, _ := json.Marshal(&checkTime.EndTime)
		checkEndTime := ConvertsTime(string(b3))
		checkEndDate := ConvertsDate(string(b3))
		b1, _ := json.Marshal(&times[i].StartTime)
		start := ConvertsTime(string(b1))
		startDate := ConvertsDate(string(b1))
		b2, _ := json.Marshal(&times[i].EndTime)
		end := ConvertsTime(string(b2))
		endDate := ConvertsDate(string(b2))
		fmt.Fprintln(w, "check start: ", checkStartTime[0], checkStartDate[2], "check end: ", checkEndTime[0], checkEndDate[1], "start: ", start[0], startDate[1], "end: ", end[0], endDate[2])
		if checkStartDate[1] == startDate[1] {
			if checkStartDate[2] == endDate[2] {
				if checkStartTime[0] > start[0] && checkStartTime[0] < end[0] || checkEndTime[0] > start[0] && checkEndTime[0] < end[0] {
					available--
					break
				} else if checkStartTime[0] == start[0] && checkEndTime[0] == end[0] {
					if checkStartTime[1] > start[1] || checkEndTime[1] < end[1] {
						available--
						break
					}
				} else if checkStartTime[0] == start[0] || checkEndTime[0] == end[0] {
					if checkStartTime[0] == start[0] {
						if checkStartTime[1] >= start[1] {
							available--
							break
						}
					} else if checkEndTime[0] == end[0] {
						if checkEndTime[1] <= end[1] {
							available--
							break
						}
					}
				} else if checkStartDate[2] > startDate[2] && checkStartDate[2] < endDate[2] || checkEndDate[2] > startDate[2] && checkEndDate[2] < endDate[2] {
					available--
					break
				}
			} else if checkStartDate[1] > startDate[1] && checkStartDate[1] < endDate[1] || checkEndDate[1] > startDate[1] && checkEndDate[1] < endDate[1] {
				available--
				break
			}
		}
	}
	fmt.Fprintln(w, available)
	if available != 1 {
		fmt.Fprintln(w, "Room has been booked")
	} else {
		fmt.Fprintln(w, "Room avaliable")
	}
}
func CreateHotel(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	w.Header().Set("Content-Type", "application/json")
	userid := r.Context().Value("user_id")
	str := fmt.Sprintf("%v", userid)
	userID, _ := strconv.ParseUint(str, 10, 64)

	var Data Hotel
	err := json.NewDecoder(r.Body).Decode(&Data)
	if err != nil {
		fmt.Fprintln(w, err)
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
			fmt.Fprint(w, result.Error)
			continue
		}
		// create each image for each room
		for j := 0; j < len(Data.Room[i].ImageRoom); j++ {
			imagerRoom := ImageRoom{
				Image:  Data.Room[i].ImageRoom[j].Image,
				RoomID: room.ID,
			}
			result := db.Create(&imagerRoom)
			if result.Error != nil {
				fmt.Fprint(w, result.Error)
				continue
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
				fmt.Fprint(w, result.Error)
				continue
			}

		}

	}
}
func CreateHotel(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	w.Header().Set("Content-Type", "application/json")
	userid := r.Context().Value("user_id")
	str := fmt.Sprintf("%v", userid)
	userID, _ := strconv.ParseUint(str, 10, 64)

	var Data Hotel
	err := json.NewDecoder(r.Body).Decode(&Data)
	if err != nil {
		fmt.Fprintln(w, err)
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
			fmt.Fprint(w, result.Error)
			continue
		}
		// create each image for each room
		for j := 0; j < len(Data.Room[i].ImageRoom); j++ {
			imagerRoom := ImageRoom{
				Image:  Data.Room[i].ImageRoom[j].Image,
				RoomID: room.ID,
			}
			result := db.Create(&imagerRoom)
			if result.Error != nil {
				fmt.Fprint(w, result.Error)
				continue
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
				fmt.Fprint(w, result.Error)
				continue
			}

		}

	}
}
