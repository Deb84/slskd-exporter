package postgres

type File struct {
	ID         int64  `gorm:"primaryKey;autoIncrement"`
	Path       string `gorm:"not null;uniqueIndex"`
	FileName   string `gorm:"index"`
	ArtistName string `gorm:"index"`
	AlbumName  string `gorm:"index"`
	Year       string
}
