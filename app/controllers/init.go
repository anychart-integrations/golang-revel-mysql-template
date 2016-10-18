package controllers

import "github.com/revel/revel"

func init() {
	revel.OnAppStart(InitDB)
}
