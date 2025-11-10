package api

import (
	"cloudnativedemo/user-service/internals/repository"
	"cloudnativedemo/user-service/internals/service"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGETUsers_Return200AndList(t *testing.T){

	if err := godotenv.Load(`../../.env.test`); err != nil {
		t.Fatalf("Kunde inte ladda test env filen %v", err)
	}

	user := os.Getenv("TEST_DB_USER")
	pass := os.Getenv("TEST_DB_PASS")
	host := os.Getenv("TEST_DB_HOST")
	port := os.Getenv("TEST_DB_PORT")
	name := os.Getenv("TEST_DB_NAME")

	if user == "" || pass == "" || host == "" || port == "" || name == "" {
		t.Fatalf("Saknar värden från test.env")
	}


	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",user,pass,host,port,name)


	repo := repository.NewUserRepository(dsn)
	service := service.NewUserService(repo)


	if _, err := service.CreateUser("Yahya", 34); err != nil {
		t.Fatalf("setup failade %v", err)
	}

	if _, err := service.CreateUser("Martin", 25); err != nil {
		t.Fatalf("Setup failade %v", err)
	}


	gin.SetMode(gin.TestMode) //"Mindre loggning i test"


	r := gin.New();

	r.Use(gin.Recovery()) //Fångar "krash/panic" -> 500 istället för att krascha testet


	r.GET("/users", func(c *gin.Context) {

		users, err := service.ListUsers()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "kunde inte hämta användaren"})			
			return
		}
		c.JSON(http.StatusOK, users)
	})


	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)


	assert.Equal(t, http.StatusOK, rec.Code, "ska ge oss 200 ok")



	var response []map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("kunde inte avkoda json %v %s", err, rec.Body.String())
	}

	assert.Len(t, response, 2, "Ska vara 2 users")

}