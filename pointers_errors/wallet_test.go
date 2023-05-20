package pointers_errors

import "testing"

func TestWallet(t *testing.T) {
	assertError := func(t *testing.T, got error, want error) {
		t.Helper()

		if got == nil {
			t.Fatal("got no error but wanted one")
		}

		if got != want {
			t.Errorf("Got %q, want %q", got, want)
		}
	}
	assertBalance := func(t *testing.T, w Wallet, want Bitcoin) {
		t.Helper()

		got := w.Balance()

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(10)

		want := Bitcoin(10)

		assertBalance(t, w, want)

	})

	t.Run("Withdraw", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(20)
		w.Withdraw(5)

		want := Bitcoin(15)

		assertBalance(t, w, want)
	})

	t.Run("overdraft withdraw", func(t *testing.T) {
		sb := Bitcoin(5)
		w := Wallet{sb}
		err := w.Withdraw(10)

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, w, 5)
	})
}
