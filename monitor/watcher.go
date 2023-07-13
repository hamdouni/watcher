/*
	Use package rjeczalik/notify to monitor the Linux file system with inotify specific events.
*/
package monitor

import (
	"github.com/rjeczalik/notify"
)

// Watch return a channel of all events happening in the srcpath directory
// tree. Dispatch only Linux inotify modification events. The channel must be
// closed with the Stop function by the caller (e.g using defer)
func Watch(srcpath string) (chan notify.EventInfo, error) {
	c := make(chan notify.EventInfo, 1)
	err := notify.Watch(srcpath+"/...", c, notify.InCloseWrite, notify.InMovedTo, notify.Remove)
	return c, err
}

// Stop the channel
func Stop(c chan notify.EventInfo) {
	defer notify.Stop(c)
}
