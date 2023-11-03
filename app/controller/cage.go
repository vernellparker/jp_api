package controller

import (
	"JurrassicParkAPI/app/service"
	"JurrassicParkAPI/model"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type CageController struct {
	cageService         *service.CageService
	verificationService *service.VerificationService
}

func NewCageController(cageService *service.CageService, verificationService *service.VerificationService) *CageController {
	return &CageController{
		cageService,
		verificationService,
	}
}

func (c *CageController) CreateCage(ctx *gin.Context) {
	cage := model.IntakeCage{}

	if err := ctx.ShouldBindJSON(&cage); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	//check if user added dinos to the cage
	if len(cage.Dinosaurs) > 0 {
		if ok, msg, err := c.verificationService.CageVerificationChecks(cage, 0); !ok {
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			}

			ctx.JSON(200, gin.H{"message": msg})
			return
		}

	}

	if err := c.cageService.CreateCage(cage); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{"data": "crate created"})
}

func (c *CageController) GetAllCages(ctx *gin.Context) {
	cages, err := c.cageService.GetAllCages()
	if err != nil {
		return
	}
	ctx.JSON(200, gin.H{"data": cages})
}

func (c *CageController) GetOneCage(ctx *gin.Context) {
	paramID := ctx.Param("id")

	cageId, err := strconv.Atoi(paramID)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	cage, err := c.cageService.GetOneCages(cageId)
	if err != nil {
		return
	}

	ctx.JSON(200, gin.H{"data": cage})
}

func (c *CageController) UpdateCage(ctx *gin.Context) {
	paramID := ctx.Param("id")

	cageId, err := strconv.Atoi(paramID)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	cage := model.IntakeCage{}

	if err := ctx.ShouldBindJSON(&cage); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	cage.CageID = cageId

	if ok, msg, err := c.verificationService.CageVerificationChecks(cage, 0); !ok {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		ctx.JSON(200, gin.H{"message": msg})
		return
	}

	if ok, msg, err := c.verificationService.CagePowerChecks(cage); !ok {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		ctx.JSON(200, gin.H{"message": msg})
		return
	}

	id, err := c.cageService.UpdateCage(cage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	if id == 0 {
		ctx.JSON(http.StatusNotImplemented, gin.H{"data": fmt.Sprintf("Cage %d failed to update.", cageId)})
		return
	}

	ctx.JSON(200, gin.H{"data": fmt.Sprintf("Cage %d has been successfully updated", cageId)})
}
