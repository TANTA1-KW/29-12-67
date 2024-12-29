package roles


import (
   "net/http"
   "project-se67/config"
   "project-se67/entity"
   "github.com/gin-gonic/gin"
)


func GetAll(c *gin.Context) {
   db := config.DB()
   var roles []entity.Roles

   db.Find(&roles)
   c.JSON(http.StatusOK, &roles)
}