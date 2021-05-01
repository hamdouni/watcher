Watcher is a **Go live reload** program for **Linux**.

It monitors the files in the actual directory and all its subfolders. If a write or delete event occurs on any file except those matching any pattern in the **.gitignore** file, the Go program is reloaded by killing its process, building and launching it.

I use it for live reloading my web app written in Go.