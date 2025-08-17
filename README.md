# Simple Calculator

A modern graphical calculator application written in Go using the Fyne GUI toolkit. Features a clean interface with
buttons for basic arithmetic operations including addition (+), subtraction (-), multiplication (*), and division (/).
Handles floating-point numbers and provides error feedback.

## Features

- Graphical user interface with numeric keypad and operation buttons
- Input display shows the current expression and results
- Supports basic arithmetic operations: +, -, *, /
- Error handling for:
    - Invalid expressions
    - Division by zero
    - Invalid number formats
- Clear button (C) to reset input
- Floating point number support

## Requirements

- Go 1.22+ (installed on your system)
- Fyne GUI toolkit
- Compatible OS: Windows, macOS, Linux

## Installation & Running

1. Install Go and Fyne dependencies:
   ```bash
   go get fyne.io/fyne/v2
   ```

2. Clone the repository:
   ```bash
   git clone https://github.com/somepgs/simple-calculator.git
   cd simple-calculator
   ```

3. Run directly:
   ```bash
   go run main.go
   ```

   Or build and run the executable:
   ```bash
   go build
   ./simple-calculator
   ```

## Usage

1. Click number buttons to enter values
2. Click operation buttons (+, -, *, /) to select operation
3. Press = to calculate a result
4. Use C to clear and start over

## Error Handling

- "Error: division by zero" - Attempt to divide by zero
- "Error: invalid expression" - Malformed or incomplete expression
- Other parsing errors for invalid number formats

## Development

- Built using Go 1.22+ and Fyne v2
- Main components:
    - GUI layout using Fyne containers and widgets
    - Expression evaluation
    - Basic arithmetic operations
    - Error handling