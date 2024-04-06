package pointersnerrors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with insufficient funds", func(t *testing.T) {

		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(30))
		assertBalance(t, wallet, startingBalance)

		assertError(t, err, ErrInsufficientFunds)
	})
	t.Run("Withdraw", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
        assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

}
func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.balance

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		// fatal means that it will stop the test
		t.Fatal("Expecting an error but didnt get one")
	}

	if got.Error() != want.Error() {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("This should not have errord ")
	}
}
