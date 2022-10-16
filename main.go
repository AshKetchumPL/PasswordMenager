package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// TODO
// - Add support for Encryption

var (
	version = "0.0.0"
	dev     = false
)

type Password struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Domain   string `json:"domain"`
	Password string `json:"password"`
}

type commandFunc func() int

type Command struct {
	Name        string
	Description string
	Func        commandFunc
	Implemented bool // idk how to implement without this
}

var (
	commands  []Command
	passwords []Password
)

func init() {
	commands = append(commands, Command{
		Name:        "help",
		Description: "Shows this",
		Func: func() int {
			Red, Green, Cyan, Yellow := color.New(color.FgRed), color.New(color.FgGreen), color.New(color.FgCyan), color.New(color.FgYellow)

			for i := 0; i < len(commands); i++ {
				Red.Print(commands[i].Name)
				Cyan.Print(" - ")
				if !commands[i].Implemented {
					Yellow.Print("Not Implemented")
					Cyan.Print(" - ")
				}
				Green.Print(commands[i].Description)
				Green.Println()
			}
			return 0
		},
		Implemented: true,
	})
	commands = append(commands, Command{
		Name:        "stats",
		Description: "Shows statistics",
		Func: func() int {
			Cyan := color.New(color.FgCyan)
			Cyan.Println("Passwords > " + strconv.Itoa(len(passwords)))
			Cyan.Println("Commands  > " + strconv.Itoa(len(commands)))
			return 0
		},
		Implemented: true,
	})
	commands = append(commands, Command{
		Name:        "add",
		Description: "Add a new password",
		Func: func() int {
			Red, Green, Cyan := color.New(color.FgRed), color.New(color.FgGreen), color.New(color.FgCyan)
			var name string
			var login string
			var domain string
			var password string
			Green.Print("Enter Name > ")
			fmt.Scanln(&name)
			Cyan.Print("Enter Login > ")
			fmt.Scanln(&login)
			Cyan.Print("Enter Domain > ")
			fmt.Scanln(&domain)
			Red.Print("Enter Password > ")
			fmt.Scanln(&password)
			passwords = append(passwords, Password{
				Name:     name,
				Login:    login,
				Domain:   domain,
				Password: password,
			})
			Cyan.Println("Password saved in ram (type \"save\" to save in file)")

			return 0
		},
		Implemented: true,
	})
	commands = append(commands, Command{
		Name:        "save",
		Description: "Save added/removes deleted passwords",
		Func: func() int {
			Green := color.New(color.FgGreen)
			Green.Println("Saving...")
			save()
			Green.Println("Saved!")
			return 0
		},
		Implemented: true,
	})
	commands = append(commands, Command{
		Name:        "search",
		Description: "Search for a password",
		Func: func() int {
			Red, Green, Cyan := color.New(color.FgRed), color.New(color.FgGreen), color.New(color.FgCyan)
			var query string
			var found bool
			var foundPassword Password
			var foundPasswordNo int
			Green.Print("Qery > ")
			fmt.Scanln(&query)
			for i := 0; i < len(passwords); i++ {
				if strings.ToLower(query) == strings.ToLower(passwords[i].Name) {
					found = true
					foundPasswordNo = i
					foundPassword = passwords[i]
					break
				}
			}
			for i := 0; i < len(passwords); i++ {
				if strings.ToLower(query) == strings.ToLower(passwords[i].Domain) {
					found = true
					foundPasswordNo = i
					foundPassword = passwords[i]
					break
				}
			}
			for i := 0; i < len(passwords); i++ {
				if strings.ToLower(query) == strings.ToLower(passwords[i].Login) {
					found = true
					foundPasswordNo = i
					foundPassword = passwords[i]
					break
				}
			}
			if found {
				Green.Println(strconv.Itoa(foundPasswordNo) + " {")
				Cyan.Println("  Name     > " + foundPassword.Name)
				Cyan.Println("  Login    > " + foundPassword.Login)
				Cyan.Println("  Domain   > " + foundPassword.Domain)
				Red.Println("  Password > " + foundPassword.Password)
				Green.Println("}")
				Green.Println()
				return 0
			}
			Red.Println("Password Not found!")
			return 0
		},
		Implemented: true,
	})
	commands = append(commands, Command{
		Name:        "listall",
		Description: "List all passwords (VERY DANGEROUS)",
		Func: func() int {
			Red, Green, Cyan := color.New(color.FgRed), color.New(color.FgGreen), color.New(color.FgCyan)
			var sure string
			Green.Print("Are You Sure? Y/N > ")
			fmt.Scanln(&sure)
			if strings.ToLower(sure) == "y" {
				for i := 0; i < len(passwords); i++ {
					Green.Println(strconv.Itoa(i) + " {")
					Cyan.Println("  Name     > " + passwords[i].Name)
					Cyan.Println("  Login    > " + passwords[i].Login)
					Cyan.Println("  Domain   > " + passwords[i].Domain)
					Red.Println("  Password > " + passwords[i].Password)
					Green.Println("}")
					Green.Println()
				}
			}
			return 0
		},
		Implemented: true,
	})
	commands = append(commands, Command{
		Name:        "about",
		Description: "About this program",
		Func: func() int {
			Red, Green, Cyan := color.New(color.FgRed), color.New(color.FgGreen), color.New(color.FgCyan)

			Green.Print("Github: ")
			Cyan.Print("github.com/AshKetchumPL/")
			Red.Println("PasswordMenager")

			fmt.Println()

			Green.Print("version: ")
			Cyan.Print(version)
			if dev {
				Cyan.Print("-dev")
			}

			fmt.Println()

			Green.Print("Program: ")
			Cyan.Println("This Password Menager was created for fun just to get some knowledge of programing")

			Green.Print("Made with: ")
			Cyan.Print("sweat, tears")
			Green.Print(" and ")
			Cyan.Print("go")
			Green.Println(" (language)")

			fmt.Println()

			Green.Print("Author: ")
			Cyan.Println("AshKetchumPL")
			Green.Print("Discord: ")
			Cyan.Println(">>ash ketchum<<#6595")
			return 0
		},
		Implemented: true,
	})
	commands = append(commands, Command{
		Name:        "exit",
		Description: "Exits this program",
		Func: func() int {
			os.Exit(0)
			return 0
		},
		Implemented: true,
	})
}

func main() {
	Red, Green, Cyan, Yellow := color.New(color.FgRed), color.New(color.FgGreen), color.New(color.FgCyan), color.New(color.FgYellow)

	Green.Print("# ")
	Cyan.Print("github.com/AshKetchumPL/")
	Red.Println("PasswordMenager")

	if _, err := os.Stat("./Password Menager"); os.IsNotExist(err) {
		setup()
	}
	jsonFile, err := os.Open("./Password Menager/passwords.apm")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &passwords)

	Yellow.Println("Warning! The Passwords are not encrypted in any way please dont save your real passwords here.")
	Cyan.Println("Type \"help\" for help")

	for {
		var inpt string
		var found bool
		Green.Print("> ")
		fmt.Scanln(&inpt)
		for i := 0; i < len(commands); i++ {
			if strings.ToLower(commands[i].Name) == inpt {
				if commands[i].Implemented {
					commands[i].Func()
				} else {
					Red.Println("Command not Implemented: ", inpt)
				}

				found = true
				break
			}
		}
		if !found {
			Red.Println("Command not found: ", inpt)
		}

	}
}

func setup() {
	Red, Green, Cyan := color.New(color.FgRed), color.New(color.FgGreen), color.New(color.FgCyan)
	Red.Print("Password Menager Folder was not found. ")
	Cyan.Println("Creating Password Menager Folder (starting setup)")
	err := os.Mkdir("./Password Menager", 0755)
	if err != nil {
		log.Fatal(err)
	}
	data := []Password{}
	file, _ := json.MarshalIndent(data, "", "	")
	_ = ioutil.WriteFile("./Password Menager/passwords.apm", file, 0644)
	Green.Println("Setup Complete")
}

func save() {
	file, _ := json.MarshalIndent(passwords, "", "	")
	_ = ioutil.WriteFile("./Password Menager/passwords.apm", file, 0644)
}
