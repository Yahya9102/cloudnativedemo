package api

import (
	"cloudnativedemo/s3-uploader/internals/s3client"
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

func StartServer(s3c *s3client.Client) {

	r := gin.Default()


	r.POST("/upload", func(c *gin.Context) {

		var input struct {
			Key string `json:"key"`
			Content string `json:"content"`
		}

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		if input.Key == "" || input.Content == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "key and content are required"})
			return
		}

		if err := s3c.PutString(context.Background(), input.Key, input.Content); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "file uploaded successfully",
			"bucket": s3c.Bucket,
			"key": input.Key,
		} )

	})



	r.DELETE("/delete/:key", func(c *gin.Context) {
		objectKey := c.Param("key")
		key, err := url.PathUnescape(objectKey) // katt bilder -> katt%sbilder -> kattbilder

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid key in url",
			})
			return
		}


		if err := s3c.DeleteObject(c.Request.Context(), key); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Status(http.StatusNoContent)



	})

	r.GET("/list", func(c *gin.Context) {
		prefix := c.Query("prefix")
		ttl := 15 * time.Minute

		items, err := s3c.ListObjects(c.Request.Context(), prefix, ttl)
		if  err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
			
		}


		c.JSON(http.StatusOK, items)

	})
	r.Run(":8082")			

}