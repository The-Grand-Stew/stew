package templates

const GoFiberModelTemplate string = `package models

type {{ . }} struct {

}

type {{ . }}Service interface{
	Create()(any,error)
}`

const GoFiberQueryTemplate string = `package queries`
const GoFiberControllerTemplate string = `package controllers`
const GoFiberRouteTemplate string = `package routes`
const GoFiberMainTemplate string = `package main`
