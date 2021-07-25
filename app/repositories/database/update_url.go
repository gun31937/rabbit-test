package database

func (repository *Repository) UpdateURL(id uint, request UpdateShortURLRequest) error {

	model := request.parseModel()

	if err := repository.Database.Model(&model).Where("id = ?", id).Update(&model).Error; err != nil {
		return err
	}
	return nil
}

func (request UpdateShortURLRequest) parseModel() URL {
	return URL{
		Hits: request.Hits,
	}
}
