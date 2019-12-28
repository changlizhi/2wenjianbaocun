package main

import (
  "github.com/gin-gonic/gin"
  "log"
  "net/http"
  "fmt"
  "net/url"
  "github.com/json-iterator/go"
)
type Tone struct{
  Name string
  Age int64
}
////// 跨域
func Cors() gin.HandlerFunc {
  return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
func main() {
	router := gin.Default()
  router.Use(Cors())
  var json = jsoniter.ConfigCompatibleWithStandardLibrary
  var data =Tone{}
  //router.OPTIONS("/upload", func(c *gin.Context){} )
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
    log.Println("m------",m["name"][0])

    json.Unmarshal([]byte(m["name"][0]), &data)
    log.Println("data.Name------",data.Name)

    fmt.Printf("\nids: %v;\n names: %v\n", ids, names)
	})
	router.Run(":8081")
}
