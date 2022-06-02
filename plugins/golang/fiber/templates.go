package fiber

import (
	"path/filepath"
	"stew/pkg/configs"
)

const Extension string = ".go"
const modelScaffold string = `package models

type {{ .DomainName | Title }} struct {

}

type {{ .DomainName | Title }}Service interface{
	Create()(any,error)
	Get()(any,error)
	GetAll()(any,error)
	Update()(any,error)
	Delete()(any,error)
}`

const queryScaffold string = `package queries

import "{{ .AppName }}/app/models"

type {{ .DomainName | Title }} models.{{ .DomainName | Title }}

func (q *{{ .DomainName | Title }})Create()(any,error){
    // YOUR CODE GOES HERE
    return nil,nil
}

func (q *{{ .DomainName | Title }})Get()(any,error){
    // YOUR CODE GOES HERE
    return nil,nil
}

func (q *{{ .DomainName | Title }})GetAll()(any,error){
    // YOUR CODE GOES HERE
    return nil,nil
}

func (q *{{ .DomainName | Title }})Update()(any,error){
    // YOUR CODE GOES HERE
    return nil,nil
}

func (q *{{ .DomainName | Title }})Delete()(any,error){
    // YOUR CODE GOES HERE
    return nil,nil
}`

const controllerScaffold string = `package controllers

import (
	"{{ .AppName | ToLower }}/app/queries"

	"github.com//fiber/v2"
)

var {{ .DomainName | ToLower }} *queries.{{ .DomainName | Title }}

func Create{{ .DomainName | Title }}(c *fiber.Ctx)error{
    //YOUR CODE GOES HERE
    return nil

}

func Get{{ .DomainName | Title  }}(c *fiber.Ctx)error{
    //YOUR CODE GOES HERE
    return nil

}

func GetAll{{ .DomainName | Title  }}(c *fiber.Ctx)error{
    //YOUR CODE GOES HERE
    return nil

}

func Update{{ .DomainName | Title }}(c *fiber.Ctx)error{
    //YOUR CODE GOES HERE
    return nil
}

func Delete{{ .DomainName | Title }}(c *fiber.Ctx)error{
    //YOUR CODE GOES HERE
    return nil
}`

const routeScaffold string = `package routes
import (
	"{{ .AppName | ToLower }}/app/controllers"

	"github.com//fiber/v2"
)

func {{ .DomainName | Title }}Routes(a *fiber.App){
    route:=a.Group("/")
    route.Post("/{{ .DomainName | ToLower }}",controllers.Create{{ .DomainName }})
    route.Get("/{{ .DomainName | ToLower }}/:id",controllers.Get{{ .DomainName }})
    route.Get("/{{ .DomainName | ToLower }}",controllers.GetAll{{ .DomainName }})
    route.Put("/{{ .DomainName | ToLower }}/:id",controllers.Update{{ .DomainName }})
    route.Delete("/{{ .DomainName | ToLower }}",controllers.Delete{{ .DomainName }})
}`

const mainScaffold string = `package main

import (
	"{{ .AppName | ToLower }}/app/routes"
    "{{ .AppName | ToLower }}/pkg/utils"
	"os"

	"github.com//fiber/v2"
)

func health(c *fiber.Ctx)error{
    return c.JSON("Health Check from {{ .AppName | ToLower }}")

}


func main(){
    os.Setenv("SERVER_PORT","{{ .AppPort }}")
   app := fiber.New() 
   app.Get("/{{ .AppName | ToLower }}/health", health)
   routes.{{.AppName | Title }}Routes(app)
   // Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "prod" {
		utils.StartServerWithGracefulShutdown(app)
	} else {
		utils.StartServer(app)
	}
}`

var Scaffold = []configs.ScaffoldLocations{
	{
		Source:      controllerScaffold,
		Destination: filepath.Join("app", "controllers"),
		Filename:    "",
	},
	{
		Source:      routeScaffold,
		Destination: filepath.Join("app", "routes"),
		Filename:    "",
	},
	{
		Source:      modelScaffold,
		Destination: filepath.Join("app", "models"),
		Filename:    "",
	},
	{
		Source:      queryScaffold,
		Destination: filepath.Join("app", "queries"),
		Filename:    "",
	},
	{
		Source:      mainScaffold,
		Destination: "cmd",
		Filename:    "main.go",
	},
}
