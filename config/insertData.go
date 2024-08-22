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
		{
			Username: "uiop@email.com",
			Email:    "uiop",
			Password: "uiop",
		},
		{
			Username: "jklm@email.com",
			Email:    "jklm",
			Password: "jklm",
		},
		{
			Username: "cvbn@email.com",
			Email:    "cvbn",
			Password: "cvbn",
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
		{
			IdVendeur:  4,
			Title:      "clavier mac",
			Price:      6000,
			Desc:       "clavier pour mac toujours blanc",
			StatusID:   2,
			CategoryID: 6,
			Tags: []objets.Tags{
				{ID: 1},
			},
			Img: pq.StringArray([]string{"../public/img/product/mac-clavier.jpg"}),
		},
		{
			IdVendeur:  6,
			Title:      "clavier sans fil",
			Price:      5000,
			Desc:       "clavier noir, usb sans fil, avec touche modifiable",
			StatusID:   2,
			CategoryID: 6,
			Tags: []objets.Tags{
				{ID: 1},
			},
			Img: pq.StringArray([]string{"../public/img/product/terrillo-walls-clavier.jpg"}),
		},
		{
			IdVendeur:  7,
			Title:      "clavier logitec rgb",
			Price:      4000,
			Desc:       "clavier logitec rgb, filaire, 2ans, méchanique",
			StatusID:   2,
			CategoryID: 6,
			Tags: []objets.Tags{
				{ID: 1},
			},
			Img: pq.StringArray([]string{"../public/img/product/clavier-logitec.jpg"}),
		},
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
			Img: pq.StringArray([]string{"../public/img/product/Screenshot 2024-08-20 at 14-44-18.jpg"}),
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
			Img: pq.StringArray([]string{"../public/img/product/Screenshot 2024-08-20 at 14-44-38.jpg"}),
		},
		{
			IdVendeur:  8,
			Title:      "fauteuille",
			Price:      4000,
			Desc:       "fauteuille noir confortable",
			StatusID:   2,
			CategoryID: 4,
			Img:        pq.StringArray([]string{"../public/img/product/Screenshot 2024-08-20 at 14-48-46 fauteuillr.jpg"}),
		},
		{
			IdVendeur:  9,
			Title:      "fauteuille",
			Price:      4000,
			Desc:       "fauteuille confortable d'interieur",
			StatusID:   2,
			CategoryID: 4,
			Img:        pq.StringArray([]string{"../public/img/product/Screenshot 2024-08-20 at 14-49-11 fauteuillr.jpg"}),
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
			Img: pq.StringArray([]string{"../public/img/product/Screenshot 2024-08-20 at 14-50-01 ecran pc gamer.jpg"}),
		},
		{
			IdVendeur:  9,
			Title:      "12 lampes solaires",
			Price:      9000,
			Desc:       "12 lampes solaires, waterproof, résistent à la pluie et au vent fort",
			StatusID:   2,
			CategoryID: 5,
			Img:        pq.StringArray([]string{"../public/img/product/Screenshot 2024-08-20 at 14-50-54 aménagement exterieur.jpg"}),
		},
		{
			IdVendeur:  8,
			Title:      "6 chaise d exterieur",
			Price:      9000,
			Desc:       "6 chaises d extérieur très confort pour une journée ensoleillée",
			StatusID:   2,
			CategoryID: 5,
			Img:        pq.StringArray([]string{"../public/img/product/Screenshot 2024-08-20 at 14-50-54 aménagement exterieur.jpg"}),
		},
		{
			IdVendeur:  8,
			Title:      "6 chaise d exterieur",
			Price:      9000,
			Desc:       "6 chaises d extérieur très confort pour une journée ensoleillée",
			StatusID:   2,
			CategoryID: 5,
			Img:        pq.StringArray([]string{"../public/img/product/Screenshot 2024-08-20 at 14-51-26 aménagement exterieur.jpg"}),
		},
		{
			IdVendeur:  2,
			Title:      "La Belle et la Bête, Rose Eternelle sous Cloche",
			Price:      12000,
			Desc:       "DIDUDE Rose Eternelle sous Cloche,La Belle et la Bête Rose,avec Lumineuse LED, Dôme en Verre et Base en Bois",
			StatusID:   2,
			CategoryID: 4,
			Img:        pq.StringArray([]string{"../public/img/product/Rose Eternelle sous Cloche.jpg"}),
		},
		{
			IdVendeur:  9,
			Title:      "4 Brillants à Lèvres",
			Price:      2000,
			Desc:       "Lot de 4 Brillants à Lèvres Aromatisés aux Fruits, Transparent Hydratant Lotion Pour les Lèvres",
			StatusID:   2,
			CategoryID: 7,
			Img:        pq.StringArray([]string{"../public/img/product/4 brillants à lèvre.jpg"}),
		},
		{
			IdVendeur:  9,
			Title:      "lampe à lave",
			Price:      2000,
			Desc:       "GIFTMARKET - Lampe à lave bleue. Lampe de chevet avec 2 ampoules incluses. Cadeau amusant pour les adolescents. Lampe rétro de 34 x 8,5 cm.",
			StatusID:   2,
			CategoryID: 10,
			Img:        pq.StringArray([]string{"../public/img/product/lampe à lave.jpg"}),
		},
		{
			IdVendeur:  9,
			Title:      "lampe keli",
			Price:      2000,
			Desc:       "Atmosphera - Lampe Keli dorée - métal H45 cm",
			StatusID:   2,
			CategoryID: 10,
			Img:        pq.StringArray([]string{"../public/img/product/lampe keli.jpg"}),
		},
		{
			IdVendeur:  5,
			Title:      "Souris Sans Fil",
			Price:      2000,
			Desc:       "Logitech M185 Souris Sans Fil, 2.4 GHz avec Mini Récepteur USB, Longévité de la Pile 12 Mois, Résolution du Capteur 1000 PPP, Ambidextre, Compatible PC, Mac, Ordinateur Portable - Gris/Noir",
			StatusID:   2,
			CategoryID: 6,
			Img:        pq.StringArray([]string{"../public/img/product/Souris Sans Fil.jpg"}),
		},
		{
			IdVendeur:  6,
			Title:      "Bureau Table Blanc",
			Price:      4000,
			Desc:       "vidaXL Bureau Table dOrdinateur Table Blanc 90x45x76 cm Aggloméré",
			StatusID:   2,
			CategoryID: 5,
			Img:        pq.StringArray([]string{"../public/img/product/Bureau Table.jpg"}),
		},
		{
			IdVendeur:  6,
			Title:      "bureau blanc",
			Price:      6000,
			Desc:       "Yaheetech Bureau Blanc pour Ordinateur PC Bureau Informatique avec Porte Clavier Coulissant Tiroir et 2 étagères de Rangement Table de Bureau",
			StatusID:   2,
			CategoryID: 10,
			Img:        pq.StringArray([]string{"../public/img/product/bureau blanc.jpg"}),
		},
		{
			IdVendeur:  7,
			Title:      "bureau gamer",
			Price:      10000,
			Desc:       "HLFURNIEU 120 x 60 cm Bureau Gaming, Bureau Gamer Informatique Ergonomique, Table Gaming en Fibre de Carbone, Gaming Desk avec Porte Gobelet et Crochet",
			StatusID:   2,
			CategoryID: 10,
			Img:        pq.StringArray([]string{"../public/img/product/bureau gamer.jpg"}),
		},
		{
			IdVendeur:  7,
			Title:      "t-shirt under armor",
			Price:      8000,
			Desc:       "Under Armour UA HG Armour Comp SS T-Shirt à Manches Courtes, T-Shirt de Compression pour Homme ",
			StatusID:   2,
			CategoryID: 9,
			Tags: []objets.Tags{
				{ID: 4},
			},
			Img: pq.StringArray([]string{"../public/img/product/t-shirt under armor.jpg"}),
		},
		{
			IdVendeur:  1,
			Title:      "sac puma",
			Price:      10000,
			Desc:       "Puma Fundamentals Sports Bag XS Sac De Sport, Puma Black ",
			StatusID:   2,
			CategoryID: 9,
			Img:        pq.StringArray([]string{"../public/img/product/sac puma.jpg"}),
		},
		{
			IdVendeur:  1,
			Title:      "corde a sauté",
			Price:      6500,
			Desc:       "Blukar Corde à Sauter, Speed Jump Rope Réglable pour Hommes & Femmes , Roulements à Billes en Acier, Poignées Antidérapantes en Mousse, Câble en Acier ",
			StatusID:   2,
			CategoryID: 9,
			Img:        pq.StringArray([]string{"../public/img/product/corde a sauté.jpg"}),
		},
		{
			IdVendeur:  2,
			Title:      "carte uno",
			Price:      2500,
			Desc:       "Games UNO Classique, Jeu De Cartes Familial pour Enfants Et Adultes, Jeu De Société pour Soirée Jeux en Famille Ou en Voyage, 2 À 10 Joueurs",
			StatusID:   2,
			CategoryID: 8,
			Img:        pq.StringArray([]string{"../public/img/product/carte uno.jpg"}),
		},
		{
			IdVendeur:  2,
			Title:      "Jeu de Dés-Qwixx",
			Price:      4000,
			Desc:       " Gigamic- Jeu de Dés-Qwixx, 8 ans to 99 ans, JNQX ",
			StatusID:   2,
			CategoryID: 8,
			Img:        pq.StringArray([]string{"../public/img/product/carte uno.jpg"}),
		},
		{
			IdVendeur:  2,
			Title:      "SKYJO, de Magilano ",
			Price:      3000,
			Desc:       "Le Tout Nouveau Jeu de Cartes/de société pour Les Petits et Les Grands pour se divertir et Passer des soirées Amusantes avec des Amis et en Famille",
			StatusID:   2,
			CategoryID: 8,
			Img:        pq.StringArray([]string{"../public/img/product/carte uno.jpg"}),
		},
	}

	for _, obj := range mps {
		result := db.Create(&obj)
		if result.Error != nil {
			panic(result.Error)
		}
	}
}
