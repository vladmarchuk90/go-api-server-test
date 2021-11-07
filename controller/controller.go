package controller

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/vladmarchuk90/go-api-server-test/model"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// Default godoc
// @Summary Show possible API routes
// @Description Root path to give hint about available methods
// @Router / [get]
func Default(c *fiber.Ctx) error {

	response := `There are available the next api methods:
	- GET groups, /groups
	- GET group by id, /groups/{groupId}
	- GET index by id, /indexes/{indexId}
	- GET blocks bu Number, Hash or the latest one, /blocks/{blockNumber|blockHash|}`

	return c.Send([]byte(response))
}

// GetGroups godoc
// @Summary Get groups
// @Description get all groups in smart contract
// @Accept  json
// @Produce json
// @Success 200 {array} big.Int
// @Failure 400,404 {object} ErrorResponse
// @Router /groups [get]
func GetGroups(c *fiber.Ctx) error {

	response, err := model.GetGroups()
	if err != nil {
		c.Status(400)
		err := ErrorResponse{Error: "Could not get group Ids"}
		errJSON, _ := json.Marshal(err)
		return c.Send(errJSON)
	}

	responseJSON, _ := json.Marshal(response)
	return c.Send(responseJSON)
}

// GetGroup godoc
// @Summary Get group's information
// @Description get detail information of specific group
// @Accept  json
// @Produce json
// @Param groupId path int true "Group ID"
// @Success 200 {object} model.Group
// @Failure 400,404 {object} ErrorResponse
// @Router /groups/{groupId} [get]
func GetGroup(c *fiber.Ctx) error {

	groupId := c.Params("groupId")

	response, err := model.GetGroup(groupId)
	if err != nil {
		c.Status(400)
		err := ErrorResponse{Error: "Could not get group by Id"}
		errJSON, _ := json.Marshal(err)
		return c.Send(errJSON)
	}

	responseJSON, _ := json.Marshal(response)
	return c.Send(responseJSON)
}

// GetIndex godoc
// @Summary Get index's information
// @Description get detail information of specific index
// @Accept  json
// @Produce json
// @Param indexId path int true "Index ID"
// @Success 200 {object} model.Index
// @Failure 400,404 {object} ErrorResponse
// @Router /indexes/{indexId} [get]
func GetIndex(c *fiber.Ctx) error {

	indexId := c.Params("indexId")

	response, err := model.GetIndex(indexId)
	if err != nil {
		c.Status(400)
		err := ErrorResponse{Error: "Could not get index by Id"}
		errJSON, _ := json.Marshal(err)
		return c.Send(errJSON)
	}

	responseJSON, _ := json.Marshal(response)
	return c.Send(responseJSON)
}

// GetBlock godoc
// @Summary Get block's header information
// @Description get all header information of the block
// @Accept  json
// @Produce json
// @Param blockParameter path string true "blockNumber/blockHash/latest"
// @Success 200 {object} types.Header
// @Failure 400,404 {object} ErrorResponse
// @Router /blocks/{blockParameter} [get]
func GetBlock(c *fiber.Ctx) error {

	blockParameter := c.Params("blockParameter")

	response, err := model.GetBlock(blockParameter)
	if err != nil {
		c.Status(400)
		err := ErrorResponse{Error: "Could not get block by parameter"}
		errJSON, _ := json.Marshal(err)
		return c.Send(errJSON)
	}

	responseJSON, _ := json.Marshal(response)
	return c.Send(responseJSON)
}
