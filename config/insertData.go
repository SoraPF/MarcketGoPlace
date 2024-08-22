package config

import (
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
	mps := []objets.Objects{
		{
			IdVendeur:  8,
			Title:      "laisse rouge",
			Price:      4000,
			Desc:       "laisse rouge pour chien 2 mètre",
			StatusID:   2,
			CategoryID: 3,
			Tags: []objets.Tags{
				{ID: 10},
			},
			Img: pq.StringArray([]string{"../public/img/product/Screenshot 2024-08-20 at 14-44-18.png"}),
		},
		{
			IdVendeur:  8,
			Title:      "laisse jaune",
			Price:      4000,
			Desc:       "laisse jaune pour chien 3 mètre, auto retractable",
			StatusID:   2,
			CategoryID: 3,
			Tags: []objets.Tags{
				{ID: 10},
			},
			Img: pq.StringArray([]string{"../public/img/product/Screenshot 2024-08-20 at 14-44-38.png"}),
		},
		{
			IdVendeur:  8,
			Title:      "fauteuille",
			Price:      4000,
			Desc:       "fauteuille noir confortable",
			StatusID:   2,
			CategoryID: 4,
			Img:        pq.StringArray([]string{"../public/img/product/Screenshot 2024-08-20 at 14-48-46 fauteuillr.png"}),
		},
		{
			IdVendeur:  9,
			Title:      "fauteuille",
			Price:      4000,
			Desc:       "fauteuille confortable d'interieur",
			StatusID:   2,
			CategoryID: 4,
			Img:        pq.StringArray([]string{"../public/img/product/Screenshot 2024-08-20 at 14-49-11 fauteuillr.png"}),
		},
		{
			IdVendeur:  9,
			Title:      "Samsung Ecran PC Gaming ",
			Price:      4000,
			Desc:       "C33GC 24\" 100Hz - 4ms, Dalle IPS, FHD (1920 x 1080),1000:1,250 cd/㎡,Eye Saver Mode, FreeSync, Inclinable, HDMI, DisplayPort",
			StatusID:   2,
			CategoryID: 6,
			Tags: []objets.Tags{
				{ID: 1},
			},
			Img: pq.StringArray([]string{"../public/img/product/Screenshot 2024-08-20 at 14-50-01 ecran pc gamer.png"}),
		},
		{
			IdVendeur:  9,
			Title:      "12 lampes solaires",
			Price:      9000,
			Desc:       "12 lampes solaires, waterproof, résistent à la pluie et au vent fort",
			StatusID:   2,
			CategoryID: 5,
			Img:        pq.StringArray([]string{"../public/img/product/Screenshot 2024-08-20 at 14-50-54 aménagement exterieur.png"}),
		},
		{
			IdVendeur:  8,
			Title:      "6 chaise d exterieur",
			Price:      9000,
			Desc:       "6 chaises d extérieur très confort pour une journée ensoleillée",
			StatusID:   2,
			CategoryID: 5,
			Img:        pq.StringArray([]string{"../public/img/product/Screenshot 2024-08-20 at 14-50-54 aménagement exterieur.png"}),
		},
		{
			IdVendeur:  8,
			Title:      "6 chaise d exterieur",
			Price:      9000,
			Desc:       "6 chaises d extérieur très confort pour une journée ensoleillée",
			StatusID:   2,
			CategoryID: 5,
			Img:        pq.StringArray([]string{"../public/img/product/Screenshot 2024-08-20 at 14-51-26 aménagement exterieur.png"}),
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
			Img: pq.StringArray([]string{"../public/img/product/beige-pants.png"}),
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
			Img: pq.StringArray([]string{"../public/img/product/black-pants.png"}),
		},
	}

	for _, obj := range mps {
		result := db.Create(&obj)
		if result.Error != nil {
			panic(result.Error)
		}
	}
}
