package postgres

type Batch struct {
	ID int64 `gorm:"primaryKey;autoIncrement"`
}
