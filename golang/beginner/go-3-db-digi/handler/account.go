package handler

import (
	"go-3-db-digi/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountInterface interface {
	Create(*gin.Context)
	Read(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	List(*gin.Context)

	My(*gin.Context)
}

type accountImplement struct {
	db *gorm.DB
}

func NewAccount(db *gorm.DB) AccountInterface {
	return &accountImplement{
		db: db,
	}
}

func (a *accountImplement) Create(c *gin.Context) {
	payload := model.Account{}

	err := c.BindJSON(&payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	result := a.db.Create(&payload)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Create success",
		"data":    payload,
	})
}

func (a *accountImplement) Read(c *gin.Context) {
	var account model.Account
	id := c.Param("id")

	if err := a.db.First(&account, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Not found",
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": account,
	})
}

func (a *accountImplement) Update(c *gin.Context) {
	payload := model.Account{}

	err := c.BindJSON(&payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	id := c.Param("id")

	account := model.Account{}
	result := a.db.First(&account, "account_id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Not found",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// Update data
	account.Name = payload.Name
	a.db.Save(account)

	c.JSON(http.StatusOK, gin.H{
		"message": "Update success",
	})
}

func (a *accountImplement) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := a.db.Where("account_id = ?", id).Delete(&model.Account{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Not found",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// No data deleted
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete success",
		"data": map[string]string{
			"account_id": id,
		},
	})
}

func (a *accountImplement) List(c *gin.Context) {
	var accounts []model.Account

	if err := a.db.Find(&accounts).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": accounts,
	})
}

func (a *accountImplement) My(c *gin.Context) {
	var account model.Account
	accountID := c.GetInt64("account_id")

	if err := a.db.First(&account, accountID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Not found",
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": account,
	})
}
