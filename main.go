package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dstotijn/go-notion"
	"github.com/gin-gonic/gin"
)

func Send_dummy_message() {
	req, err := http.NewRequest("POST", "https://ntfy.sh/lajase_alert_1234567890",
		strings.NewReader("Je vais pouvoir commencer à jouer"))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Title", "Pense-bête")
	req.Header.Set("Priority", "urgent")
	req.Header.Set("Tags", "warning,skull")
	if _, err := http.DefaultClient.Do(req); err != nil {
		log.Fatal(err)
	}
}

func QueryDatabaseNotion(c *gin.Context) {
	SECRET := os.Getenv("NOTION_SECRET_TOKEN")
	NOTION_DB_ID := os.Getenv("NOTION_DB_ID")
	client := notion.NewClient(SECRET)

	query := notion.DatabaseQuery{
		// Filter: &notion.DatabaseQueryFilter{},
		Sorts: []notion.DatabaseQuerySort{
			{
				Property: "Name",
				// Timestamp: "",
				Direction: "ascending",
			},
		},
		// StartCursor: "",
		PageSize: 1,
	}

	db, err := client.QueryDatabase(context.Background(), NOTION_DB_ID, &query)
	if err != nil {
		fmt.Println("Big error", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": db.Results,
	})
}

func TopicSub() {
	// Attente du message retour
	resp, err := http.Get("https://ntfy.sh/lajase_response_1234567890/json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		println(scanner.Text())
	}
}

func main() {
	// lancement de la souscription pour les retours
	go TopicSub()

	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// <-c
	router := gin.Default()

	// Test
	router.GET("/test", func(c *gin.Context) {
		Send_dummy_message()
		c.JSON(http.StatusOK, gin.H{
			"message": "message sent",
		})
	})

	router.GET("/notion", func(c *gin.Context) {
		QueryDatabaseNotion(c)
	})

	router.GET("/action", func(ctx *gin.Context) {
		body := `{
	       "topic": "lajase_alert_1234567890",
	       "message": "You left the house. Turn down the A/C?",
	       "actions": [
	         {
	           "action": "http",
	           "label": "Turn down",
	           "url": "https://ntfy.sh/lajase_response_1234567890",
	           "body": "{\"temperature\": 65}"
	         }
	       ]
	   }`
		req, _ := http.NewRequest("POST", "https://ntfy.sh/", strings.NewReader(body))
		http.DefaultClient.Do(req)
	})

	router.Run(":8080")
}
