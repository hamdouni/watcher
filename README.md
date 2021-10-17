Watcher is a **Go live reload** program for **Linux**.

Watcher monitors the files in a directory and all its subfolders. 
If a write or delete event occurs on any file, the Go program is reloaded.

Watcher ignores any pattern in the **.gitignore** and the **.watcherignore** files.

## Usage

`-d path` the folder to watch (default current folder)
`-r prog` the program to run (default empty)

**Examples**

Watch the current folder (default) and run the program in ./cmd/server

```sh
watcher -r ./cmd/server/ 
```

Watch folder ./pkg and run the program in ./cmd/server

```sh
watcher -r ./cmd/server -d ./pkg
```

Example of a *.watcherignore* file

```
$ cat .watcherignore
.watcherignore
.git
cmd/client
*.md
*.yml
*.sh
*.mod
*.sum
Makefile
```
