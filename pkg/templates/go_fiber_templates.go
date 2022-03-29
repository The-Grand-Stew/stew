package templates

const GoFiberModelTemplate string = `package models

type {{ . }} struct {

}

type {{ . }}Service interface{
	Create()(any,error)
	Get()(any,error)
	GetAll()(any,error)
	Update()(any,error)
	Delete()(any,error)
}`

const GoFiberQueryTemplate string = `package queries

type {{ . }} models.{{ . }}

func (q *{{ . }})Create()(any,error){
    // YOUR CODE GOES HERE
    return nil,nil
}

func (q *{{ . }})Get()(any,error){
    // YOUR CODE GOES HERE
    return nil,nil
}

func (q *{{ . }})GetAll()(any,error){
    // YOUR CODE GOES HERE
    return nil,nil
}

func (q *{{ . }})Update()(any,error){
    // YOUR CODE GOES HERE
    return nil,nil
}

func (q *{{ . }})Delete()(any,error){
    // YOUR CODE GOES HERE
    return nil,nil
}`

const GoFiberControllerTemplate string = `package controllers

func Create{{ . }}(c *fiber.Ctx)error{
    //YOUR CODE GOES HERE

}

func Get{{ . }}(c *fiber.Ctx)error{
    //YOUR CODE GOES HERE

}

func GetAll{{ . }}c *fiber.Ctx)error{
    //YOUR CODE GOES HERE

}

func Update{{ . }}c *fiber.Ctx)error{
    //YOUR CODE GOES HERE

}

func Delete{{ . }}c *fiber.Ctx)error{
    //YOUR CODE GOES HERE

}`

const GoFiberRouteTemplate string = `package routes

func {{ . }}Routes(a *fiber.App){
    route:=a.Group("")
    route.Post("/{{ . | ToLower }}",controllers.Create{{ . }})
    route.Get("/{{ . }}/:id",controllers.Get{{ . }})
    route.Get("/{{ . }}",controllers.GetAll{{ . }})
    route.Put("/{{ . }}",controllers.Update{{ . }})
    route.Delete("/{{ . }}",controllers.Delete{{ . }})
}`

const GoFiberMainTemplate string = `package main


func main(){
   app := fiber.New(config) 
   routes.{{ . }}Routes(app)
   // Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "prod" {
		utils.StartServerWithGracefulShutdown(app)
	} else {
		utils.StartServer(app)
	}
}`
