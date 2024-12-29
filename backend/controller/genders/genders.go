package genders


import (

   "net/http"


   "project-se67/config"

   "project-se67/entity"

   "github.com/gin-gonic/gin"

)


func GetAll(c *gin.Context) {
   db := config.DB()
   var genders []entity.Genders

   db.Find(&genders)
   c.JSON(http.StatusOK, &genders)


}