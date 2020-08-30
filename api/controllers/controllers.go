package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/DeVil2O/moviebookingsystem/api/database"
	"github.com/DeVil2O/moviebookingsystem/api/models"
	"github.com/dgrijalva/jwt-go"
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

	CreateTickets("mainAdmin", "Chirag Garg", "8218517963")

}

func CreateTickets(adminId string, Customername string, Phonenumber string) {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, _ := flake.NextID()
	ticket := &models.Ticket{id, Customername, Phonenumber, time.Now(), time.Now().Add(180 * time.Minute), time.Now(), false}
	collection, err := database.GetDBCollection()

	cursor, err := collection.Find(context.TODO(), bson.M{"adminid": adminId})

	change := bson.M{"$push": bson.M{"admin.$.tickets": ticket}}
	erro := collection.UpdateOne(cursor, change)
	if erro != nil {
		panic(erro)
	} else {
		fmt.Println("success")
	}
	if err != nil {
		log.Fatal(err)
	}
	var episodes []bson.M
	if err = cursor.All(context.TODO(), &episodes); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodes)
}
