# Brainfuck-interpreter
A Brainfuck interpreter written in Go.

## Introduction
This is a simple Brainfuck interpreter implemented in Go. Brainfuck is a minimalistic programming language with only eight commands, designed for fun and educational purposes.

## Prerequisites
To run this interpreter, you need to have Go installed on your system.

### macOS Installation
If you don't have Go installed, you can install it using Homebrew:

```bash
brew install go
```

### Linux Installation
You can install Go on most Linux distributions using your package manager. For example, on Ubuntu:

```bash
sudo apt update
sudo apt install golang-go
```

## Installation

### Clone the repository
Clone the repository or download the `brainfuck.go` file:

```bash
git clone https://github.com/yourusername/brainfuck-interpreter.git
cd brainfuck-interpreter
```

### Create the Go file
If you haven't already, create the `brainfuck.go` file and paste the interpreter code into it:

```bash
nano brainfuck.go
```

Then paste the Go code provided in the `brainfuck.go` file and save it.

### Run the program
To run the program, use the following command:

```bash
go run brainfuck.go
```

This will execute the Brainfuck interpreter with the example Brainfuck code that outputs 'A'.

## Usage

### Running with Default Example
The interpreter comes with a default example that outputs the character 'A'. To run this example, simply execute:

```bash
go run brainfuck.go
```

### Running with Custom Brainfuck Code
You can pass your own Brainfuck code as a command-line argument. For example:

```bash
go run brainfuck.go "++++++++[>++++++++<-]>+."
```

This should output the character 'A'.
