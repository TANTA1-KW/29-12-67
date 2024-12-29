package status


import (
   "net/http"
   "project-se67/config"
   "project-se67/entity"
   "github.com/gin-gonic/gin"
)


func GetAll(c *gin.Context) {
   db := config.DB()
   var stats []entity.Stats

   db.Find(&stats)
   c.JSON(http.StatusOK, &stats)
}