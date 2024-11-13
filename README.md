Watcher is a **Go live reload or retest** program for **Linux**.

Watcher monitors the files in a directory and all its subfolders. If a write or
delete event occurs on any file, the Go program or the tests are reloaded.

If a **.gitignore** file is present, Watcher will use it to exclude files from
monitoring. Use **.watcherignore** file to add patterns to ignore that can not
be added to .gitignore. 

Watcher also accepts to pass arguments to the program via the `args` flag.
Arguments must be surrounded with quotes.

## Usage

* `-dir path` folder to watch (default current folder)
* `-run prog` program to run (default empty)
* `-args "arguments..."` pass arguments to the program (default empty). This
  has no effect with -test.
* `-test` launch tests instead of the program
* `-verbosetest` launch tests with verbosity
* `-help` view the command usage 

**Examples**

Watch the current folder (default) and run the program in ./cmd/server with
some arguments (note the quotes)

```sh
watcher -run ./cmd/server/ -args "-name John -age 24"
```

Watch folder ./pkg and run the program in ./cmd/server

```sh
watcher -run ./cmd/server -dir ./pkg
```

Watch current folder and retest 

```sh 
watcher -test
```
Same thing but test with verbosity

```sh
watcher -verbosetest
```

Watch current folder and retest all subfolders

```sh 
watcher -test -run ./...
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
