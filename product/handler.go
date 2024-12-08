package product

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pewe21/newbelajar/dto"
)

type Handler interface {
	Create(ctx *fiber.Ctx) error
	Get(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{service: service}
}

func (h handler) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()
	var req dto.CreateProductDTO
	err := ctx.BodyParser(&req)

	if err != nil {
		return errors.New("Error parsing request body")
	}

	err = h.service.Create(c, req)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Product created",
	})

}

func (h handler) Get(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	products, err := h.service.Get(c)

	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"data": products,
	})
}

func (h handler) GetById(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	id := ctx.Params("id") //string
	// 127.0.0.1:3000/:id => id (string)

	ids, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": errors.New("cannot convert id to int"),
		})
	}
	product, errs := h.service.GetById(c, ids)

	if errs != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": errs,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"data": product,
	})
}

func (h handler) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	id := ctx.Params("id")

	data := dto.UpdateProductDTO{}

	errs := ctx.BodyParser(&data)

	fmt.Println(data)
	if errs != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": errs,
		})
	}

	ids, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": errors.New("cannot convert id to int"),
		})
	}

	err = h.service.Update(c, ids, data)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Product updated",
	})
}

func (h handler) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	id := ctx.Params("id")

	ids, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": errors.New("cannot convert id to int"),
		})
	}

	err = h.service.Delete(c, ids)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Product deleted",
	})
}
