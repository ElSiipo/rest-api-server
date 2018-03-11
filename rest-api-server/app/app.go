package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Instruction struct {
	ID          string `db:"id" json:"id"`
	EventStatus string `db:"event_status" json:"event_status"`
	EventName   string `db:"event_name" json:"event_name"`
}

var instructions []Instruction

func GetInstructions(c *gin.Context) {
	// _, err := dbmap.Select(&instructions, "SELECT * FROM instruction")
	c.JSON(200, instructions)
	// curl -i http://localhost:8080/api/v1/instructions
}

func GetInstruction(c *gin.Context) {
	instructionID := c.Param("id")

	for _, item := range instructions {
		if item.ID == instructionID {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": item})
		}
	}
}

func PostInstruction(c *gin.Context) {
	instructionID := c.Param("id")
	eventStatus := c.Param("event_status")
	eventName := c.Param("event_name")

	instructions = append(instructions, Instruction{ID: instructionID, EventStatus: eventStatus, EventName: eventName})
	c.JSON(200, gin.H{"ok": "POST api/v1/instructions"})
}

func UpdateInstruction(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "PUT api/v1/instructions/1"})
}

func DeleteInstruction(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "DELETE api/v1/instructions/1"})
}
