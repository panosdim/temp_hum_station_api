package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

// Update with correct token
const token = ""

func init() {
	var err error
	//open a db connection
	db, err = gorm.Open("sqlite3", "sensorData.db")
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	db.AutoMigrate(&sensorDataModel{})
}

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/", createRecord)
		v1.GET("/", fetchAll)
	}
	router.Run()

	defer db.Close()
}

type sensorDataModel struct {
	gorm.Model
	Date        uint32  `json:"date"`
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
}

// createRecord add a new record in database
func createRecord(c *gin.Context) {
	if token != c.PostForm("token") {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid token provided"})
	} else {
		temperature, _ := strconv.ParseFloat(c.PostForm("temperature"), 2)
		humidity, _ := strconv.ParseFloat(c.PostForm("humidity"), 2)
		date, _ := strconv.ParseUint(c.PostForm("date"), 10, 32)
		data := sensorDataModel{Date: uint32(date), Temperature: float32(temperature), Humidity: float32(humidity)}
		db.Create(&data)
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Record created successfully!", "resourceId": data.ID})
	}
}

// fetchAll fetch all records from database
func fetchAll(c *gin.Context) {
	var records []sensorDataModel

	db.Find(&records)

	if len(records) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No records found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": records})
}
