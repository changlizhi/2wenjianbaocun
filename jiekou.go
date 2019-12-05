package 2wenjianbaocun
import(
  "github.com/gin-gonic"
  "github.com/json-iterator/go"
)
func main() {
	router := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
    var json = jsoniter.ConfigCompatibleWithStandardLibrary
    json.Marshal(&file.Filename)
		// 上传文件至指定目录
		// c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":8080")
}
