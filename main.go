package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"net/http"
	"os"
	"strings"
)

type Env struct {
	WebsiteTitle       string `yaml:"website_title" json:"website_title"`
	WebsiteAuthor      string `yaml:"website_author" json:"website_author"`
	WebsiteDescription string `yaml:"website_description" json:"website_description"`
	WebsiteKeywords    string `yaml:"website_keywords" json:"website_keywords"`
	Logo               string `yaml:"logo" json:"logo"`

	GithubUsername string `json:"github_username" yaml:"github_username"`
	Email          string `json:"email" yaml:"email"`
	Pay            string `json:"pay" yaml:"pay"`
	Qq             string `json:"qq" yaml:"qq"`

	Sites    []Site    `json:"sites"`
	Projects []Project `json:"projects"`
}

type (
	Site struct {
		Title string `json:"title" yaml:"title"`
		Url   string `json:"url" yaml:"url"`
		Desc  string `json:"desc" yaml:"desc"`
	}

	Project struct {
		Title string `json:"title" yaml:"title"`
		Url   string `json:"url" yaml:"url"`
		Desc  string `json:"desc" yaml:"desc"`
	}
)

var cfg Env

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		if _, err := os.Stat("config.yml"); err != nil {
			c.JSON(500, fmt.Sprintf("文件 [config.yml] 不存在！"))
			c.Abort()
		}

		content, _ := os.ReadFile("config.yml")
		_ = yaml.NewDecoder(strings.NewReader(string(content))).Decode(&cfg)

		b, _ := json.Marshal(cfg)

		var body map[string]interface{}
		_ = json.Unmarshal(b, &body)
		c.HTML(http.StatusOK, "index.html", body)
	})
	_ = r.Run()
}
