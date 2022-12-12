package factory

import (
	homestayDelivery "github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay/delivery"
	homestayRepo "github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay/repository"
	homestayService "github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	homestayRepoFactory := homestayRepo.New(db)
	homestayServiceFactory := homestayService.New(homestayRepoFactory)
	homestayDelivery.New(homestayServiceFactory, e)

}
