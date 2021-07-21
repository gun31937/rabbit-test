package database

func (repository *Repository) CountAllURL() (*uint64, error) {

	var count uint64

	result := repository.Database.Unscoped().Model(&URL{}).Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}

	return &count, nil
}
