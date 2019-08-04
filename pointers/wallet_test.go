package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("should deposit bitcoins to wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("should withdraw funds from wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("should not be able to overdraw funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(21))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, insufficientFundsError)
	})

}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()

	balance := wallet.Balance()

	if balance != want {
		t.Errorf("got %s, expected %s", balance, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("Expected error but got none.")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Fatalf("Expected no errors but got %v.", got)
	}
}
