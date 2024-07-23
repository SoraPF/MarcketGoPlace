package config

import (
	"Marcketplace/model/objets"
	"fmt"
	"os"

	"gorm.io/gorm"
)

func InsertImage(db *gorm.DB, id uint, imagePath string) {
	var category objets.Categories
	result := db.First(&category, id)
	if result.Error != nil {
		fmt.Println(fmt.Errorf("category with ID %d not found", id))
	}
	updateCategory(db, &category, "", imagePath)
}

func updateCategory(db *gorm.DB, category *objets.Categories, newTitle string, newImagePath string) error {
	if newTitle != "" {
		category.Title = newTitle
	}

	if newImagePath != "" {
		imageBytes, err := os.ReadFile(newImagePath)
		if err != nil {
			return fmt.Errorf("failed to read new image file: %v", err)
		}
		category.Image = imageBytes
	} else {
		category.Image = nil
	}

	result := db.Save(category)
	if result.Error != nil {
		return fmt.Errorf("failed to update category: %v", result.Error)
	}

	fmt.Printf("Category ID: %d updated successfully\n", category.ID)
	return nil
}
