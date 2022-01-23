package storage

import (
	"database/sql"
	"github.com/LebedevNd/BannerRotation/internal/storage/banner"
	"github.com/LebedevNd/BannerRotation/internal/storage/group"
	"github.com/LebedevNd/BannerRotation/internal/storage/slot"
)

type Storage struct {
	slot   slot.Slot
	group  group.Group
	banner banner.Banner
	db     *sql.DB
}
