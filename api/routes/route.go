package routes

import (
	"github.com/devfurkankizmaz/go-lib-management-app/api/handlers"
	"github.com/devfurkankizmaz/go-lib-management-app/api/middleware"
	"github.com/devfurkankizmaz/go-lib-management-app/configs"
	"github.com/devfurkankizmaz/go-lib-management-app/repository"
	"github.com/devfurkankizmaz/go-lib-management-app/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Setup(env *configs.Env, db *gorm.DB, g *echo.Echo) {
	// All Public APIs
	public := g.Group("/api/auth")
	NewRegisterRouter(env, db, public)
	NewLoginRouter(env, db, public)
	NewRefreshRouter(env, db, public)

	// All Protected APIs
	protected := g.Group("", middleware.MiddlewareJWT)
	NewProfileRouter(env, db, protected)
	NewBookRouter(env, db, protected)
}

func NewRegisterRouter(env *configs.Env, db *gorm.DB, group *echo.Group) {
	r := repository.NewUserRepository(db)
	h := handlers.RegisterHandler{
		RegisterService: service.NewRegisterService(r),
		Env:             env,
	}
	group.POST("/register", h.Register)
}

func NewLoginRouter(env *configs.Env, db *gorm.DB, group *echo.Group) {
	r := repository.NewUserRepository(db)
	h := &handlers.LoginHandler{
		LoginService: service.NewLoginService(r),
		Env:          env,
	}
	group.POST("/login", h.Login)
}

func NewRefreshRouter(env *configs.Env, db *gorm.DB, group *echo.Group) {
	r := repository.NewUserRepository(db)
	h := &handlers.RefreshHandler{
		RefreshService: service.NewRefreshService(r),
		Env:            env,
	}
	group.POST("/refresh", h.Refresh)
}

func NewProfileRouter(env *configs.Env, db *gorm.DB, group *echo.Group) {
	r := repository.NewUserRepository(db)
	h := &handlers.ProfileHandler{
		ProfileService: service.NewProfileService(r),
	}
	group.GET("/api/me", h.Fetch)
}

func NewBookRouter(env *configs.Env, db *gorm.DB, group *echo.Group) {
	r := repository.NewBookRepository(db)
	h := &handlers.BookHandler{
		BookService: service.NewBookService(r),
	}
	group.POST("/api/books", h.Create)
	group.GET("/api/books", h.FetchAllByUserID)
	group.GET("/api/books/:bookId", h.FetchByID)
	group.PUT("/api/books/:bookId", h.UpdateByID)
	group.DELETE("/api/books/:bookId", h.DeleteByID)
}
