package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func NewCompany(c *gin.Context) {
	var newCompany Company
	if err := c.ShouldBindJSON(&newCompany); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newCompany.ID = xid.New().String()
	companies = append(companies, newCompany)
	c.JSON(http.StatusCreated, newCompany)
}

func GetCompanies(c *gin.Context) {
	c.JSON(http.StatusOK, companies)
}

func HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API rodando"})
}

func UpdateCompany(c *gin.Context) {
	id := c.Param("id")
	var company Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	index := -1
	for i := 0; i < len(companies); i++ {
		if companies[i].ID == id {
			index = 1
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Company not found",
		})
		return
	}

	companies[index] = company
	c.JSON(http.StatusOK, company)

}

func DeleteCompany(c *gin.Context) {
	id := c.Param("id")
	index := -1
	for i := 0; i < len(companies); i++ {
		if companies[i].ID == id {
			index = 1
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Company not found",
		})
		return
	}

	companies = append(companies[:index], companies[index+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"message": "Company has been deleted",
	})
}
