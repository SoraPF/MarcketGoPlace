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
	mps := objets.Objects{
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

	status_Creation := objets.Statuses{Title: "create"}
	status_Insale := objets.Statuses{Title: "in sale"}
	status_Blocked := objets.Statuses{Title: "blocked"}
	status_Sold := objets.Statuses{Title: "sold"}

	tag_ordinateur := objets.Tags{Title: "Ordinateurs"}
	tag_manettes := objets.Tags{Title: "Manettes"}
	tag_playstation := objets.Tags{Title: "Playstation"}
	tag_tricot := objets.Tags{Title: "Tricots"}
	tag_short := objets.Tags{Title: "Chorts"}
	tag_vase := objets.Tags{Title: "Vases"}
	tag_tele := objets.Tags{Title: "Télévisions"}
	tag_table := objets.Tags{Title: "Tables"}
	tag_croquette := objets.Tags{Title: "Croquettes"}
	tag_laisse := objets.Tags{Title: "Laisses"}
	tag_verni := objets.Tags{Title: "Vernis"}
	tag_balle := objets.Tags{Title: "Balle"}
	tag_plante := objets.Tags{Title: "Plantes"}
	tag_voiture := objets.Tags{Title: "Voitures"}
	tag_scooter := objets.Tags{Title: "Scooter"}

	db.Create(&status_Creation)
	db.Create(&status_Insale)
	db.Create(&status_Blocked)
	db.Create(&status_Sold)

	db.Create(&tag_ordinateur)
	db.Create(&tag_manettes)
	db.Create(&tag_playstation)
	db.Create(&tag_tricot)
	db.Create(&tag_short)
	db.Create(&tag_vase)
	db.Create(&tag_tele)
	db.Create(&tag_table)
	db.Create(&tag_croquette)
	db.Create(&tag_laisse)
	db.Create(&tag_verni)
	db.Create(&tag_balle)
	db.Create(&tag_plante)
	db.Create(&tag_voiture)
	db.Create(&tag_scooter)

	db.Create(&mps)
}
