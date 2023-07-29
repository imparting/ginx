package account

import (
	"fmt"
	"gotest/indirect/db"
	"gotest/indirect/model"
	"gotest/indirect/service/interfaces/dto"
)

type SAccount struct {
}

func (s *SAccount) Create(in dto.AccountCreateInDto) dto.AccountCreateOutDto {
	// 判断表是否存在
	if !db.HasTable(&model.User{}) {
		if err := db.AutoMigrate(&model.User{}); err != nil {
			panic("failed to create table")
		}
	}

	user := model.User{
		Username: in.Account,
		Password: in.Password,
	}

	res := db.Table("user").Create(&user)
	u := &model.User{}
	db.Table("user").Where("id = ?", user.ID).First(u)
	fmt.Println(res)

	return dto.AccountCreateOutDto{
		Account:  in.Account,
		Password: in.Password,
		Count:    int64(u.ID),
	}
}
