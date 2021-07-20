package database

func (repository *Repository) CreateUrl(request CreateShortUrlRequest) error {

	if err := repository.Database.Create(request).Error; err != nil {
		return err
	}
	return nil
}

func (CreateShortUrlRequest) TableName() string {
	return "urls"
}
