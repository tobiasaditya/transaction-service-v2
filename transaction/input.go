package transaction

import "transaction-service-v2/user"

type InputTransaction struct {
	Type        string `json:"trxType" binding:"required"`
	Amount      int    `json:"amount" binding:"required"`
	Description string `json:"desc" binding:"required"`
	Method      string `json:"method" binding:"required"`
	User        user.User
}
