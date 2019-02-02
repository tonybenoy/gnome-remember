package main

import (
	"github.com/mqu/go-notify"
	"time"
)

func main() {
	var header string
	var notifys string
	var dialinfo string
	var delay int
	dialinfo = ""
	notifys = "This"
	header = "Main"
	delay = 5000
	notifynow(header, delay, notifys, dialinfo)

}
func addtofile(completion string, priority string, completiondate time)
func notifynow(Header string, delay int, Notification string, DialInfo string) {

	notify.Init("Remember")
	remember := notify.NotificationNew(Header, Notification, DialInfo)
	notify.NotificationSetTimeout(remember, 5000)
	remember.Show()
}
