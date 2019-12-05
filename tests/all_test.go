package tests

import(
  "github.com/imroc/req"
  "testing"
  "log"
  "os"
  "net/url"
  "github.com/json-iterator/go"

)

func TestReq(t *testing.T){
  r := req.New()
  resp ,_ :=r.Get("https://www.hanfuxin.com")
  log.Println("resp---",resp)

  resp ,_ =r.Get("https://www.hanfuxin.com/myweb")
  log.Println("resp---",resp)
}
func TestFileupload(t *testing.T){
  // r :=req.New()
  // resp,_:=r.Post("http://localhost:8080/upload",
  //    req.File("C:/Users/clz/allworkspaces/diannaobaozhi/1.jpg"),
  //    req.File("C:/Users/clz/allworkspaces/diannaobaozhi/1.jpg"))
    file1, _ := os.Open("C:/Users/clz/allworkspaces/diannaobaozhi/1.jpg")
    file2, _ := os.Open("C:/Users/clz/allworkspaces/diannaobaozhi/2.jpg")
    progress := func(current, total int64) {
    	log.Println(float32(current)/float32(total)*100, "%")
    }
    resp,_:=req.Post("http://localhost:8080/upload", req.FileUpload{
      File:      file1,
      FieldName: "files",       // FieldName 是表单字段名
      FileName:  "avatar1.png", // Filename 是要上传的文件的名称，我们使用它来猜测mimetype，并将其上传到服务器上
    },req.FileUpload{
      File:      file2,
      FieldName: "files",       // FieldName 是表单字段名
      FileName:  "avatar2.png", // Filename 是要上传的文件的名称，我们使用它来猜测mimetype，并将其上传到服务器上
    },req.UploadProgress(progress))
    log.Println("upload complete")
    log.Println("resp---",resp);
}
// func TestFiledownload(t *testing.T){
//   progress := func(current, total int64) {
//   	log.Println(float32(current)/float32(total)*100, "%")
//   }
//   r, _ := req.Get("https://www.hanfuxin.com/img/login-bg.8ab74202.jpg", req.DownloadProgress(progress))
//   r.ToFile("login.jpg")
//   log.Println("download complete")
// }
type Tone struct{
  Name string
  Age int64
}
func TestParam(t *testing.T){
  var json = jsoniter.ConfigCompatibleWithStandardLibrary
  var data =Tone{
    Name:"ccc",
    Age:11111,
  }
  strbyte,_ := json.Marshal(&data)
  log.Println("str---",string(strbyte))
  v := url.Values{}
  v.Add("name", string(strbyte))
  body := v.Encode()
  log.Println("v---",v)
  log.Println("body---",body)
  // url decode
  param := req.Param{
  	"names":  body,
  	"pwd": "roc",
  }
  resp,_:=req.Post(
    "http://localhost:8080/post?ids[a]=123&ids[b]=abc",param)
  log.Println("resp-------",resp)

}
