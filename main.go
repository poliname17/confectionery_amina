package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/poliname17/confectionery/internal/database"
	"github.com/poliname17/confectionery/internal/handlers"
	"github.com/poliname17/confectionery/internal/models"
	"gorm.io/gorm"
)

// Railway автоматически передаёт URL базы данных в переменной окружения DATABASE_URL.
// PORT — также задаётся автоматически.

func main() {
	// Подключение к базе
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}
	db := database.ConnectDatabase(dsn)
	db.AutoMigrate(&models.Chef{}, &models.News{}, &models.Dessert{}, &models.Review{})
	seedData(db)

	// Настройка роутера
	router := gin.Default()
	router.LoadHTMLGlob("web/*")
	router.Static("/static", "./web") // если есть CSS/JS/изображения

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

	// --- API маршруты ---
	router.GET("/news", newsHandler.GetAllNews)
	router.GET("/catalog", dessertHandler.GetAllDesserts)
	router.GET("/reviews", reviewHandler.GetAllReviews)

	// --- Запуск сервера ---
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server running on port", port)
	router.Run(":" + port)
}

// --- Первичное заполнение базы данными ---
func seedData(db *gorm.DB) {
	var count int64
	db.Model(&models.Chef{}).Count(&count)
	if count == 0 {
		db.Create(&models.Chef{
			Name:        "Амина Кадирова",
			Description: "Авторские торты и десерты на заказ в Казани",
			PhotoURL:    "https://iili.io/KD35orX.jpg",
			Contacts:    "+79033416353\n@ms_aminaaa",
		})
	}

	db.Model(&models.News{}).Count(&count)
	if count == 0 {
		db.Create(&models.News{
			Title:   "У нас появился сайт!",
			Content: "Теперь вы можете найти всю важную информацию здесь",
			Date:    "13.11.2025",
		})
	}

	db.Model(&models.Dessert{}).Count(&count)
	if count == 0 {
		db.Create(&models.Dessert{
			Name:        "Торт 'Ягода-ваниль'",
			Description: "ванильный бисквит, молочная пропитка, ягодная начинка, сливочный крем",
			Price:       "2200₽/кг",
		})
		db.Create(&models.Dessert{
			Name:        "Торт 'Морковный'",
			Description: "морковный бисквит, сливочный крем, карамель, вишневая начинка",
			Price:       "2200₽/кг",
		})
		db.Create(&models.Dessert{
			Name:        "Торт 'Тропики'",
			Description: "ванильный бисквит, молочная пропитка, начинка 'манго-маракуйя', сливочный крем",
			Price:       "2200₽/кг",
		})
		db.Create(&models.Dessert{
			Name:        "Торт 'Вишня-шоколад'",
			Description: "шоколадный бисквит, вишневая пропитка, сливочный крем, вишневая начинка",
			Price:       "2200₽/кг",
		})
		db.Create(&models.Dessert{
			Name:        "Торт 'Сникерс'",
			Description: "шоколадный бисквит, молочная пропитка, сливочный крем, крем из молочного шоколада, карамель, арахис",
			Price:       "2200₽/кг",
		})
		db.Create(&models.Dessert{
			Name:        "Бенто торты",
			Description: "Начинки: 'Ягода-ваниль', 'Шоколад-вишня', 'Тропики' , 'Сникерс'",
			Details:     "вес: 350 гр / в набор входит свечка, вилочка, открытка и подарочная упаковка / *сахарная печать оплачивается отдельно",
			Price:       "1300₽/шт",
			ImageURL:    "https://iili.io/KDB332R.png",
		})
		db.Create(&models.Dessert{
			Name:        "Капкейки",
			Description: "Начинки: 'Ягода-ваниль', 'Шоколад-вишня', 'Сникерс'",
			Details:     "заказ от 6 шт одного вкуса",
			Price:       "280₽/шт",
			ImageURL:    "https://iili.io/KDqyNQj.png",
		})
		db.Create(&models.Review{
			Message:  "Торт очень вкусный, сытный, сочный. Благодарю, чувствуется, что сделан с душой и талантом!",
			ImageURL: "https://iili.io/KDk1roQ.png",
		})
		db.Create(&models.Review{
			Message:  "Торт на день рождения был просто восхитительным! Он не только красивый, но и безумно вкусный.",
			ImageURL: "https://iili.io/KDkNNhF.jpg",
		})
		db.Create(&models.Review{
			Message:  "Торт был вкусный и красивый. Спасибо большое. Сын доволен!",
			ImageURL: "https://iili.io/KDkWeun.jpg",
		})
		db.Create(&models.Review{
			Message:  "Спасибо за торт! Вкусный как всегда ❤️",
			ImageURL: "https://iili.io/KDkUgP1.png",
		})
		db.Create(&models.Review{
			Message:  "Амина, торт бесподобный! Спасибо 😊",
			ImageURL: "https://iili.io/KDkeqVn.png",
		})
	}
}
