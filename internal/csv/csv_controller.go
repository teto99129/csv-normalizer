package csv

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

type CSVController struct {
  Service *CSVService
}

func (controller *CSVController)RegistEndpoint(r *gin.Engine) {
  r.POST("/csv", func(c *gin.Context) {
    file, _, err := c.Request.FormFile("file")
    if err != nil {
      c.Status(http.StatusBadRequest)
    }
    res := controller.Service.CSVToJson(file)
    c.JSON(http.StatusOK, res)
  })

}
