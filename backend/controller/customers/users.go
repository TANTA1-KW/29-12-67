package customers


import (

    "net/http"

    "github.com/gin-gonic/gin"
 
    "project-se67/config"
 
    "project-se67/entity"
 

)


func GetAll(c *gin.Context) {

   var users []entity.Customers

   db := config.DB()
   results := db.Preload("Gender").Find(&users)
   if results.Error != nil {
       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
       return
   }

   c.JSON(http.StatusOK, users)
}


func Get(c *gin.Context) {

   ID := c.Param("id")

   var user entity.Customers

   db := config.DB()
   results := db.Preload("Gender").First(&user, ID)
   if results.Error != nil {
       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
       return
   }
   if user.ID == 0 {
       c.JSON(http.StatusNoContent, gin.H{})
       return
   }

   c.JSON(http.StatusOK, user)
}


func Update(c *gin.Context) {

   var user entity.Customers

   UserID := c.Param("id")
   db := config.DB()
   result := db.First(&user, UserID)
   if result.Error != nil {
       c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
       return
   }

   if err := c.ShouldBindJSON(&user); err != nil {
       c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
       return
   }


   result = db.Save(&user)
   if result.Error != nil {
       c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
       return
   }

   c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
}


func Delete(c *gin.Context) {


   id := c.Param("id")
   db := config.DB()
   if tx := db.Exec("DELETE FROM customers WHERE id = ?", id); tx.RowsAffected == 0 {
       c.JSON(http.StatusBadRequest, gin.H{"error": "ERROR: ID NOT FOUND"})
       return
   }

   c.JSON(http.StatusOK, gin.H{"message": "Deleted Account Successful"})
}