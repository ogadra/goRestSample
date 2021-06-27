package main

import (
	"database/sql/driver"
	"log"
	"net/http"
	"os"
	"time"
	"fmt"

	"github.com/ant0ine/go-json-rest/rest"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var timeZoneJST = time.FixedZone("Asia/Tokyo", 9*60*60)

type DBTime struct {
	time.Time
}

func (t DBTime) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.In(timeZoneJST).Format(`"2006-01-02 15:04:05"`)), nil
}

func (t DBTime) MarshalText() ([]byte, error) {
    return []byte(t.Time.In(timeZoneJST).Format(`"2006-01-02 15:04:05"`)), nil
}

func (t *DBTime) Scan(value interface{}) error {
	t.Time = value.(time.Time).In(timeZoneJST)
	return nil
}

func (t DBTime) Value() (driver.Value, error){
	return t.Time, nil
}

type Impl struct {
    DB *gorm.DB
}

func connect(path string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql",path)
	if err != nil{
		log.Println("Not ready")
		time.Sleep(time.Second)
		return connect(path)
	}
	return db, nil
}

func (i *Impl) InitDB() {

    // MySQLとの接続。ユーザ名：gorm パスワード：password DB名：country
	var err error
	dsn := "go_test:" + os.Getenv("PASS") + "@tcp(db:3306)/go_database?parseTime=true"
    i.DB, err = connect(dsn)
	if err != nil{
		log.Println("DB connection error")
	}

    i.DB.LogMode(true)
}

// DBマイグレーション
func (i *Impl) InitSchema() {
    i.DB.AutoMigrate(&Recipe{})
}

type Recipe struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Making_time string `json:"making_time"`
	Serves string `json:"serves"`
	Ingredients string `json:"ingredients"`
	Cost int `json:"cost"`
	Created_at DBTime `json:"created_at"`
	Updated_at DBTime `json:"updated_at"`
}

func main() {

	err := godotenv.Load(fmt.Sprintf("%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal(err)
	}

	i := Impl{}
	i.InitDB()
	i.InitSchema()

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/recipes", i.GetAllRecipes),
		rest.Post("/recipes", i.PostRecipe),
		rest.Get("/recipes/:id", i.GetRecipe),
		rest.Patch("/recipes/:id", i.PatchRecipe),
		rest.Delete("/recipes/:id", i.DeleteRecipe),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("server started.")
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}