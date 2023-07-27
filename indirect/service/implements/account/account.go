package account

import "gotest/indirect/service/interfaces/dto"

type SAccount struct {
}

func (s *SAccount) Create(in dto.AccountCreateInDto) dto.AccountCreateOutDto {
	return dto.AccountCreateOutDto{
		Account:  in.Account,
		Password: in.Password,
	}
}
