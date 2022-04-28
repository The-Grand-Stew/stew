package gofiber

//TODO: MOVE TO GITHUB REPO/GISTS
const GoFiberModelTemplate string = `package models

type {{ .DomainName }} struct {

}

type {{ .DomainName }}Service interface{
	Create()(any,error)
	Get()(any,error)
	GetAll()(any,error)
	Update()(any,error)
	Delete()(any,error)
}`

const GoFiberQueryTemplate string = `package queries

import "{{ .AppName }}/app/models"

type {{ .DomainName }} models.{{ .DomainName }}

func (q *{{ .DomainName }})Create()(any,error){
    // YOUR CODE GOES HERE
    return nil,nil
}

func (q *{{ .DomainName }})Get()(any,error){
    // YOUR CODE GOES HERE
    return nil,nil
}

func (q *{{ .DomainName }})GetAll()(any,error){
    // YOUR CODE GOES HERE
    return nil,nil
}

func (q *{{ .DomainName }})Update()(any,error){
    // YOUR CODE GOES HERE
    return nil,nil
}

func (q *{{ .DomainName }})Delete()(any,error){
    // YOUR CODE GOES HERE
    return nil,nil
}`

const GoFiberControllerTemplate string = `package controllers

import (
	"{{ .AppName }}/app/queries"

	"github.com/gofiber/fiber/v2"
)

var {{ .DomainName | ToLower }} *queries.{{ .DomainName }}

func Create{{ .DomainName }}(c *fiber.Ctx)error{
    //YOUR CODE GOES HERE
    return nil

}

func Get{{ .DomainName }}(c *fiber.Ctx)error{
    //YOUR CODE GOES HERE
    return nil

}

func GetAll{{ .DomainName }}(c *fiber.Ctx)error{
    //YOUR CODE GOES HERE
    return nil

}

func Update{{ .DomainName }}(c *fiber.Ctx)error{
    //YOUR CODE GOES HERE
    return nil
}

func Delete{{ .DomainName }}(c *fiber.Ctx)error{
    //YOUR CODE GOES HERE
    return nil
}`

const GoFiberRouteTemplate string = `package routes
import (
	"{{ .AppName }}/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func {{ .DomainName }}Routes(a *fiber.App){
    route:=a.Group("/")
    route.Post("/{{ .DomainName | ToLower }}",controllers.Create{{ .DomainName }})
    route.Get("/{{ .DomainName | ToLower }}/:id",controllers.Get{{ .DomainName }})
    route.Get("/{{ .DomainName | ToLower }}",controllers.GetAll{{ .DomainName }})
    route.Put("/{{ .DomainName | ToLower }}/:id",controllers.Update{{ .DomainName }})
    route.Delete("/{{ .DomainName | ToLower }}",controllers.Delete{{ .DomainName }})
}`

const GoFiberMainTemplate string = `package main

import (
	"{{ .appName }}/app/routes"
    "{{ .appName }}/pkg/utils"
	"os"

	"github.com/gofiber/fiber/v2"
)

func health(c *fiber.Ctx)error{
    return c.JSON("Health Check from {{ .appName }}")

}


func main(){
    os.Setenv("SERVER_PORT","{{ .appPort }}")
   app := fiber.New() 
   app.Get("/{{.appName}}/health", health)
   {{ .routes }}
   // Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "prod" {
		utils.StartServerWithGracefulShutdown(app)
	} else {
		utils.StartServer(app)
	}
}`
