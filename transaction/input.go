package transaction

import "transaction-service-v2/user"

type InputTransaction struct {
	TrxType string `json:"trxType" binding:"required"`
	Amount  string `json:"amount" binding:"required"`
	Desc    string `json:"desc" binding:"required"`
	Method  string `json:"method" binding:"required"`
	User    user.User
}
