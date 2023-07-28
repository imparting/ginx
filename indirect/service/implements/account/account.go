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
	user := model.User{
		Username: in.Account,
		Password: in.Password,
	}
	res := db.Table("user").Create(&user)
	fmt.Println(res)

	return dto.AccountCreateOutDto{
		Account:  in.Account,
		Password: in.Password,
		Count:    int64(user.ID),
	}
}
