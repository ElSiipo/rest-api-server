package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"rest-api-server/models"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetConnectionFromFile() string {

	absPath, _ := filepath.Abs("../rest-api-server/app/connectionString.txt")

	byteArray, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(err)
	}

	return string(byteArray[:])
}

func GetInstructions(c *gin.Context) {
	s := GetConnectionFromFile()
	session, err := mgo.Dial(s + "/vera")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	fmt.Println("Creating db session...")
	collection := session.DB("vera").C("instructions")
	var results []models.Instruction

	err = collection.Find(bson.M{}).All(&results)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Results All: ", results)
	}

	// c.JSON(200, instructions)
	// curl -i http://localhost:8080/api/v1/instructions
}

func PostInstruction(c *gin.Context) {
	// instructionID := c.Param("id")
	// eventStatus := c.Param("event_status")
	// eventName := c.Param("event_name")

	fmt.Println("dialing mongodb..")

	s := GetConnectionFromFile()
	session, err := mgo.Dial(s + "/vera")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB("test").C("instruction")
	err = collection.Insert(&models.Instruction{EventStatus: "Test Event Status", EventName: "Test event name"})

	if err != nil {
		log.Fatal(err)
	}

	// instructions = append(instructions, Instruction{ID: instructionID, EventStatus: eventStatus, EventName: eventName})
	// c.JSON(200, gin.H{"ok": "POST api/v1/instructions"})
}

// func UpdateInstruction(c *gin.Context) {
// 	c.JSON(200, gin.H{"ok": "PUT api/v1/instructions/1"})
// }

// func DeleteInstruction(c *gin.Context) {
// 	c.JSON(200, gin.H{"ok": "DELETE api/v1/instructions/1"})
// }
