package config

import (
	"Marcketplace/model/objets"
	"fmt"

	"gorm.io/gorm"
)

func InsertImages(db *gorm.DB) {
	InsertImage(db, 1, "../public/img/vetement.jpg")
	InsertImage(db, 2, "../public/img/vehicule.png")
	InsertImage(db, 3, "../public/img/animaux.png")
	InsertImage(db, 4, "../public/img/amenagement_interieur.jpg")
	InsertImage(db, 5, "../public/img/amenagement_exterieur.jpg")
	InsertImage(db, 6, "../public/img/high-tech.jpg")
	InsertImage(db, 7, "../public/img/beaute_et_bien_etre.jpg")
	InsertImage(db, 8, "../public/img/jeux.jpg")
	InsertImage(db, 9, "../public/img/loisire_et_sport.jpg")
	InsertImage(db, 10, "../public/img/bureau.png")
}

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
		category.Image = newImagePath
	} else {
		category.Image = ""
	}

	result := db.Save(category)
	if result.Error != nil {
		return fmt.Errorf("failed to update category: %v", result.Error)
	}

	fmt.Printf("Category ID: %d updated successfully\n", category.ID)
	return nil
}

func InsertObject(db *gorm.DB) {
	/*mps := {
		IdVendeur:  1,
		Title:      "manette de PS4",
		Price:      8990,
		Desc:       "manette de PS4, dualshock 4, 1 ans d'âge",
		StatusID:   1,
		CategoryID: 8,
		Tags: []objets.Tags{
			{ID: 2},
			{ID: 3},
		},
	}

	db.Create(&mps)*/
	var obj objets.Objects
	id := uint(1)
	result := db.First(&obj, id)
	if result.Error != nil {
		fmt.Println(fmt.Errorf("object with ID %d not found", id))
		return
	}

	obj.Img = []string{"../public/img/product/m1.jpg", "../public/img/product/m2.jpg"}
	result = db.Save(&obj)
	if result.Error != nil {
		fmt.Println(fmt.Errorf("failed to update object: %v", result.Error))
		return
	}

	fmt.Printf("Object ID: %d updated successfully\n", obj.ID)
}
