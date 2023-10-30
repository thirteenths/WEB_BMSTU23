package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/thirteenths/WEB_BMSTU23/backend/docs"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/service"
)

type Handler struct {
	service *service.Service
	router  *gin.Engine
}

func NewHandler(service *service.Service) *Handler {
	// Gin instance
	router := gin.New()
	return &Handler{service: service, router: router}
}

func (h *Handler) GetRoutes() *gin.Engine {
	return h.router
}

func (h *Handler) InitRoutes() *gin.Engine {
	// Routes
	auth := h.router.Group("/auth")
	{
		auth.POST("/singup", h.singUp)
		auth.POST("/singin", h.singIn)
	}

	// Routes
	comics := h.router.Group("/comics")
	{
		comics.GET("/", h.GetAllComic)
		comics.GET("/:id", h.GetComic)

		comics.POST("/", h.CreateComic, h.userIdentity)
		comics.DELETE("/:id", h.DeleteComic)

		comics.PUT("/:id", h.UpdateComic)
		comics.PATCH("/:id", h.UpdateComicKek)
	}

	// Routes
	posters := h.router.Group("/posters")
	{
		posters.GET("/", h.GetAllPoster)
		posters.GET("/:id", h.GetPoster)

		posters.POST("/", h.CreatePoster)
		posters.DELETE("/:id", h.DeleteComic)

		posters.POST("/:id", h.CheckInPoster, h.userIdentity)
		posters.PUT("/:id", h.UpdatePoster)
	}

	// Routes
	events := h.router.Group("/events")
	{
		events.GET("/", h.GetAllEvent)
		events.GET("/:id", h.GetEvent)

		events.POST("/", h.CreateEvent)
		events.DELETE("/:id", h.DeleteEvent)

		events.PUT("/:id", h.UpdateEvent)
	}

	// Routes
	places := h.router.Group("/places")
	{
		places.GET("/", h.GetAllPlace)
		places.GET("/:id", h.GetPlace)

		places.POST("/", h.CreatePlace)
		places.DELETE("/:id", h.DeletePlace)

		places.PUT("/:id", h.UpdatePlace)
	}

	// Routes
	h.router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return h.router
}
