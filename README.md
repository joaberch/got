# Got
[![Go Report Card](https://goreportcard.com/badge/github.com/joaberch/got)](https://goreportcard.com/report/github.com/joaberch/got)
A minimalist version control system built in Go.

---

## Features
- Initializes a local repository
- Add files to staging area and commit it with a message
- View current status of tracked files
- Restore previous commits using its hash
- Pure Go CLI with **0** external dependencies

---

## How to use

### Windows

1. **Download the latest release from *[Github Releases](https://github.com/joaberch/got/releases)***
2. **Extract the ZIP archive**
3. **Run ``setup-windows.bat`` as an administrator**.

This will :
- Move the binary to a ``utils`` folder in the user directory.
- Add the ``utils`` folder to the system ``PATH`` variable so ``gosearch`` can be run everywhere.

> Manual setup :
> - Move the binary to a specific folder
> - Add that folder to the system ``PATH``

### Linux / MacOS
1. **Download the latest release from *[Github Releases](https://github.com/joaberch/got/releases)***.
2. Extract the TAR archive.
3. Run the setup script : `bash setup-linux.sh`.

This will :
- Move the binary to `~/utils`.
- Add the utils folder to the system ``PATH`` using ``.bashrc``.

> Manual setup :
> - Move the binary to ``~\utils``
> - Add ``export PATH="$PATH:$HOME/utils"`` to .bashrc or .zshrc

---

## Usage

### Help and version

```
got help                        # Show help
got version                     # Show current version
```

### Initialize a repository

```
got init                        # Create a .got repository
```

### Add and commit files

```
got add main.go                 # Add the file to the staging area
got commit "Initial commit"     # Creates a commit with a message
```

### Restore a file

```
got restore abc123commithash    # Restore the file from commit hash
```

---

## Requirements
- Go 1.24.4 or higher
- Windows, Linux, macOS
- Admin rights (optional for setup scripts)

---

## Roadmap
- [x] Basic init/stage/commit
- [x] Restore commit
- [ ] Use the branch system
- [ ] got diff
- [ ] Push on remote server
- [ ] Restore using a more user-friendly way

---

## Contributing
Pull requests are welcome!

---

## License
This project is licensed under the MIT License.
See [LICENSE](https://github.com/joaberch/got/blob/main/LICENSE) for more information.
