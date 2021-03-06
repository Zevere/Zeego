package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"zeego/web/models"
)

type GraphQLController struct {
	beego.Controller
	zevereSchema *graphql.Schema
}

func NewGraphQLController(zevereSchema *graphql.Schema) *GraphQLController {
	return &GraphQLController{zevereSchema: zevereSchema}
}

func (c *GraphQLController) Post() {
	var req models.GraphQLRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Ctx.ResponseWriter.Status = 400
		c.Ctx.WriteString("")
	}

	result := graphql.Do(graphql.Params{
		Schema:        *c.zevereSchema,
		RequestString: req.Query,
	})

	c.Data["json"] = result
	c.ServeJSON()
}
