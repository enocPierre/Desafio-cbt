package controllers

import (
	"github.com/estudo-go/Desafio-cbt/database"
	"github.com/estudo-go/Desafio-cbt/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ShowTransaction(c *gin.Context) {
	tipo := c.Param("tipo")

	newtipo, err := strconv.Atoi(tipo)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Tipo has to be String",
		})
		return
	}

	db := database.GetDatabase()

	var transacao model.Transacao
	err = db.First(&transacao, newtipo).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find transaction: " + err.Error(),
		})
		return
	}
	c.JSON(200, transacao)
}
