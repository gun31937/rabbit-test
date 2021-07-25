package database

func (repository *Repository) ListURL(filter ListURLFilterRequest) ([]URL, error) {

	var urls []URL

	// Will return deleted items too.
	query := repository.Database.Unscoped().Model(&URL{}).Order("id ASC")

	if filter.ShortCode != nil {
		query = query.Where("short_code = ?", *filter.ShortCode)
	}

	if filter.Keyword != nil {
		query = query.Where("full_url like ?", `%`+*filter.Keyword+`%`)
	}

	err := query.Find(&urls).Error

	if err != nil {
		return nil, err
	}

	return urls, nil
}
