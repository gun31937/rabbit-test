package database

func (repository *Repository) DeleteURL(shortCode string) error {

	if err := repository.Database.Where("short_code = ?", shortCode).Delete(&URL{}).Error; err != nil {
		return err
	}

	return nil
}
