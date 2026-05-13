package di

import (
	genreApp "game-hub-backend/internal/application/genre"
	userApp "game-hub-backend/internal/application/user"

	httpDelivery "game-hub-backend/internal/delivery/http"
	"game-hub-backend/internal/infra/auth"
	"game-hub-backend/internal/infra/config"
	"game-hub-backend/internal/infra/database/repository"

	"gorm.io/gorm"
)

type AppHandlers struct {
	UserHandler  *httpDelivery.UserHandler
	GenreHandler *httpDelivery.GenreHandler
}

func InitHandlers(db *gorm.DB, cfg *config.Config) *AppHandlers {

	// Services
	jwtService := auth.NewJWTService(cfg.JWTSecret)

	// User
	userRepo := repository.NewUserRepository(db)
	userUseCase := userApp.NewUserUseCase(userRepo, jwtService)
	userHandler := httpDelivery.NewUserHandler(userUseCase)

	//Genres
	genreRepo := repository.NewGenreRepository(db)
	genreUseCase := genreApp.NewGenreUseCase(genreRepo)
	genreHandler := httpDelivery.NewGenreHandler(genreUseCase)

	return &AppHandlers{
		UserHandler:  userHandler,
		GenreHandler: genreHandler,
	}
}
