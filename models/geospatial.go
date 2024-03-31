package models

import "gorm.io/gorm"

// GeoData model
type GeoData struct {
	gorm.Model
	Title   string
	Content string
}

// CreateGeoData creates a new GeoData entry
func CreateGeoData(geoData *GeoData) error {
	return db.Create(geoData).Error
}

// GetGeoData retrieves a GeoData entry by ID
func GetGeoData(id uint) (*GeoData, error) {
	geoData := &GeoData{}
	err := db.First(geoData, id).Error
	if err != nil {
		return nil, err
	}
	return geoData, nil
}

// UpdateGeoData updates a GeoData entry
func UpdateGeoData(geoData *GeoData) error {
	return db.Save(geoData).Error
}

// DeleteGeoData deletes a GeoData entry
func DeleteGeoData(id uint) error {
	return db.Delete(&GeoData{}, id).Error
}
