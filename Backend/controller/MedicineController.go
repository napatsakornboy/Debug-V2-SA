package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/panupongKanin/ProjectSA-arm/entity"
)

/////DOCTOR//////

func CreateDoctor(c *gin.Context) {
	var doctor entity.DOCTOR
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&doctor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": doctor})

}

func GetDoctor(c *gin.Context) {
	var GetDoctor entity.DOCTOR
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM doctors WHERE id = ?", id).Scan(&GetDoctor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": GetDoctor})
}

func ListDoctor(c *gin.Context) {

	var doctor []entity.DOCTOR
	if err := entity.DB().Raw("SELECT * FROM doctors").Scan(&doctor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": doctor})

}

//////MEDICINE//////

func CreateMedicine(c *gin.Context) {

	var medicine entity.MEDICINE

	if err := c.ShouldBindJSON(&medicine); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&medicine).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": medicine})

}

func GetMedicine(c *gin.Context) {
	var GetMedicine entity.MEDICINE
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM medicine WHERE id = ?", id).Scan(&GetMedicine).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": GetMedicine})
}

func ListMedicine(c *gin.Context) {

	var medicine []entity.MEDICINE
	if err := entity.DB().Raw("SELECT * FROM medicines").Scan(&medicine).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": medicine})

}

//////WHERE//////

func CreateWhere(c *gin.Context) {

	var where entity.WHERE
	if err := c.ShouldBindJSON(&where); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&where).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": where})

}

func GetWhere(c *gin.Context) {
	var GetWhere entity.WHERE
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM wheres WHERE id = ?", id).Scan(&GetWhere).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": GetWhere})
}

func ListWhere(c *gin.Context) {

	var where []entity.WHERE
	if err := entity.DB().Raw("SELECT * FROM wheres").Scan(&where).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": where})

}

//////BASKET/////

func CreateBasket(c *gin.Context) {

	var BASKET entity.BASKET
	var WHERE entity.WHERE
	var MEDICINE entity.MEDICINE
	var DOCTOR entity.DOCTOR

	if err := c.ShouldBindJSON(&BASKET); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", BASKET.WHERE_ID).First(&WHERE); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Where not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", BASKET.MEDICINE_ID).First(&MEDICINE); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Medicine not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", BASKET.DOCTOR_ID).First(&DOCTOR); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Doctor not found"})
		return
	}

	basket := entity.BASKET{
		Amount:   BASKET.Amount,
		Add_time: BASKET.Add_time,
		MEDICINE: MEDICINE,
		WHERE:    WHERE,
		DOCTOR:   DOCTOR,
	}

	if err := entity.DB().Create(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": basket})

}

func GetBasket(c *gin.Context) {
	var GetBasket entity.BASKET
	id := c.Param("id")
	if err := entity.DB().Preload("DOCTOR").Preload("MEDICINE").Preload("WHERE").Preload("Symtom").Preload("User").Raw("SELECT * FROM baskets WHERE id = ?", id).Find(&GetBasket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": GetBasket})
}

func ListBasket(c *gin.Context) {

	var basket []entity.BASKET
	if err := entity.DB().Preload("Symtom").Raw("SELECT * FROM baskets").Find(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": basket})

}
