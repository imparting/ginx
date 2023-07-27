package service

import (
	"gotest/indirect/service/implements/account"
	"gotest/indirect/service/interfaces"
)

func Account() interfaces.IAccount {
	return &account.SAccount{}
}
