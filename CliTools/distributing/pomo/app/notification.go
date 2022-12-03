//go:build !containers && !disable_notification

package app

import "github.com/dtherhtun/Learning-go/CliTools/distributing/notify"

func sendNotification(msg string) {
	n := notify.New("Pomodoro", msg, notify.SeverityNormal)
	n.Send()
}
