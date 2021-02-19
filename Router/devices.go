package Router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kriangkrai/Deivces/Controller"
	"github.com/kriangkrai/Deivces/Models"
)

func Get(c *fiber.Ctx) error {
	device := c.Params("device")
	data := Controller.ReadDoc(device)
	return c.JSON(data)
}

func Gets(c *fiber.Ctx) error {
	data := Controller.ReadDocAll()
	return c.JSON(data)
}

func Insert(c *fiber.Ctx) error {
	var input Models.DataModel
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	m, err := Controller.InsertDoc(input)
	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(m)
}
func Update(c *fiber.Ctx) error {
	var input Models.DataModel

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	m, err := Controller.UpdateDoc(input)
	if err != nil {
		return c.JSON(err.Error())
	}
	return c.JSON(m)
}

func Delete(c *fiber.Ctx) error {
	device := c.Params("device")
	m, err := Controller.DeleteDoc(device)
	if err != nil {
		return c.JSON(err.Error())
	}
	return c.JSON(m)
}
