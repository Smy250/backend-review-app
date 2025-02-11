package visited

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func GetVisitedCountRepository(userId uint) (visitedCount uint, err error) {
	db := config.DB
	var user schema.User

	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return visitedCount, err
	}

	visitedCount = uint(db.Model(&user).Association("VisitedPlaces").Count())
	return visitedCount, nil
}

func GetVisitorsCountRepository(placeID uint) (visitorsCount uint, err error) {
	db := config.DB
	var count int64
	err = db.Table("place_visitors").Where("place_id = ?", placeID).Count(&count).Error
	return uint(count), err //visitorsCount, err
}

func CreateVisitedPlaceRepository(userID, placeID uint) error {
	db := config.DB

	var user schema.User
	if err := db.First(&user, userID).Error; err != nil {
		return err
	}

	var place schema.Place
	if err := db.First(&place, placeID).Error; err != nil {
		return err
	}

	if err := db.Model(&user).Association("VisitedPlaces").Append(&place); err != nil {
		return err
	}

	return nil
}

func DeleteVisitedPlaceRepository(userID, placeID uint) error {
	db := config.DB

	var user schema.User
	if err := db.First(&user, userID).Error; err != nil {
		return err
	}

	var place schema.Place
	if err := db.First(&place, placeID).Error; err != nil {
		return err
	}

	if err := db.Model(&user).Association("VisitedPlaces").Delete(&place); err != nil {
		return err
	}

	return nil
}
