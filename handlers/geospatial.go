package handlers

import (
	"net/http"
	"strconv"

	"geo-data-backend/models"

	"github.com/gin-gonic/gin"
)

// CreateGeoData creates a new GeoData entry
func CreateGeoData(c *gin.Context) {
	var geoData models.GeoData
	if err := c.ShouldBindJSON(&geoData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.CreateGeoData(&geoData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create geospatial data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "GeoData created successfully", "geoData": geoData})
}

// GetGeoData retrieves a GeoData entry by ID
func GetGeoData(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	geoData, err := models.GetGeoData(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "GeoData not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"geoData": geoData})
}

// UpdateGeoData updates a GeoData entry
func UpdateGeoData(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var geoData models.GeoData
	if err := c.ShouldBindJSON(&geoData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	geoData.ID = uint(id)
	err = models.UpdateGeoData(&geoData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update geospatial data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "GeoData updated successfully", "geoData": geoData})
}

// DeleteGeoData deletes a GeoData entry
func DeleteGeoData(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = models.DeleteGeoData(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete geospatial data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "GeoData deleted successfully", "id": id})
}
