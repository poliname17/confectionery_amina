package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/poliname17/confectionery/internal/database"
	"github.com/poliname17/confectionery/internal/handlers"
	"github.com/poliname17/confectionery/internal/models"
	"gorm.io/gorm"
)

// docker-compose down -v 
// docker-compose up -d --build

func main() {
	dsn := "host=db user=postgres password=postgres dbname=confectionery port=5432 sslmode=disable"
	db := database.ConnectDatabase(dsn)
	db.AutoMigrate(&models.Chef{}, &models.News{}, &models.Dessert{}, &models.Review{})
	seedData(db)

	router := gin.Default()
	router.LoadHTMLGlob("web/*")

	chefHandler := handlers.ChefHandler{DB: db}
	newsHandler := handlers.NewsHandler{DB: db}
	dessertHandler := handlers.DessertHandler{DB: db}
	reviewHandler := handlers.ReviewHandler{DB: db}

	// --- Главная страница ---
	router.GET("/", func(c *gin.Context) {
		chef := chefHandler.GetChefData()
		news := newsHandler.GetAllNewsData()
		desserts := dessertHandler.GetAllDessertsData()
		reviews := reviewHandler.GetAllReviewsData()

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Chef":     chef,
			"News":     news,
			"Desserts": desserts,
			"Reviews":  reviews,
		})
	})

	// --- API маршруты (по желанию) ---
	router.GET("/news", newsHandler.GetAllNews)
	router.GET("/catalog", dessertHandler.GetAllDesserts)
	router.GET("/reviews", reviewHandler.GetAllReviews)

	log.Println("Server running on port 8080")
	router.Run(":8080")
}
func seedData(db *gorm.DB) {
	var count int64
	db.Model(&models.Chef{}).Count(&count)
	if count == 0 {
		db.Create(&models.Chef{
			Name:        "Амина Кадирова",
			Description: "Авторские торты и десерты на заказ в Казани",
			PhotoURL:    "https://iili.io/KD35orX.jpg",
			Contacts:    "+79033416353\n @ms_aminaaa",
		})
	}
	db.Model(&models.News{}).Count(&count)
	if count == 0 {
		db.Create(&models.News{
			Title:   "Осеннее меню!",
			Content: "В меню появились новые десерты с тыквой и корицей 🍂",
			Date:    "2025-10-01",
		})
	}
	db.Model(&models.Dessert{}).Count(&count)
	if count == 0 {
		db.Create(&models.Dessert{
			Name:        "Торт 'Ягода-ваниль'",
			Description: "ванильный бисквит, молочная пропитка, ягодная начинка, сливочный крем",
			Details:     "",
			Price:       "2200₽/кг",
			ImageURL:    "",
		})
		db.Create(&models.Dessert{
			Name:        "Торт 'Морковный'",
			Description: "морковный бисквит, сливочный крем, карамель, вишневая начинка",
			Details:     "",
			Price:       "2200₽/кг",
			ImageURL:    "",
		})
		db.Create(&models.Dessert{
			Name:        "Торт 'Тропики'",
			Description: "ванильный бисквит, молочная пропитка, начинка 'манго-маракуйя', сливочный крем",
			Details:     "",
			Price:       "2200₽/кг",
			ImageURL:    "",
		})
		db.Create(&models.Dessert{
			Name:        "Торт 'Вишня-шоколад'",
			Description: "шоколадный бисквит, вишневая пропитка, сливочный крем, вишневая начинка",
			Details:     "",
			Price:       "2200₽/кг",
			ImageURL:    "",
		})
		db.Create(&models.Dessert{
			Name:        "Торт 'Сникерс'",
			Description: "шоколадный бисквит, молочная пропитка, сливочный крем, крем из молочного шоколада, карамель, арахис",
			Details:     "",
			Price:       "2200₽/кг",
			ImageURL:    "",
		})
		db.Create(&models.Dessert{
			Name:        "Бенто торты",
			Description: "Начинки: 'Ягода-ваниль', 'Шоколад-вишня', 'Тропики' , 'Сникерс'",
			Details:     "вес: 350 гр/n в набор входит свечка, вилочка, открытка и подарочная упаковка/n *сахарная печать оплачивается отдельно",
			Price:       "1300₽/шт",
			ImageURL:    "https://iili.io/KDB332R.png",
		})
		db.Create(&models.Dessert{
			Name:        "Капкейки ",
			Description: "Начинки: 'Ягода-ваниль', 'Шоколад-вишня', 'Сникерс'",
			Details:     "заказ от 6 шт одного вкуса",
			Price:       "280₽/шт",
			ImageURL:    "https://iili.io/KDqyNQj.png",
		})
		db.Create(&models.Review{
			Message:  "Торт очень вкусный, сытный, сочный. благодарю, чувствуется что сделан с душой и талантом. Благодарю",
			ImageURL: "https://iili.io/KDk1roQ.png",
		})
		db.Create(&models.Review{
			Message:  "Торт на день рождения был просто восхитительным! Он не только невероятно красивый, но и безумно вкусный. Коржи такие нежные и пропитанные, крем воздушный и в меру сладкий. Именинник и гости были в восторге! Спасибо вам огромное за то, что сделали наш праздник еще более особенным!",
			ImageURL: "https://iili.io/KDkNNhF.jpg",
		})
		db.Create(&models.Review{
			Message:  "Торт был вкусный и красивый. Спасибо большое. Сын доволен",
			ImageURL: "https://iili.io/KDkWeun.jpg",
		})
		db.Create(&models.Review{
			Message:  "Спасибо за торт! Вкусный как всегда очень",
			ImageURL: "https://iili.io/KDkUgP1.png",
		})
		db.Create(&models.Review{
			Message:  "Амина торт бесподобный! Спасибо))",
			ImageURL: "https://iili.io/KDkeqVn.pnghttps://iili.io/KDkUgP1.png",
		})
	}
}
