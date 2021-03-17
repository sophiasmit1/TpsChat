package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/richard-lyman/lithcrypt" //Пак с шифром
)

var source string

//Из введеного достает команду
func comand() string {
	arr := []byte(source) //из переменной q делаем массив arr
	y := arr[0:2]         //из arr вырезаем символы от 0 до 2 в переменную y
	return string(y)      //Функцие proverka даем значение y в расширении string
}

//Из введеного достает id
func id() string {
	arr := []byte(source) //из переменной q делаем массив arr
	y := arr[3:4]         //переменной даем символы от arr от 3 до 4 символа
	return string(y)      //Функцие proverka даем значение y в расширении string
}

//Из введеного достает текст
func text() string {
	arr := []byte(source) //из переменной q делаем массив arr
	x := arr[5:len(arr)]  //переменной даем символы от arr от 5 символа до длины arr
	return string(x)      //Функцие proverka2 даем значение y в расширении string
}

//Функция... Эмм... Пробела?....
func probel() string {
	arr := []byte(source) //из переменной q делаем массив arr
	y := arr[2:3]         //переменной даем символы от arr от 2 до 3 символа
	return string(y)      //Функцие proverka даем значение y в расширении string
}

//Функция ввода слова
func Scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}

//Функция шифра
func shir() string {
	password := []byte("some")
	payload := []byte(source)
	encrypted, encrypt_error := lithcrypt.Encrypt(password, payload)
	if encrypt_error != nil {
		fmt.Println("Failed to encrypt:", encrypt_error)
		os.Exit(1)
	}
	return string(encrypted)
}

//Функция получения сообщения и расшифровка
func rshifr(conn net.Conn) string {
	password := []byte("some")
	//defer conn.Close()

	input := make([]byte, (1024))
	n, err := conn.Read(input)
	if n == 0 || err != nil {
		fmt.Println("Read error123:", err)
		os.Exit(1)
	}
	original, decrypt_error := lithcrypt.Decrypt(password, input[0:n])
	if decrypt_error != nil {
		fmt.Println("Failed to decrypt:", decrypt_error)
		os.Exit(1)
	}
	return string(original)
}

func main() {
	//Идет подключение к серверу
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		//Ввод логина
		fmt.Print("login: ")
		source = Scan1()
		shir := shir()
		//Отправляем серверу наш логин
		if n, err := conn.Write([]byte(shir)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}
		//Проверка на правильность логина
		proverka := rshifr(conn)
		if proverka == "errl" { //errl Ошибка логина
			fmt.Println("Логин введен не верно\nПовторите попытку: ")
		} else {
			break
		}
	}
	for {
		//Ввод пароля
		fmt.Print("Password: ")
		source = Scan1()
		shir := shir()
		if n, err := conn.Write([]byte(shir)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}
		proverka := rshifr(conn)
		//Проверка пароля
		if proverka == "errp" { //errl Ошибка логина
			fmt.Println("Пароль введен не верно\nПовторите попытку: ")
		} else {
			break
		}
	}

	for {
		fmt.Println("Вы успешно автоватезировались!")
		fmt.Println("Чтобы узнать список команд напишите /help")
		//Ввод слова с клавиатуры
		source = Scan1()
		//Сразу приравниваем функции к переменным
		comand := comand()
		id := id()
		text := text()
		probel := probel()

		if comand == "/h" {
			fmt.Println("\t\tСписок команд:\n/list-Посмотреть список клиентов онлайн\n/write [id] [текст] - напиать пользователю\n/back - Выйти из диолога с собеседником\n/exit - Выход из программы.")
		}
		if comand == "/l" {
			source = comand
			encrypted := []byte(shir())
			if n, err := conn.Write(encrypted); n == 0 || err != nil {
				fmt.Println(err)
				return
			}
		}

		otvet := rshifr(conn)
		fmt.Println(otvet)
		if comand == "/w" {
			switch {
			case id == "1":
				for text == "/back" {
					fmt.Println("Вы вошли в диолог с пользователем ", id)
					source = Scan1()
					message := comand + probel + id + probel + text
					source = message
					encrypted := []byte(shir())
					if n, err := conn.Write(encrypted); n == 0 || err != nil {
						fmt.Println(err)
						return
					}
				}
			case id == "2":
				for text == "/back" {
					fmt.Println("Вы вошли в диолог с пользователем ", id)
					source = Scan1()
					message := comand + probel + id + probel + text
					source = message
					encrypted := []byte(shir())
					if n, err := conn.Write(encrypted); n == 0 || err != nil {
						fmt.Println(err)
						return
					}
				}
			case id == "3":
				for text == "/back" {
					fmt.Println("Вы вошли в диолог с пользователем ", id)
					source = Scan1()
					message := comand + probel + id + probel + text
					source = message
					encrypted := []byte(shir())
					if n, err := conn.Write(encrypted); n == 0 || err != nil {
						fmt.Println(err)
						return
					}
				}
			case id == "4":
				for text == "/back" {
					fmt.Println("Вы вошли в диолог с пользователем ", id)
					source = Scan1()
					message := comand + probel + id + probel + text
					source = message
					encrypted := []byte(shir())
					if n, err := conn.Write(encrypted); n == 0 || err != nil {
						fmt.Println(err)
						return
					}
				}
			case id == "5":
				for text == "/back" {
					fmt.Println("Вы вошли в диолог с пользователем ", id)
					source = Scan1()
					message := comand + probel + id + probel + text
					source = message
					encrypted := []byte(shir())
					if n, err := conn.Write(encrypted); n == 0 || err != nil {
						fmt.Println(err)
						return
					}
				}
			}
		}

		if comand == "/e" {
			os.Exit(1)
		} else {
			fmt.Println("Введите сначало команду\nДля просмотра команд введите /help")
		}
	}
}
