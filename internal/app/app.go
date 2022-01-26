package app

import "github.com/LebedevNd/BannerRotation/internal/models/database"

type App struct {
	*database.BannerModel
	*database.SlotModel
	*database.GroupModel
}
