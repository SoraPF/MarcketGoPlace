package config

import (
	"Marcketplace/model/entities"
	"Marcketplace/model/objets"
	"fmt"

	"github.com/lib/pq"
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
	/*
		mps := {
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

		db.Create(&mps)
	*/
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

func AutoIncrement(db *gorm.DB) {
	users := []entities.User{
		{
			Username: "azer@email.com",
			Email:    "azer",
			Password: "azer",
		},
		{
			Username: "qsdf@email.com",
			Email:    "qsdf",
			Password: "qsdf",
		},
		{
			Username: "test@email.com",
			Email:    "test",
			Password: "test",
		},
	}
	for _, obj := range users {
		result := db.Create(&obj)
		if result.Error != nil {
			panic(result.Error)
		}
	}

	mps := []objets.Objects{
		{
			IdVendeur:  1,
			Title:      "tshite-blanc",
			Price:      2000,
			Desc:       "tshite-blanc peu utiliser taille XL",
			StatusID:   2,
			CategoryID: 1,
			Tags: []objets.Tags{
				{ID: 4},
			},
			Img: pq.StringArray([]string{"../public/img/product/tshite-blanc.jpg"}),
		},
		{
			IdVendeur:  2,
			Title:      "tshite rose supreme",
			Price:      5000,
			Desc:       "tshite rose supreme non utiliser taille L",
			StatusID:   2,
			CategoryID: 1,
			Tags: []objets.Tags{
				{ID: 4},
			},
			Img: pq.StringArray([]string{"../public/img/product/superme-shirt.jpg"}),
		},
		{
			IdVendeur:  5,
			Title:      "drop the label shirt",
			Price:      3500,
			Desc:       "shirt blanc de drop the label très peu utiliser taille L",
			StatusID:   2,
			CategoryID: 1,
			Tags: []objets.Tags{
				{ID: 4},
			},
			Img: pq.StringArray([]string{"../public/img/product/drop-the-label-shirt.jpg"}),
		},
		{
			IdVendeur:  5,
			Title:      "pantalon de nuit beige",
			Price:      4000,
			Desc:       "pontalon de nuit pour femme beige taille L",
			StatusID:   2,
			CategoryID: 1,
			Tags: []objets.Tags{
				{ID: 5},
			},
			Img: pq.StringArray([]string{"../public/img/product/beige-pants.jpg"}),
		},
		{
			IdVendeur:  5,
			Title:      "pantalon de nuit noir",
			Price:      4000,
			Desc:       "pontalon de nuit pour femme noir taille XL",
			StatusID:   2,
			CategoryID: 1,
			Tags: []objets.Tags{
				{ID: 5},
			},
			Img: pq.StringArray([]string{"../public/img/product/black-pants.jpg"}),
		},
		{
			IdVendeur:  3,
			Title:      "short de plage bleu",
			Price:      3500,
			Desc:       "short de plage bleu pour femme taille M",
			StatusID:   2,
			CategoryID: 1,
			Tags: []objets.Tags{
				{ID: 5},
			},
			Img: pq.StringArray([]string{"../public/img/product/tamara-bellis-BGmVdP6thkU-unsplash.jpg"}),
		},
		{
			IdVendeur:  3,
			Title:      "mini-short jean",
			Price:      3000,
			Desc:       "mini-short jean femme taille M",
			StatusID:   2,
			CategoryID: 1,
			Tags: []objets.Tags{
				{ID: 5},
			},
			Img: pq.StringArray([]string{"../public/img/product/engin-akyurt-Hd4nlxLgIbA-unsplash.jpg"}),
		},
		{
			IdVendeur:  5,
			Title:      "mini-short jean",
			Price:      4000,
			Desc:       "mini-short jean femme taille M",
			StatusID:   2,
			CategoryID: 1,
			Tags: []objets.Tags{
				{ID: 5},
			},
			Img: pq.StringArray([]string{"../public/img/product/engin-akyurt-Hd4nlxLgIbA-unsplash.jpg"}),
		},
	}

	for _, obj := range mps {
		result := db.Create(&obj)
		if result.Error != nil {
			panic(result.Error)
		}
	}
}
