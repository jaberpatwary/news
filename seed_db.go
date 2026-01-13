package main

import (
	"app/src/model"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=root dbname=news port=5432 sslmode=disable TimeZone=Asia/Dhaka"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Seeding database with 20 news items...")

	categories := []string{"জাতীয়", "রাজনীতি", "খেলাধুলা", "প্রযুক্তি", "স্বাস্থ্য", "বিনোদন"}

	for i := 1; i <= 20; i++ {
		isFeatured := i <= 6
		article := model.Article{
			Title:    fmt.Sprintf("টেস্ট নিউজ শিরোনাম %d", i),
			Content:  fmt.Sprintf("এটি একটি বিস্তারিত টেস্ট নিউজ কন্টেন্ট। এই নিউজটি শুধুমাত্র লেআউট চেক করার জন্য তৈরি করা হয়েছে। নিউজ নম্বর %d।", i),
			Category: categories[i%len(categories)],
			Author:   "অ্যাডমিন",
			Created:  time.Now().Add(time.Duration(-i) * time.Hour),
			Featured: isFeatured,
		}
		db.Create(&article)
	}

	fmt.Println("Seeding complete! 6 featured and 14 regular news added.")
}
