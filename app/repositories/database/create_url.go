package database

func (repository *Repository) CreateURL(request CreateShortURLRequest) (*uint, error) {

	model := request.parseModel()

	if err := repository.Database.Create(&model).Error; err != nil {
		return nil, err
	}

	return &model.ID, nil
}

func (request CreateShortURLRequest) parseModel() URL {
	return URL{
		ShortCode: request.ShortCode,
		FullUrl:   request.FullURL,
		Expiry:    request.Expiry,
	}
}
