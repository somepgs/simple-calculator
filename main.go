package main

import (
	"errors"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")

	display := widget.NewEntry()
	display.SetPlaceHolder("0")

	appendToDisplay := func(text string) {
		display.SetText(display.Text + text)
	}

	buttons := [][]string{
		{"7", "8", "9", "/"},
		{"4", "5", "6", "*"},
		{"1", "2", "3", "-"},
		{"0", "C", "=", "+"},
	}

	grid := container.NewGridWithColumns(4)
	for _, row := range buttons {
		for _, btnText := range row {
			btn := widget.NewButton(btnText, func() {
				if btnText == "=" {
					result, err := evaluate(display.Text)
					if err != nil {
						display.SetText("Error: " + err.Error())
					} else {
						display.SetText(strconv.FormatFloat(result, 'f', -1, 64))
					}
				} else if btnText == "C" {
					display.SetText("")
				} else {
					appendToDisplay(btnText)
				}
			})
			grid.Add(btn)
		}
	}

	content := container.NewVBox(display, grid)
	w.SetContent(content)
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}

func evaluate(expr string) (float64, error) {
	expr = strings.TrimSpace(expr)

	opIndex := strings.IndexAny(expr, "+-*/")
	if opIndex == -1 || opIndex == len(expr)-1 || opIndex == 0 {
		return 0, errors.New("invalid expression")
	}

	num1Str := expr[:opIndex]
	op := string(expr[opIndex])
	num2Str := expr[opIndex+1:]

	num1, err := strconv.ParseFloat(num1Str, 64)
	if err != nil {
		return 0, err
	}
	num2, err := strconv.ParseFloat(num2Str, 64)
	if err != nil {
		return 0, err
	}
	return calculate(num1, num2, op)
}

func calculate(num1, num2 float64, op string) (float64, error) {
	switch op {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("division by zero")
		}
		return num1 / num2, nil
	default:
		return 0, errors.New("unknown operator")
	}
}
