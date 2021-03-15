package boot

import "antapi/db"

func SetUpCollections() error {
	// MongDB need not to create table
	if db.UseMongo {
		return nil
	}

	return nil
}

func SetUpDefaultsData() error {
	return nil
}
