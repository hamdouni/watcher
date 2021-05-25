Watcher is a **Go live reload** program for **Linux**.

It monitors the files in the actual directory and all its subfolders. If a write or delete event occurs on any file except those matching any pattern in the **.gitignore** file, the Go program is reloaded by killing its process, building and launching it.

Since version 0.0.2, you can also use a **.watcherignore** file to store alternate patterns not suitable for the .gitignore file. For example, I put all html/js stuff that I don't want to trigger the Go recompilation.

I use it for live reloading my web app written in Go.