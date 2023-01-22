package main

import (
	// "context"
	// "database/sql"

	"log"
	"tour-of-heroes-api-go/controllers"
	"tour-of-heroes-api-go/models"

	"github.com/gin-gonic/gin"
	// _ "github.com/microsoft/go-mssqldb"
)

// var db *sql.DB

var server = "localhost"
var port = 1433
var user = "sa"
var password = "Password1!"
var database = "heroes"

// type hero struct {
// 	ID          int    `json:"id"`
// 	Name        string `json:"name"`
// 	AlterEgo    string `json:"alterEgo"`
// 	Description string `json:"description"`
// }

func main() {

	log.SetPrefix("tour-of-heroes-api: ")
	log.SetFlags(0)
	log.Print("Connecting to the database")

	// ConnectDb()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Tour of Heroes API",
		})
	})

	router.GET("/api/hero", controllers.GetHeroes)
	router.POST("/api/hero", controllers.CreateHero)
	// router.GET("/api/hero/:id", getHeroById)

	models.ConnectDatabase()

	router.Run("localhost:8080")
}

// Establishes SQL Server connection
// func ConnectDb() {
// 	db, err := gorm.Open(sqlserver.Open("sqlserver://sa:Password1!@localhost:1433?database=heroes"), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Error creating connection pool: ", err.Error())
// 	}
// 	log.Printf("Connected!\n")
// }

// func ConnectDb() {

// 	// Build connection string
// 	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
// 		server, user, password, port, database)

// 	var err error
// 	// Create a connection pool
// 	db, err = sql.Open("sqlserver", connString)

// 	if err != nil {
// 		log.Fatal("Error creating connection pool: ", err.Error())
// 	}

// 	ctx := context.Background()

// 	err = db.PingContext(ctx)

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	log.Printf("Connected!\n")
// }

// func getHeroes(c *gin.Context) {
// 	log.Printf("getHeroes")
// 	ctx := context.Background()

// 	if db == nil {
// 		err := errors.New("getHeroes: db is null")
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
// 	}

// 	// Check if the database is alive
// 	err := db.PingContext(ctx)

// 	if err != nil {
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
// 	}

// 	// Prepare query
// 	tsql := fmt.Sprintf("SELECT Id,Name,AlterEgo,Description FROM Heroes;")

// 	// Execute query
// 	rows, err := db.QueryContext(ctx, tsql)

// 	if err != nil {
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
// 	}

// 	defer rows.Close()

// 	var heroes []hero

// 	// Iterate through the result set.
// 	for rows.Next() {
// 		var name, description, alterEgo string
// 		var id int

// 		// Get values from row.
// 		err := rows.Scan(&id, &name, &alterEgo, &description)
// 		if err != nil {
// 			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
// 		}

// 		fmt.Printf("ID: %d, Name: %s\n", id, name)
// 		hero := hero{ID: id, Name: name, AlterEgo: alterEgo, Description: description}
// 		heroes = append(heroes, hero)
// 	}

// 	c.IndentedJSON(http.StatusOK, heroes)
// }

// func getHeroById(c *gin.Context) {
// 	log.Printf("getHeroById")

// 	id := c.Param("id")

// 	ctx := context.Background()

// 	if db == nil {
// 		err := errors.New("getHeroes: db is null")
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
// 	}

// 	// Check if the database is alive
// 	err := db.PingContext(ctx)

// 	if err != nil {
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
// 	}

// 	// Prepare query
// 	tsql := fmt.Sprintf("SELECT Id,Name,AlterEgo,Description FROM Heroes WHERE Id=%v;", id)

// 	log.Print(tsql)

// 	// Execute query
// 	rows, err := db.Query(tsql)

// 	//Check how many rows were returned
// 	if rows.Next() == false {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Hero not found"})
// 	}

// 	// Check if there are any errors
// 	if err != nil {
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
// 	}

// 	defer rows.Close()

// 	// Check how many rows were returned
// 	if rows.Next() == false {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Hero not found"})
// 	}

// 	// Get values from row.
// 	var name, description, alterEgo string
// 	var idInt int

// 	for rows.Next() {
// 		err := rows.Scan(&idInt, &name, &alterEgo, &description)
// 		if err != nil {
// 			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
// 		}

// 		log.Printf("ID: %d, Name: %s\n", idInt, name)
// 	}

// 	hero := hero{ID: idInt, Name: name, AlterEgo: alterEgo, Description: description}

// 	c.IndentedJSON(http.StatusOK, hero)

// }
