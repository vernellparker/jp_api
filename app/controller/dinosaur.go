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

type DinosaurController struct {
	dinosaurService     *service.DinosaurService
	verificationService *service.VerificationService
}

func NewDinosaurController(dinosaurService *service.DinosaurService, verificationService *service.VerificationService) *DinosaurController {
	return &DinosaurController{
		dinosaurService,
		verificationService,
	}
}

func (c *DinosaurController) CreateDinosaur(ctx *gin.Context) {
	dino := model.Dinosaur{}

	if err := ctx.ShouldBindJSON(&dino); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.dinosaurService.CreateDinosaur(dino); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{"data": "dino created"})
}
func (c *DinosaurController) PreloadDinosaurs(ctx *gin.Context) {
	err := c.dinosaurService.PreloadTestData()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("Error: Test data could not be loaded, %s", err)})
	}
	ctx.JSON(200, gin.H{"data": "Test data loaded"})
}

func (c *DinosaurController) GetAllDinosaurs(ctx *gin.Context) {
	dinos, err := c.dinosaurService.GetAllDinosaurs()
	if err != nil {
		// Todo fix meg
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("Error: Test data could not be loaded, %s", err)})

	}
	ctx.JSON(200, gin.H{"data": dinos})
}

func (c *DinosaurController) GetOneDinosaur(ctx *gin.Context) {
	paramID := ctx.Param("id")

	dinoId, err := strconv.Atoi(paramID)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	dinosaur, err := c.dinosaurService.GetOneDinosaur(dinoId)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(200, gin.H{"data": dinosaur})
}

func (c *DinosaurController) UpdateDinosaur(ctx *gin.Context) {
	paramID := ctx.Param("id")

	dinoId, err := strconv.Atoi(paramID)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	dino := model.Dinosaur{}

	if err := ctx.ShouldBindJSON(&dino); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	dino.DinoID = dinoId
	if ok, msg, err := c.verificationService.DinoMoveUpdateChecks(dino); !ok {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		ctx.JSON(200, gin.H{"message": msg})
		return
	}

	if _, err := c.dinosaurService.UpdateDinosaur(dino); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{"data": fmt.Sprintf("Dinosaur %d has been successfully updated", dinoId)})
}
