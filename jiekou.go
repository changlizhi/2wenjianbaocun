package main

import (
  "github.com/gin-gonic/gin"
  "log"
  "net/http"
  "fmt"
  "net/url"
  // "github.com/json-iterator/go"
)
type Tone struct{
  Name string
  Age int64
}
func main() {
	router := gin.Default()

  // var json = jsoniter.ConfigCompatibleWithStandardLibrary
  // var data =Tone{}

  router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["files"]
    log.Println("files",files)
		for _, file := range files {
			log.Println(file.Filename)

			// 上传文件至指定目录
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
  router.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostForm("names")

    m, _ := url.ParseQuery(names)
    log.Println("m------",m)

    // json.Unmarshal(m, &data)
    // log.Println("data------",data)

    fmt.Printf("\nids: %v;\n names: %v\n", ids, names)
	})
	router.Run(":8080")
}
