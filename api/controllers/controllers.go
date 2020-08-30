package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/DeVil2O/moviebookingsystem/api/database"
	"github.com/DeVil2O/moviebookingsystem/api/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/sony/sonyflake"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var admin models.Admin
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &admin)
	var res models.ResponseResult
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	collection, err := database.GetDBCollection()

	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	var result models.Admin
	err = collection.FindOne(context.TODO(), bson.D{{"AdminId", admin.AdminId}}).Decode(&result)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			hash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 5)

			if err != nil {
				res.Error = "Error While Hashing Password, Try Again"
				json.NewEncoder(w).Encode(res)
				return
			}
			admin.Password = string(hash)

			_, err = collection.InsertOne(context.TODO(), admin)
			if err != nil {
				res.Error = "Error While Creating User, Try Again"
				json.NewEncoder(w).Encode(res)
				return
			}
			res.Result = "Registration Successful"
			json.NewEncoder(w).Encode(res)
			return
		}

		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Result = "Username already Exists!!"
	json.NewEncoder(w).Encode(res)
	return
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var admin models.Admin
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &admin)
	if err != nil {
		log.Fatal(err)
	}

	collection, err := database.GetDBCollection()

	if err != nil {
		log.Fatal(err)
	}
	var result models.Admin
	var res models.ResponseResult

	err = collection.FindOne(context.TODO(), bson.D{{"adminid", admin.AdminId}}).Decode(&result)

	if err != nil {
		res.Error = "Invalid username"
		json.NewEncoder(w).Encode(res)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(admin.Password))

	if err != nil {
		res.Error = "Invalid password"
		json.NewEncoder(w).Encode(res)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"adminid": result.AdminId,
		"name":    result.Name,
	})

	tokenString, err := token.SignedString([]byte("secretKey"))

	if err != nil {
		res.Error = "Error is occuring while generating token,Try again"
		json.NewEncoder(w).Encode(res)
		return
	}

	result.Token = tokenString
	result.Password = ""
	json.NewEncoder(w).Encode(result)
	setSession(admin.AdminId, 30*time.Minute, w)

}

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func setSession(adminid string, ttl time.Duration, response http.ResponseWriter) {
	value := map[string]string{
		"adminid": adminid,
	}
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		panic(err)
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		fmt.Println(loc)
		cookie := &http.Cookie{
			Name:   "session",
			Value:  encoded,
			MaxAge: 30,
		}
		http.SetCookie(response, cookie)
		fmt.Println(cookie.Expires)
	}

}

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	adminid := params["adminid"]
	fmt.Println(adminid)

	err := CreateTickets(adminid, "Chirag Garg", "8218517963")
	var res models.ResponseResult
	if err != nil {
		res.Error = "HouseFull"
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Result = "Ticket Created Successfully"
	json.NewEncoder(w).Encode(res)
	return

}

func CreateTickets(adminId string, Customername string, Phonenumber string) error {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, _ := flake.NextID()
	ticket := &models.Ticket{id, Customername, Phonenumber, time.Now(), time.Now().Add(180 * time.Minute), time.Now(), false}
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("theatrebooking").C("admin")

	cursor := bson.M{"adminid": adminId}
	fmt.Println(cursor)

	change := bson.M{"$push": bson.M{"tickets": ticket}}
	if err != nil {
		panic(err)
	} else {
		collection, _ := database.GetDBCollection()
		var result models.Admin
		err = collection.FindOne(context.TODO(), bson.D{}).Decode(&result)
		if len(result.Tickets) <= 20 {
			erro := c.Update(cursor, change)
			if erro != nil {
				panic(erro)
			}

			fmt.Println("***********Ticket Booked*********")

		} else {
			fmt.Println("********HouseFull*********")
			err := 1
			if err == 1 {
				return fmt.Errorf("Housefull")
			}

		}
	}
	return err

}

func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	adminid := params["adminid"]
	ticketid := params["ticketid"]
	fmt.Println(adminid)
	u, _ := strconv.ParseUint(ticketid, 10, 64)
	UpdateTicketTimings(adminid, u, w)

}

func UpdateTicketTimings(adminId string, Ticketid uint64, w http.ResponseWriter) {
	collection, _ := database.GetDBCollection()
	var result models.Admin
	collection.FindOne(context.TODO(), bson.D{}).Decode(&result)
	for i, s := range result.Tickets {
		fmt.Printf("%T\n", s.TicketId)
		fmt.Println(i)
		if s.TicketId == Ticketid {
			fmt.Println(time.Now().Sub(s.StartTime))
			var res models.ResponseResult
			if time.Now().Sub(s.StartTime) >= 2400*time.Minute {
				err, _ := collection.ReplaceOne(context.TODO(), bson.M{"expired": false}, bson.M{"expired": true})
				if err != nil {

					fmt.Printf("%T\n", s.StartTime)
					res.Error = "Ticket is expired. Please Buy Another Ticket."
					json.NewEncoder(w).Encode(res)
					return
				}
			}

			s.StartTime = s.StartTime.Add(30 * time.Minute)
			fmt.Println(s.StartTime)
			result, err := collection.UpdateOne(context.TODO(), bson.M{"ticketid": Ticketid}, bson.D{{"$set", bson.D{{"starttime", s.StartTime}}}})

			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(result)

			res.Result = fmt.Sprintf("Timings Updated Successfully to Start Time %s , End Time %s ", s.StartTime.String(), s.StartTime.Add(180*time.Minute).String())

			json.NewEncoder(w).Encode(res)
			return

		}
	}
}

func GetTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	adminid := params["adminid"]
	timings := params["timings"]
	fmt.Println(adminid)
	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, timings)
	if err != nil {
		fmt.Println(err)
	}
	GetTickets(adminid, t, w)
}

func GetTickets(adminId string, timings time.Time, w http.ResponseWriter) []models.Ticket {
	collection, _ := database.GetDBCollection()
	var result models.Admin
	collection.FindOne(context.TODO(), bson.D{}).Decode(&result)

	var res []models.Ticket
	for i, s := range result.Tickets {
		fmt.Println(i)
		if s.StartTime == timings {
			fmt.Println(s)
			res = append(res, *s)

		}
	}

	json.NewEncoder(w).Encode(res)
	return res

}

func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	adminid := params["adminid"]
	ticketid := params["ticketid"]
	fmt.Println(adminid)

	u, _ := strconv.ParseUint(ticketid, 10, 64)
	DeleteTickets(adminid, u, w)
}

func DeleteTickets(adminId string, Ticketid uint64, w http.ResponseWriter) {
	var res models.ResponseResult
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("theatrebooking").C("admin")

	cursor := bson.M{"adminid": adminId}
	fmt.Println(cursor)

	change := bson.M{"$pull": bson.M{"tickets": bson.M{"ticketid": Ticketid}}}
	erro := c.Update(cursor, change)
	if erro != nil {
		panic(erro)
	}

	res.Result = fmt.Sprintf("Ticket No. %d  Deleted Successfully", Ticketid)

	json.NewEncoder(w).Encode(res)
	return

}

func UserDetailsTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	adminid := params["adminid"]
	ticketid := params["ticketid"]
	fmt.Println(adminid)

	u, _ := strconv.ParseUint(ticketid, 10, 64)
	UserDetailsTickets(adminid, u, w)
}

func UserDetailsTickets(adminId string, Ticketid uint64, w http.ResponseWriter) []models.Ticket {
	collection, _ := database.GetDBCollection()
	var result models.Admin
	collection.FindOne(context.TODO(), bson.D{}).Decode(&result)

	var res []models.Ticket
	for i, s := range result.Tickets {
		fmt.Println(i)
		if s.TicketId == Ticketid {
			fmt.Println(s)
			res = append(res, *s)

			json.NewEncoder(w).Encode(res)
			return res
		}
	}
	return res
}

func MarkTicketExpired(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	adminid := params["adminid"]
	ticketid := params["ticketid"]
	fmt.Println(adminid)

	u, _ := strconv.ParseUint(ticketid, 10, 64)
	MarkTicketExpireds(adminid, u, w)
}

func MarkTicketExpireds(adminId string, Ticketid uint64, w http.ResponseWriter) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("theatrebooking").C("admin")

	cursor := bson.M{"tickets.ticketid": Ticketid}
	fmt.Println(cursor)

	change := bson.M{"$set": bson.M{"tickets.$.expired": true}}
	erro := c.Update(cursor, change)
	if erro != nil {
		panic(erro)
	}

	// res.Result = fmt.Sprintf("Ticket No. %d  Deleted Successfully", Ticketid)

	// json.NewEncoder(w).Encode(res)
	// return

	// collection, _ := database.GetDBCollection()
	// var result models.Admin
	// collection.FindOne(context.TODO(), bson.D{}).Decode(&result)

	// var res []models.Ticket
	// for i, s := range result.Tickets {
	// 	fmt.Println(i)
	// 	if s.TicketId == Ticketid {
	// 		fmt.Println(s)
	// 		res = append(res, *s)

	// 		json.NewEncoder(w).Encode(res)
	// 		return res
	// 	}
	// }
	// return res
}
