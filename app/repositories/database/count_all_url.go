package database

func (repository *Repository) CountAllURL() (*uint64, error) {

	var count uint64

	// Count deleted item too
	result := repository.Database.Unscoped().Model(&URL{}).Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}

	return &count, nil
}
