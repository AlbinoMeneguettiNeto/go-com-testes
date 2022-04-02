package pne

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("\n Got:  %s\n Want: %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		assertBalance(t, wallet, 10)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}

		wallet.Withdraw(10)

		assertBalance(t, wallet, 10)
	})
}
