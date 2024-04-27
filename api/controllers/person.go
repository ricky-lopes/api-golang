package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricky-lopes/api/entities"
)

type PersonController struct {
	persons []entities.Person
}

func NewPersonControler() *PersonController {
	return &PersonController{}
}

func (p *PersonController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, p.persons)
}

func (p *PersonController) Create(ctx *gin.Context) {
	person := entities.NewPerson()

	if err := ctx.BindJSON(&person); err != nil {
		return
	}

	p.persons = append(p.persons, *person)

	ctx.JSON(http.StatusOK, person)
}

func (p *PersonController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	for idx, person := range p.persons {
		if person.ID == id {
			p.persons = append(p.persons[0:idx], p.persons[idx+1:]...)
			ctx.JSON(http.StatusOK, person.ID)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Person not found",
	})
}
