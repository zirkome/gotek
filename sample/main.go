package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"github.com/kokaz/gotek"
	"golang.org/x/crypto/ssh/terminal"
)

func readLine(prompt string) (input string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err = reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return string(input), nil
}

func readPassword(prompt string) (passwd string, err error) {
	fmt.Print(prompt)

	password, err := terminal.ReadPassword(syscall.Stdin)
	if err != nil {
		return "", err
	}
	fmt.Print("\n")

	return string(password), nil
}

func main() {
	login, _ := readLine("Login: ")
	password, _ := readPassword("Password: ")
	client := &gotek.Client{Login: login, Password: password}

	dashboard, err := client.SignIn()
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println(dashboard.Infos.ID)

	schedule, err := client.GetSchedule()
	if err != nil {
		fmt.Printf(err.Error())
	}
	for _, event := range schedule {
		if len(event.Scolaryear) > 0 {
			fmt.Printf("%#v\n", event)
		}
	}
}
