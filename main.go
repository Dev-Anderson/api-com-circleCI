package main

import (
	"github.com/gin-gonic/gin"
)

type Company struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	CEO     string `json:"ceo"`
	Revenue string `json:"revenue"`
}

var companies = []Company{
	{ID: "1", Name: "Dell", CEO: "Michael Dell", Revenue: "92.2 billion"},
	{ID: "2", Name: "Netflix", CEO: "Reed Hastings", Revenue: "20.2 billion"},
	{ID: "3", Name: "Microsoft", CEO: "Satya Nadella", Revenue: "320 million"},
	{ID: "4", Name: "Microsoft outro", CEO: "Satya Nadella asdf", Revenue: "30 million"},
}

func main() {
	router := gin.Default()
	router.GET("/", HomePage)
	router.GET("/companies", GetCompanies)
	router.POST("/company", NewCompany)
	router.PUT("/company/:id", UpdateCompany)
	router.DELETE("/company/:id", DeleteCompany)
	router.Run()
}
