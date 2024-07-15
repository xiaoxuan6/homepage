package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"net/http"
)

type Env struct {
	WebsiteTitle       string `env:"website_title" json:"website_title"`
	WebsiteAuthor      string `env:"website_author" json:"website_author"`
	WebsiteDescription string `env:"website_description" json:"website_description"`
	WebsiteKeywords    string `env:"website_keywords" json:"website_keywords"`
	Logo               string `env:"logo" json:"logo"`

	GithubUsername string `json:"github_username" env:"github_username"`
	Email          string `json:"email" env:"email"`
	Pay            string `json:"pay" env:"pay"`
	Qq             string `json:"qq" env:"qq"`

	// site
	Site1Url   string `json:"site_1_url" env:"site_1_url"`
	Site1Title string `json:"site_1_title" env:"site_1_title"`
	Site1Desc  string `json:"site_1_desc" env:"site_1_desc"`

	Site2Url   string `json:"site_2_url" env:"site_2_url"`
	Site2Title string `json:"site_2_title" env:"site_2_title"`
	Site2Desc  string `json:"site_2_desc" env:"site_2_desc"`

	Site3Url   string `json:"site_3_url" env:"site_3_url"`
	Site3Title string `json:"site_3_title" env:"site_3_title"`
	Site3Desc  string `json:"site_3_desc" env:"site_3_desc"`

	Site4Url   string `json:"site_4_url" env:"site_4_url"`
	Site4Title string `json:"site_4_title" env:"site_4_title"`
	Site4Desc  string `json:"site_4_desc" env:"site_4_desc"`

	// project
	Project1Url   string `json:"project_1_url" env:"project_1_url"`
	Project1Title string `json:"project_1_title" env:"project_1_title"`
	Project1Desc  string `json:"project_1_desc" env:"project_1_desc"`

	Project2Url   string `json:"project_2_url" env:"project_2_url"`
	Project2Title string `json:"project_2_title" env:"project_2_title"`
	Project2Desc  string `json:"project_2_desc" env:"project_2_desc"`

	Project3Url   string `json:"project_3_url" env:"project_3_url"`
	Project3Title string `json:"project_3_title" env:"project_3_title"`
	Project3Desc  string `json:"project_3_desc" env:"project_3_desc"`

	Project4Url   string `json:"project_4_url" env:"project_4_url"`
	Project4Title string `json:"project_4_title" env:"project_4_title"`
	Project4Desc  string `json:"project_4_desc" env:"project_4_desc"`
}

var cfg Env

func main() {

	_ = godotenv.Load()
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(fmt.Sprintf("读取 env 错误：%s", err.Error()))
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		b, _ := json.Marshal(cfg)

		var body map[string]interface{}
		_ = json.Unmarshal(b, &body)
		c.HTML(http.StatusOK, "index.html", body)
	})
	_ = r.Run()
}
