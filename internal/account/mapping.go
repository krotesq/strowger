package account

// takes an account and returns the account dto
func toAccountDTO(account *account) accountDTO {
	return accountDTO{
		ID:                  account.ID,
		Username:            account.Username,
		Active:              account.Active,
		FailedLoginAttempts: account.FailedLoginAttempts,
		CreatedAt:           account.CreatedAt,
	}
}
