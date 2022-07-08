package v1

import (
	"reflect"
	"testing"
)

func TestWallet(t *testing.T) {

	arrert := func(t *testing.T, want, got Bitcoin) {

		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %s want %s", got, want)
		}

	}

	assertBalance := func(t *testing.T, wallet Wallet, got Bitcoin) {

		t.Helper()
		if got != wallet.Balance() {
			t.Errorf("got %s want %s", got, wallet.Balance())
		}
	}

	assertError := func(t *testing.T, got error, want string) {
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got.Error() != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 100}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(90)

		arrert(t, want, got)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		assertError(t, err, "cannot withdraw, insufficient funds")
	})

}
