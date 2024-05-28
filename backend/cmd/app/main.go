package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	entity2 "mirgalievaal-project/backend/internal/entity"
	handler "mirgalievaal-project/backend/internal/handler/http"
	"mirgalievaal-project/backend/internal/repository"
	repo_sqlite "mirgalievaal-project/backend/internal/repository/sqlite"
	"mirgalievaal-project/backend/internal/service"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func checkEntities() {
	u := entity2.User{
		ID: 1,
		UserRegister: entity2.UserRegister{
			UserLogin: entity2.UserLogin{Email: "mirgalieva.ai@phystech.edu",
				Password: "jbsnlbsnklbs"},
		},
	}
	fmt.Printf("%+v\n\n", u)
	p := entity2.Product{
		CreatedAt:   time.Now(),
		ID:          1,
		Name:        "thermos",
		Description: "very good thermos",
	}

	fmt.Printf("%+v\n\n", p)
}

func checkRepository(repo *repository.Repository) {
	fmt.Println("### Checking repository")

	repo.UserRepository.Create(&entity2.User{UserRegister: entity2.UserRegister{UserLogin: entity2.UserLogin{Email: "mirgalieva.ai@phystech.edu", Password: "jsanjknkj"}}})
	users, _ := repo.UserRepository.GetAll()
	fmt.Printf("Users:%+v\n", *users)

	repo.ProductRepository.Create(&entity2.Product{CreatedAt: time.Now(), SellerID: 1})
	products, _ := repo.ProductRepository.GetAll()
	fmt.Printf("Products:%+v", products)

	newProduct := (*products)[0]
	newProduct.UpdatedAt = time.Now()
	repo.ProductRepository.Update(&newProduct)

	products, _ = repo.ProductRepository.GetAll()
	fmt.Printf("Products: %+v\n", *products)
	fmt.Println("### Checking repository end")
}

func checkService(s *service.Service) {
	fmt.Println("########## Checking service")

	fmt.Println("Try to register an user:")
	userReg := entity2.UserRegister{UserLogin: entity2.UserLogin{Email: "mirgalieva", Password: "adhcakbjkb"}, FirstName: "Alsu", LastName: "Mirgalieva"}
	err := s.User.Register(&userReg)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("OK")
	}
	fmt.Println("Try to log in with INCORRECT password")
	err = s.User.Login(&entity2.UserLogin{Email: "mirgalieva", Password: "snjdvnljdsvnl"})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("OK")
	}

	fmt.Println("Try to log in with CORRECT password")
	err = s.User.Login(&entity2.UserLogin{Email: "mirgalieva", Password: "adhcakbjkb"})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("OK")
	}

	user, err := s.User.Get(3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%+v\n", *user)
	}
	fmt.Println("Try to create new product for user")
	product, err := s.Product.Create(&entity2.Product{SellerID: 1, Name: "thermos", Price: "300", Description: "nice thermos"})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%+v\n", *product)
	}
	fmt.Println("Getting info about user 3")
	user, err = s.User.Get(3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%+v\n", *user)
	}
	fmt.Println("####### Ended chcking service")
	fmt.Println()
}

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	fmt.Println("Hello world!")
	checkEntities()
	dbUri, ok := os.LookupEnv("DB_URI")
	if !ok {
		log.Println("cannot get DB_URI from ENV")
		dbUri = "test.db"
	}
	db, err := repo_sqlite.NewSQLIte(dbUri)
	if err != nil {
		log.Panic("failed to initialize database:", err)
	} else {
		log.Println("database is initialized")
	}
	repo := repository.NewRepository(db)
	//checkRepository(repo)

	service := service.NewService(repo)
	//checkService(service)
	h := handler.NewHandler(service)
	srv := &http.Server{
		Addr: ":8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h.NewRouter(), // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
