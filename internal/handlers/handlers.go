package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	MRUApp "github.com/koushilab/MRUAppTest/internal/algorithm"
)

type MRUs interface {
	getMRUs(c *gin.Context)
}

type MRUInfo struct{}

func (mrus MRUInfo) getMRUs(c *gin.Context) {

	data := MRUApp.MakeTestInput()

	result, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	} else {
		c.String(http.StatusOK, fmt.Sprintf(string(result)))
	}
}
