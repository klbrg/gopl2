package storage

import (
	"strings"
	"testing"
)

// setNotifier installs f as the notification function for the
// duration of the test, restoring the old one afterwards.
func setNotifier(t *testing.T, f func(username, msg string)) {
	t.Helper()
	saved := notifyUser
	t.Cleanup(func() { notifyUser = saved })
	notifyUser = f
}

// setUsage records a byte count for user for the duration of the test.
func setUsage(t *testing.T, user string, n int64) {
	t.Helper()
	saved, ok := usage[user]
	t.Cleanup(func() {
		if ok {
			usage[user] = saved
		} else {
			delete(usage, user)
		}
	})
	usage[user] = n
}

func TestCheckQuotaNotifiesUser(t *testing.T) {
	var notifiedUser, notifiedMsg string
	setNotifier(t, func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	})

	const user = "joe@example.org"
	setUsage(t, user, 980000000) // 980MB used

	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
	}
	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, "+
			"want substring %q", notifiedMsg, wantSubstring)
	}
}

func TestCheckQuotaSilentBelowThreshold(t *testing.T) {
	called := false
	setNotifier(t, func(string, string) { called = true })

	const user = "jane@example.org"
	setUsage(t, user, 500000000) // 500MB used

	CheckQuota(user)
	if called {
		t.Errorf("notifyUser called at 50%% of quota")
	}
}
