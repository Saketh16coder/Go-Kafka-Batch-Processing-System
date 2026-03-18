package model

import "errors"

func (t *Transaction) Validate() error {

	if t.UserID <= 0 {
		return errors.New("user_id must be greater than 0")
	}

	if t.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	if t.Type != "credit" && t.Type != "debit" {
		return errors.New("type must be 'credit' or 'debit'")
	}

	return nil
}
