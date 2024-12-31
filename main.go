package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/refinepoint/mock-api-response/mockdata"
)

type MockData struct {
	Method string          `json:"method"`
	Data   json.RawMessage `json:"data"`
}

// Function to scan mock data rows
func scanMockData(rows *sql.Rows) ([]*MockData, error) {
	defer rows.Close()

	mockData := make([]*MockData, 0)
	for rows.Next() {
		var m MockData
		//err := rows.Scan(&m.ID, &m.CustomDomain, &m.ProjectID, &m.CreatedBy, &m.Method, &m.Data)
		err := rows.Scan(&m.Data)
		if err != nil {
			return nil, err
		}

		mockData = append(mockData, &m)

	}
	return mockData, nil
}

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Unable to load environment proeprties")
	}

	// Postgres connection pooling setup
	idleConn := 4
	maxConnections := 4
	maxConnLifetime := 2 * time.Minute
	poolConn, err := sqlx.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer poolConn.Close()
	poolConn.SetMaxOpenConns(maxConnections)
	poolConn.SetMaxIdleConns(idleConn)
	poolConn.SetConnMaxLifetime(maxConnLifetime)

	// Initialize the HTTP router
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost",
			"https://mock-api.refinepoint.com",
			"https://mock-response.refinepoint.com",
			"https://mockapi.refinepoint.com",
			"https://mockapi-docs.refinepoint.com",
		},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost" ||
				origin == "https://mock-api.refinepoint.com" ||
				origin == "https://mock-response.refinepoint.com" ||
				origin == "https://mockapi.refinepoint.com" ||
				origin == "https://mockapi-docs.refinepoint.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	store := persistence.NewInMemoryStore(time.Hour)

	router.StaticFile("/", "./index.html")

	router.GET("/mockdata/samples", cache.CachePage(store, time.Hour, func(c *gin.Context) {
		// Get the size query parameter
		sizeParam := c.DefaultQuery("_size", "1") // Default size is 128
		size, err := strconv.Atoi(sizeParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid size parameter"})
			return
		}

		// Get the delay query parameter
		delayParam := c.DefaultQuery("_delay", "0") // Default delay is 0 (no delay)
		delay, err := strconv.Atoi(delayParam)
		if err != nil || delay < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid delay parameter"})
			return
		}
		// Get the jsonFields query parameter
		jsonFieldsParam := c.DefaultQuery("_fields", "") // Default to empty (no fields specified)

		// Set encloseInArray based on size
		encloseInArrayStr := "true"
		if size == 1 {
			encloseInArrayStr = c.DefaultQuery("_encloseInArray", "false")
		}
		encloseInArray, _ := strconv.ParseBool(encloseInArrayStr)

		jsonFields := strings.Split(jsonFieldsParam, ",") // Split into a slice

		for i, field := range jsonFields {
			jsonFields[i] = strings.TrimSpace(field) // Trim spaces around each field
		}

		// Create a map for additional query parameters
		additionalParams := make(map[string]string)

		// Loop through all query parameters manually
		for paramName, values := range c.Request.URL.Query() {
			// Skip common parameters
			if strings.HasPrefix(paramName, "_") {
				continue
			}
			if !isValidParamName(values[0]) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameter: " + paramName + "system columns(size , delay , encloseInArray , type should begin with _ )"})
				return
			}
			additionalParams[paramName] = values[0] // Single value, store directly

		}

		if delay > 0 {
			time.Sleep(time.Duration(delay) * time.Millisecond) // Delay in milliseconds
		}

		dummyData, err := mockdata.GenerateCustomJSON(size, jsonFields, additionalParams, encloseInArray)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, dummyData)
	}))

	router.Any("/c/:projectid/:id", cache.CachePage(store, time.Hour, func(c *gin.Context) {

		method := c.Request.Method
		log.Printf("Request method: %s, Path: %s", method, c.Request.URL.Path)

		projectID := c.Param("projectid")
		id := c.Param("id")

		query := "SELECT data FROM refinepoint.mockdata WHERE projectid = $1 AND id = $2 AND method  = upper( $3)"
		rows, err := poolConn.Query(query, projectID, id, method)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		mockData, err := scanMockData(rows)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if len(mockData) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No such url found"})
			return
		}

		c.Data(http.StatusOK, "application/json", mockData[0].Data)
	}))

	// Route to create mock data
	router.POST("/mockdata", func(c *gin.Context) {
		projectID := c.DefaultQuery("projectid", "")
		method := c.DefaultQuery("method", "GET")
		methodQuery := strings.ToUpper(method)

		if projectID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ProjectID is required"})
			return
		}

		// Get the created_by value from the request header
		// Get the created_by value from the request header
		createdBy := c.GetHeader("Created-By")

		// Use the default value if the header is missing
		if createdBy == "" {
			createdBy = "default@refinepoint.com"
		}

		// Bind the entire request body (without wrapping in a data block)
		var requestBody json.RawMessage
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate that the 'data' field is valid JSON

		// Insert into the mockdata table and get the generated id
		var newID string
		query := "INSERT INTO refinepoint.mockdata (customdomain, projectid, created_by, method, data) VALUES (NULL, $1, $2, $3, $4) RETURNING id"
		err := poolConn.QueryRow(query, projectID, createdBy, methodQuery, requestBody).Scan(&newID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the generated ID in the response
		c.JSON(http.StatusCreated, gin.H{"message": "Mock data created", "id": newID})
	})

	// Start the HTTP server
	if err := router.Run(":8000"); err != nil {
		log.Fatalf("Unable to start HTTP server: %v\n", err)
	}
}

// Function to check if the query parameter name is valid
func isValidParamName(paramName string) bool {
	validParams := []string{"uuid", "uid", "id", "number", "decimal", "float", "long", "integer", "text", "string", "phone", "boolean", "date", "datetime", "time", "email", "mail"}
	for _, valid := range validParams {
		if paramName == valid {
			return true
		}
	}
	return false
}
