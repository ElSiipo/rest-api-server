package app

import (
	"fmt"
	"log"
	"rest-api-server/models"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func GetInstructions(c *gin.Context) {

	session, err := mgo.Dial("mongodb://earl:AllenWalker1!@ds111279.mlab.com:11279/vera")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	fmt.Println("Creating db session...")
	collection := session.DB("instructions").C("instruction")
	var results []models.Instruction

	err = collection.Find(nil).All(&results)
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

	session, err := mgo.Dial("mongodb://earl:AllenWalker1!@ds111279.mlab.com:11279/vera")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB("test").C("instruction")
	err = collection.Insert(&models.Instruction{ID: "Ale", EventStatus: "Test Event Status", EventName: "Test event name"})

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
