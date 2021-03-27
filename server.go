package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type router []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	IP   string `json:"ip"`
}

func indexPage(c *gin.Context) {
	if c.Param("version") == "v1" {
		c.HTML(200, "v1.tmpl", nil)
	} else if c.Param("version") == "" {
		c.HTML(200, "v2.tmpl", nil)
	} else {
		pageNotAvailable(c)
	}
}

func LookingGlass(c *gin.Context) {
	Router := c.PostForm("Router")
	Action := c.PostForm("Action")
	IP := c.PostForm("IP")

	data, err := ioutil.ReadFile("router.json")

	var router router
	json.Unmarshal(data, &router)

	var RouterIP string
	RouterIP = "NULL"
	for n := range router {
		if router[n].Name == Router {
			fmt.Println("Router IP: ", router[n].IP)
			RouterIP = router[n].IP
		}
	}

	RouterURL := "http://" + RouterIP + ":32991"
	resp, err := http.PostForm(RouterURL,
		url.Values{"Action": {Action}, "IP": {IP}})
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	c.String(200, string(body))
}

func pageNotAvailable(c *gin.Context) {
	c.HTML(404, "404.tmpl", nil)
}

func main() {

	fmt.Print("\n")
	fmt.Print("-------------------\n")
	fmt.Print("\n")
	fmt.Print("FRRouting Looking Glass\n")
	fmt.Print("Port listing at 32280\n")
	fmt.Print("Repo: https://github.com/steveyiyo/frrouting-lg\n")
	fmt.Print("Author: SteveYi\n")
	fmt.Print("Demo: https://lg.steveyi.net\n")
	fmt.Print("\n")
	fmt.Print("-------------------\n")
	fmt.Print("\n")

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.LoadHTMLGlob("static/*")

	router.GET("/", indexPage)
	router.GET("/:version", indexPage)
	router.POST("/", LookingGlass)

	router.NoRoute(pageNotAvailable)

	router.Run("127.0.0.1:32280")

}
