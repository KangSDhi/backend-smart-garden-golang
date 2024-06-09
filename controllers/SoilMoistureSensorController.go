package controllers

import (
	"backend-smart-garden-golang/dto"
	"backend-smart-garden-golang/entity"
	"backend-smart-garden-golang/repository"
	"backend-smart-garden-golang/utils"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateDataSoilMoistureSensor(ctx *fiber.Ctx) error {
	input := new(dto.RequestSoilMoistureSensor)

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

	soilMoistureSensorInput := entity.SoilMoitureSensor{
		NilaiSensorAnalog: input.NilaiSensorAnalog,
		LabelSensorAnalog: input.LabelSensorAnalog,
		TanggalSensor:     utils.StringDatetimeToTime(input.TanggalSensor),
	}

	soilMoistureSensorOutput, err := repository.SaveDataSoilMoistureSensor(soilMoistureSensorInput)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"errors": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"message": "Data Soil Moisture Sensor Created",
		"data": fiber.Map{
			"nilai_analog_sensor": soilMoistureSensorOutput.NilaiSensorAnalog,
			"label_analog_sensor": soilMoistureSensorOutput.LabelSensorAnalog,
			"tanggal_sensor":      soilMoistureSensorOutput.TanggalSensor,
		},
	})

}
