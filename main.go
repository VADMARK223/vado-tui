package main

import (
	"fmt"
	"os"
	"vado-tui/app"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	clearTerminal()

	logFileName := "debug.log"
	_ = os.Remove(logFileName)
	logFile, logErr := tea.LogToFile(logFileName, "[TUI] ")
	if logErr != nil {
		fmt.Println("Ошибка создания лога:", logErr)
		os.Exit(1)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			fmt.Println("Ошибка закрытия лога:", logErr)
		}
	}(logFile)

	/*if term.IsTerminal(uintptr(int(os.Stdout.Fd()))) {
		log.Println("Is terminal")
	} else {
		log.Println("Is not terminal")
	}*/

	p := tea.NewProgram(app.NewModel(), tea.WithAltScreen())
	_, err := p.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("✅ The program has completed. See debug.log for logs.")
}

func clearTerminal() {
	// Это ESC (escape)-символ в восьмеричном виде. В ASCII у него код 27, в шестнадцатеричном — 0x1B. Это сигнал терминалу: “дальше идёт управляющая команда”.
	// Начало CSI (Control Sequence Introducer) — указывает, что дальше идут параметры.
	// \033[2J — очищает весь экран
	// \033[H — очищает весь экран
	fmt.Print("\033[2J\033[H")
}
