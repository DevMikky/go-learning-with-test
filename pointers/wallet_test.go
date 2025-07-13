package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit test", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw test", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(20)}

		want := Bitcoin(10)

		err := wallet.Withdraw(Bitcoin(10))

		assertNotError(t, err)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {

		startingBalance := Bitcoin(20)

		wallet := Wallet{startingBalance}

		gotErr := wallet.Withdraw(Bitcoin(100))

		wantError := ErrInsufficientFunds

		assertError(t, gotErr, wantError)

		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {

	t.Helper()

	if wallet.Balance() != want {
		t.Errorf("got %s; want %s", wallet.Balance(), want)
	}
}

func assertError(t *testing.T, gotErr error, wantErr error) {

	t.Helper()

	if gotErr == nil {
		t.Fatalf("got == nil, want %q", gotErr)
	}

	if gotErr.Error() != wantErr.Error() {
		t.Errorf("got %q; want %q", gotErr, wantErr)
	}
}

func assertNotError(t *testing.T, gotErr error) {

	t.Helper()

	if gotErr != nil {
		t.Fatalf("got %q, want nil", gotErr)
	}
}
