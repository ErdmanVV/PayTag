package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"example/service_payment"
	"example/tokengenerate"
	"example/user_service"
)

func main() {
	// Создаем новый маршрутизатор
	router := mux.NewRouter()

	// Создаем API банка
	bankAPI := service_payment.NewBankAPI("https://example.com/bank/api")

	// Создаем сервис платежей
	paymentService := service_payment.NewPaymentService(bankAPI)

	// Обрабатываем платежи
	router.HandleFunc("/payment", paymentService.HandlePayment).Methods("POST")

	// Создаем сервис генерации токенов
	tokenGenerator := tokengenerate.NewTokenGenerator()

	// Создаем сервис пользователей
	db, err := sql.Open("postgres", "user:password@localhost/database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userService := user_service.NewUserService(db)

	// Обрабатываем регистрацию и вход пользователей
	router.HandleFunc("/register", userService.RegisterUser).Methods("POST")
	router.HandleFunc("/login", userService.LoginUser).Methods("POST")

	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", router)
}
