package seeder

import (
	"log"
	"main/internal/models"

	"gorm.io/gorm"
)

func userSeeder(db *gorm.DB) {
	// now := time.Now()
	var users = []models.User{
		{
			Name:     "Mikhael Kristian",
			Email:    "kristianmikhael667@gmail.com",
			Password: "kijang123",
		},
		{
			Name:     "Mamaria Berigetan",
			Email:    "mamariaberigetan@gmail.com",
			Password: "monitor123",
		},
		{
			Name:     "Mayones Keju",
			Email:    "mayoneskeju@gmail.com",
			Password: "lambo123",
		},
	}

	if err := db.Create(&users).Error; err != nil {
		log.Printf("Cannot seed data users, with error %v\n", err)
	}
	log.Println("Success seed data users")
}
