package seeder

import (
	"log"
	"main/internal/models"
	"time"

	"gorm.io/gorm"
)

func userSeeder(db *gorm.DB) {
	// now := time.Now()
	var users = []models.User{
		{
			Fullname: "Mikhael Kristian",
			Email:    "kristianmikhael667@gmail.com",
			Password: "kijang123",
			RoleID: 1,
			DivisionID: 1,
			Common: models.Common{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
		{
			Fullname: "Mamaria Berigetan",
			Email:    "mamariaberigetan@gmail.com",
			Password: "monitor123",
			RoleID: 2,
			DivisionID: 1,
			Common: models.Common{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
		{
			Fullname: "Mayones Keju",
			Email:    "mayoneskeju@gmail.com",
			Password: "lambo123",
			RoleID: 2,
			DivisionID: 2,
			Common: models.Common{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	}

	if err := db.Create(&users).Error; err != nil {
		log.Printf("Cannot seed data users, with error %v\n", err)
	}
	log.Println("Success seed data users")
}
