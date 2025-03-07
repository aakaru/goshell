# 🐚 GoShell

A lightweight interactive shell written in Go! 🚀

## ✨ Features

- 🏃‍♂️ Fast and lightweight command execution
- 🔍 Smart tab completion for commands and files
- 🏠 Built-in commands (`cd`, `pwd`, `echo`, `type`, `exit`)
- 🔄 External command execution
- 🖥️ Beautiful interactive prompt

## 📺 Shell in Action

![GoShell Screenshot](/api/placeholder/650/400)

```
GoShell v0.1 - Type 'exit' to exit.
username@hostname:/home/username$ pwd
/home/username
username@hostname:/home/username$ echo Hello from GoShell!
Hello from GoShell!
username@hostname:/home/username$ type echo
echo is a built in command
username@hostname:/home/username$ cd Documents
username@hostname:/home/username/Documents$ ls
file1.txt  file2.txt  project/
username@hostname:/home/username/Documents$ exit
Exiting GoShell.
```

## 📁 Project Structure

```
goshell/
├── main.go        # Entry point and initialization
├── executor.go    # Command execution engine
├── completion.go  # Tab completion implementation
├── builtins.go    # Built-in command implementations
└── utils.go       # Utility functions
```

## 🔧 Installation

### Prerequisites

- 📦 Go 1.16 or higher
- 📚 Git

```bash
# Clone the repository
git clone https://github.com/yourusername/goshell.git

# Navigate to the project
cd goshell

# Install dependencies
go mod tidy

# Build the project
go build
```

## 🚀 Usage

Run the shell with:

```bash
./goshell
```

### 📝 Built-in Commands

| Command | Description | Usage |
|---------|-------------|-------|
| `cd` | 📂 Change directory | `cd [directory]` |
| `pwd` | 📌 Print working directory | `pwd` |
| `echo` | 🔊 Echo text to standard output | `echo [text]` |
| `type` | 🔍 Display command type | `type [command]` |
| `exit` | 🚪 Exit the shell | `exit [code]` |

## 🛠️ Development

### Building from Source

```bash
go build -o goshell
```

### Dependencies

- 🧩 [github.com/c-bata/go-prompt](https://github.com/c-bata/go-prompt) - Interactive prompt library

## 🤝 Contributing

Contributions are welcome! 🎉 Please feel free to submit a Pull Request.

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- Thanks to the Go community for their amazing tools and libraries
- Special thanks to the creators of go-prompt for making this shell possible