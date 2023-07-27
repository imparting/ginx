package interfaces

import "gotest/indirect/service/interfaces/dto"

type IAccount interface {
	Create(in dto.AccountCreateInDto) dto.AccountCreateOutDto
}
