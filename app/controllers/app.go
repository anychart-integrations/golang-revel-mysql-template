package controllers

import (
	"github.com/anychart-integrations/golang-revel-mysql-template/app/models"
	"github.com/revel/revel"
	"fmt"
	"encoding/json"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	// get data
	fruits, err := Dbm.Select(models.Fruit{},
		`SELECT name, value FROM fruits ORDER BY value DESC limit 5`)
	if err != nil {
		panic(err)
	}
	fmt.Printf("There are %d fruits.\n", len(fruits))

	// convert data to json
	b, err := json.Marshal(fruits)
	if err != nil {
		panic(err)
	}

	// pass data into template
	label := "Anychart Golang Revel template"
	c.ViewArgs["chartData"] = string(b)
	c.ViewArgs["chartTitle"] = "Top 5 fruits"
	return c.Render(label)
}
