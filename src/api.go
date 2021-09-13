package main

import (
	"crypto/sha1"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/jichall/stone/src/crypt"
	"github.com/jichall/stone/src/models"
	"github.com/labstack/echo/v4"
)

func fetchAccounts(c echo.Context) error {
	return c.JSON(200, accounts)
}

func createAccount(c echo.Context) error {
	account := new(models.Account)
	if err := c.Bind(account); err != nil {
		return nil
	}

	account.ID = accountsId
	account.Balance = 0
	account.Secret = crypt.Encrypt(account.Secret)
	account.CreatedAt = time.Now()

	accounts[accountsId] = account
	accountsId++

	return c.JSON(200, account)
}

func fetchBalance(c echo.Context) error {

	id, err := strconv.ParseInt(c.Param("id"), 10, 8)

	if err != nil {
		return c.JSON(400, "failed to parse id param")
	}

	account := accounts[id]

	return c.JSON(200, &models.Balance{Amount: account.Balance})
}

func login(c echo.Context) error {

	auth := new(models.Authentication)

	if err := c.Bind(auth); err != nil {
		return c.JSON(400, "failed to parse data")
	}

	// does the account exists?
	var account *models.Account

	for _, acc := range accounts {
		if acc.CPF == auth.CPF {
			account = acc
			break
		}
	}

	if account != nil {
		hasher := sha1.New()
		hasher.Write([]byte(auth.CPF + auth.Secret))

		token := hex.EncodeToString(hasher.Sum(nil)[:16])

		// create a token entity and save it in the storage
		t := models.Token{Token: token}
		tokens[account.ID] = &t

		return c.JSON(200, t)
	}

	return c.JSON(404, "account does not exists")
}

func fetchTransfers(c echo.Context) error {

	// TODO: retrieve id from token and authentication

	transaction := new(models.Transaction)

	if err := c.Bind(transaction); err != nil {
		return c.JSON(400, "failed to parse data")
	}

	// does the origin and the destiny account exists?
	var origin *models.Account
	var destiny *models.Account

	for _, acc := range accounts {
		// performance penalty over here (?) It probably would be better to
		// convert the AccountOrigin and AccountDestination to int64 and compare
		// them directly
		id := strconv.Itoa(int(acc.ID))

		if transaction.AccountOrigin == id {
			origin = acc
			break
		}
		if transaction.AccountDestination == id {
			destiny = acc
			break
		}
	}

	if origin != nil && destiny != nil {
		if origin.Balance <= transaction.Amount {
			return c.JSON(400, "insufficient funds")
		}

		origin.Balance -= transaction.Amount
		destiny.Balance += transaction.Amount
	}

	return c.JSON(400, "inexistent origin or destiny account")
}

func createTransfer(c echo.Context) error {

	// TODO: authentication

	return nil
}
