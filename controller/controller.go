package controller

import (
	"encoding/json"

	"github.com/gofiber/fiber"
	"github.com/vladmarchuk90/go-api-server-test/model"
)

func Default(c *fiber.Ctx) {

	response := `There are available the next api methods:
	- GET groups, /groups
	- GET group by id, /groups/{groupId}
	- GET index by id, /indexes/{indexId}
	- GET blocks bu Number, Hash or the latest one, /blocks/{blockNumber|blockHash|}`

	c.Send(string(response))
}

func GetGroups(c *fiber.Ctx) {

	response, err := json.Marshal(model.GetGroups())
	if err != nil {
		c.Send("Could not get groups Ids")
	}

	c.Send(string(response))
}

func GetGroup(c *fiber.Ctx) {

	groupId := c.Params("groupId")

	response, err := json.Marshal(model.GetGroup(groupId))
	if err != nil {
		c.Send("Could not get group by Id")
	}

	c.Send(string(response))
}

func GetIndex(c *fiber.Ctx) {

	indexId := c.Params("indexId")

	response, err := json.Marshal(model.GetIndex(indexId))
	if err != nil {
		c.Send("Could not get index by Id")
	}

	c.Send(string(response))
}

func GetBlock(c *fiber.Ctx) {

	blockParameter := c.Params("blockParameter")

	response, err := json.Marshal(model.GetBlock(blockParameter))
	if err != nil {
		c.Send("Could not get block by parameter")
	}

	c.Send(string(response))
}
