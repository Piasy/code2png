package controllers

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
	"strconv"
	"strings"

	"github.com/astaxie/beego"

	"github.com/Piasy/code2png/models"

)

var LANGS map[string]string = map[string]string {
	"java"		: "java",
	"python"	: "py",
	"c++"		: "cpp",
	"javascript"	: "js",
}

// Operations about object
type CodeImagesController struct {
	beego.Controller
}

// @Title Create
// @Description create image for source code
// @Param	lang		query 	string	true		"language of code"
// @Param	code		query 	string	true		"content of code"
// @Success 200 {object} models.CodeImage
// @Failure 403 body is empty
// @router / [post]
func (this *CodeImagesController) Post() {
	lang := LANGS[this.GetString("lang")]
	if len(lang) == 0 {
		this.Ctx.ResponseWriter.WriteHeader(403)
		return
	}
	code := this.GetString("code")
	beego.Info(code)
	ts := time.Now().Unix()
	name := "tmp/code_" + strconv.FormatInt(ts, 10) + "." + lang
	code = strings.Replace(code, "\\r", "\r", -1)
	code = strings.Replace(code, "\\n", "\n", -1)
	ioutil.WriteFile(name, []byte(code), 0644)
	cmd := exec.Command("./code2png.sh", name)
	cmd.Start()
	cmd.Wait()
	image := models.CodeImage{Url: fmt.Sprintf("http://code2png.babits.top/images/code_%d.%s.png", ts, lang)}
	this.Data["json"] = image
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.ServeJSON()
}

