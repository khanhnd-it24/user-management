package databases

import "user-management/core/entities"

func AutoMigrate() {

	err := DB.AutoMigrate(&entities.User{})
	if err != nil {
		panic("could not connect to db")
		return
	}
	Execute(DB)

}
