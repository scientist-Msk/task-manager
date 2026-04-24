package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	commands "study/comands"
	"study/strct"

	"github.com/k0kubun/pp"
)

func main() {
	tasks := make(map[strct.Task]bool)
	tasksD := make(map[strct.TaskD]bool)
	var Evnts [][]string

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Введите команду: ")
		if ok := scanner.Scan(); !ok {
			fmt.Println("Ошибка ввода!")
			return
		}
		text := scanner.Text()

		if text == "" {
			fmt.Println("Вы ничего не ввели!")
			Evnts = append(Evnts, []string{"<Пустой ввод> | Описание: Вы ничего не ввели!"})
			continue
		}
		fields := strings.Fields(text)

		cmd := fields[0]

		if cmd == "exit" {
			fmt.Println("Вы завершили выполнение программы, до встречи!")
			Evnts = append(Evnts, fields)
			return
		}

		if cmd == "help" {
			commands.Help()
			Evnts = append(Evnts, fields)
			continue
		}
		if cmd == "add" {
			if len(fields) < 3 {
				fmt.Println("Неверный формат ввода!")
				Evnts = append(Evnts, append(fields, "| Описание: Неверный формат ввода!"))
				continue
			}
			textTask := ""
			for i := 2; i < len(fields); i++ {
				if i > 2 {
					textTask += " "
				}
				textTask += fields[i]
			}
			key := strct.NewTask(fields[1], textTask)
			tasks[key] = false

			fmt.Println("Задача добавлена")
			Evnts = append(Evnts, fields)
			continue
		}

		if cmd == "list" {
			pp.Println(tasks, tasksD)
			Evnts = append(Evnts, fields)
			continue
		}

		if cmd == "done" {
			if len(fields) < 2 {
				fmt.Println("Укажите заголовок!")
				Evnts = append(Evnts, append(fields, "| Описание: Укажите заголовок!"))
				continue
			}
			key := strct.TaskComplete(tasks, fields[1])
			if key.HeadingD != "" {
				tasksD[key] = true
				fmt.Println("Задача", fields[1], "выполнена!")
			} else {
				fmt.Println("Задача не найдена!")
				Evnts = append(Evnts, append(fields, "| Описание: Задача не найдена!"))
				continue
			}
			Evnts = append(Evnts, fields)
			continue

		}
		if cmd == "events" {
			Evnts = append(Evnts, fields)
			for i, v := range Evnts {
				fmt.Printf("Событие %d. Пользователь написал: %s\n", i+1, v)
			}
			continue
		}

		if cmd == "del" {
			if len(fields) < 2 {
				fmt.Println("Укажите заголовок!")
				Evnts = append(Evnts, append(fields, "| Описание: Укажите заголовок!"))
				continue
			}
			if strct.DeleteTaskByHeading(tasks, fields[1]) {
				fmt.Println("Задача удалена!")
				Evnts = append(Evnts, fields)
				continue
			} else {
				fmt.Println("Задача не найдена!")
				Evnts = append(Evnts, append(fields, "| Описание: Задача не найдена!"))
				continue
			}
		}
		fmt.Println("Передана неизвестная команда!")
		Evnts = append(Evnts, append(fields, "| Описание: Передана неизвестная команда!"))
		continue
	}
}
