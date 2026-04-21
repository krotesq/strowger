package account

// takes an account and returns the account dto
func toAccountDTO(account *account) accountDto {
	return accountDto{
		UUID:                account.UUID,
		Username:            account.Username,
		Active:              account.Active,
		FailedLoginAttempts: account.FailedLoginAttempts,
		CreatedAt:           account.CreatedAt,
	}
}