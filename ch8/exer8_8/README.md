
- `walkDir` function will loop through `[]os.DirEntry` in the directory. It will recursively call `walkDir` if entry is a directory, else it will send file size to `fileSize` channel. To limit the number of concurrent goroutine at the time, we use semophore mechanism, create a buffer channel named `token`.

- Create `DirWalker` struct to calculate total size of a directory. It contains fields:
    - `dir` string: path to the directory 
    - `fizeSize` channel: to send and receive from goroutines 
    - `n` sync.WaitGroup: to count active goroutines and continue when no goroutine is running.


### running result

Cal total of all file in dirs

```
du -h -s $HOME /usr
78G $HOME
11G /usr
```

```
go run main.go $HOME /usr
89.0 GB
```

Cal total of every dir

```
du -h -s $HOME /usr
78G $HOME
11G /usr
```

```
go run main.go $HOME /usr
/usr 262508 file 10.2 GB
$HOME 542155 file 78.8 GB
```

another version: create a goroutine when walk every dir (include subdir)
```
➜  goplexercise git:(main) ✗ time du -h -s $HOME /usr
78G     $HOME
11G     /usr
du -h -s $HOME /usr  0.76s user 1.88s system 98% cpu 2.665 total
```

```
➜  goplexercise git:(main) ✗ time go run main.go $HOME /usr
[$HOME /usr]
/usr 262508 file 10.2 GB
$HOME 551787 file 78.8 GB
running in 1.350180911s
go run main.go $HOME /usr  5.72s user 4.59s system 699% cpu 1.475 total
```

