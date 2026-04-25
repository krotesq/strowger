package account

// takes an account and returns the account dto
func toAccountDTO(account *account) accountDTO {
	return accountDTO{
		UUID:                account.UUID,
		Username:            account.Username,
		Active:              account.Active,
		FailedLoginAttempts: account.FailedLoginAttempts,
		CreatedAt:           account.CreatedAt,
	}
}