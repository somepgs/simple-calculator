# Simple Calculator

A simple GUI calculator written in Go and Fyne. It supports basic operations: addition (+), subtraction (-), multiplication (*), and division (/). Handles floating-point numbers and errors like division by zero.

## Features
- Interactive input: Enter expressions like "5 + 3" or "2.5 * 4".
- Error handling: Checks for invalid numbers, unknown operators, and division by zero.

## Requirements
- Go 1.22+ (installed on your system, e.g., via `sudo apt install golang-go` on Ubuntu).

## How to Run
1. Clone the repository:  
   `git clone https://github.com/somepgs/simple-calculator.git`  
   `cd simple-calculator`

2. Build the executable:  
   `go build -o calc`

3. Run:  
   `./calc`  
   Then enter your expression when prompted (e.g., "10 - 2.5").

## Examples
- Input: `5 + 3` → Output: `Result: 8`
- Input: `10 / 2` → Output: `Result: 5`
- Input: `10 / 0` → Output: `Error: division by zero`
- Input: `abc + 3` → Output: `Error: num1 is not a number`

## Development
- Built using standard Go libraries: bufio, strconv, errors.
- To test: Run with different inputs and check outputs.