package controllers

import (
	"backend-smart-garden-golang/dto"
	"backend-smart-garden-golang/entity"
	"backend-smart-garden-golang/repository"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"time"
)

func CreateDataGarden(ctx *fiber.Ctx) error {
	input := new(dto.RequestGarden)

	if err := ctx.BodyParser(input); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"code":   fiber.StatusUnprocessableEntity,
			"errors": err.Error(),
		})
	}

	validate := validator.New()

	if err := validate.Struct(input); err != nil {
		var errorsMap map[string]interface{}
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			errorsMap = make(map[string]interface{})
			for _, fieldError := range validationErrors {
				fieldName := fieldError.Field()
				tagName := fieldError.Tag()
				if fieldName != "" {
					switch tagName {
					case "required":
						errorsMap[fieldName] = map[string]string{"error": fmt.Sprintf("%s Mohon Diisi!", fieldName)}
					case "datetime":
						errorsMap[fieldName] = map[string]string{"error": fmt.Sprintf("%s Format Tanggal Tidak Sesuai!", fieldName)}
					}
				}
			}
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"errors": errorsMap,
		})
	}

	gardenInput := entity.Garden{
		NamaNode:    input.NamaNode,
		Kelembapan:  input.Kelembapan,
		TanggalNode: tanggalNodeParser(input.TanggalNode),
	}

	gardenOutput, err := repository.SaveDataGarden(gardenInput)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"errors": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Ok zone",
		"data": fiber.Map{
			"nama_node":  gardenOutput.NamaNode,
			"kelembapan": fmt.Sprintf("%v%s", gardenOutput.Kelembapan, "%"),
			"tanggal":    gardenOutput.TanggalNode,
		},
	})
}

func GetLastDataGarden(ctx *fiber.Ctx) error {
	garden, err := repository.GetLastRecordGarden()

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":   fiber.StatusNotFound,
			"errors": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Ok Zone",
		"data": fiber.Map{
			"nama_node":  garden.NamaNode,
			"kelembapan": garden.Kelembapan,
			"tanggal":    garden.TanggalNode,
		},
	})
}

func tanggalNodeParser(tanggalNode string) time.Time {
	layoutTime := "2006-01-02 15:04:05"

	loc, _ := time.LoadLocation("Asia/Jakarta")

	t, err := time.ParseInLocation(layoutTime, tanggalNode, loc)

	if err != nil {
		fmt.Println("Error parsing time:", err)
	}

	return t
}
