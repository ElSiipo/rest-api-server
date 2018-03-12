package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func GetInstructions(c *gin.Context) {
	session, err := mgo.Dial("")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB("test").C("people")
	err = collection.Insert(&Instruction{ID: "Ale", EventStatus: "+55 53 8116 9639", EventName: ""})

	if err != nil {
		log.Fatal(err)
	}

	// result := Instruction{}
	// err = c.Find(bson.M{"name": "Ale"}).One(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Phone:", result.Phone)

	// _, err := dbmap.Select(&instructions, "SELECT * FROM instruction")
	// c.JSON(200, instructions)
	// curl -i http://localhost:8080/api/v1/instructions
}

// func GetInstruction(c *gin.Context) {
// 	instructionID := c.Param("id")

// 	for _, item := range instructions {
// 		if item.ID == instructionID {
// 			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": item})
// 		}
// 	}
// }

// func PostInstruction(c *gin.Context) {
// 	instructionID := c.Param("id")
// 	eventStatus := c.Param("event_status")
// 	eventName := c.Param("event_name")

// 	instructions = append(instructions, Instruction{ID: instructionID, EventStatus: eventStatus, EventName: eventName})
// 	c.JSON(200, gin.H{"ok": "POST api/v1/instructions"})
// }

// func UpdateInstruction(c *gin.Context) {
// 	c.JSON(200, gin.H{"ok": "PUT api/v1/instructions/1"})
// }

// func DeleteInstruction(c *gin.Context) {
// 	c.JSON(200, gin.H{"ok": "DELETE api/v1/instructions/1"})
// }
