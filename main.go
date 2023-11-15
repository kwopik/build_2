package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.String(200, "Hello!!!!!test2!!!!!!!!!!!1, Go with Docker8086 test branchtest!")
    })
    r.Run(":8080")
}
