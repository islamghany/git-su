package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/islamghany/git-su/fileio"
)

// //////////////////////////// Types ////////////////////////////////
type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
type Options struct {
	id    string
	email string
	name  string
}
type UsersKeys map[string]User

// //////////////////////////// Global Variables ////////////////////////////////
const usersFileName = ".git-su-users.json"

var errNotFound error = errors.New("file not found")
var errUserIdNotFound error = errors.New("user id not found")
var fileHanlder *fileio.FileIO
var users UsersKeys
var options Options

// //////////////////////////// Helper Functions ////////////////////////////////
// parseArgs parses the command line arguments and sets the options
func parseArgs() error {
	for _, arg := range os.Args[2:] {
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("error: Invalid argument format: %s", arg)

		}
		flagName := parts[0]
		flagValue := parts[1]
		switch flagName {
		case "-id":
			options.id = flagValue
		case "-email":
			options.email = flagValue
		case "-name":
			options.name = flagValue
		default:
			continue

		}
	}
	return nil
}

// getExistedUsers reads the users from the file and sets the users global variable
func getExistedUsers() error {
	usersBytes, err := fileHanlder.ReadFromFile(usersFileName)
	if err != nil {
		return errNotFound
	}
	err = json.Unmarshal(usersBytes, &users)
	if err != nil {
		return err
	}
	return nil

}

// addUser adds a new user to the users global variable
func addUser(id string, email string, name string) {
	users[id] = User{
		Email: email,
		Name:  name,
	}
}

// removeUser removes a user from the users global variable
func removeUser(id string) {
	_, ok := users[id]
	if ok {
		delete(users, id)
	}
}

// persistUsers writes the users global variable to the file
func persistUsers() error {
	usersBytes, err := json.Marshal(users)
	if err != nil {
		return err
	}
	err = fileHanlder.WriteToFile(usersFileName, usersBytes)
	if err != nil {
		return err
	}
	return nil
}

////////////////////////////// Main Functions ////////////////////////////////

// git-su add -id=<id> -email=<email> -name=<name>
func handleAddNewUser() error {
	if options.id == "" {
		return errors.New("id is required")
	}
	if options.email == "" {
		return errors.New("email is required")
	}
	if options.name == "" {
		return errors.New("name is required")
	}
	addUser(options.id, options.email, options.name)
	err := persistUsers()
	if err != nil {
		return err
	}
	return nil
}

// git-su remove -id=<id> : or <rm> Remove the user with the given id
func handleRemoveUser() error {
	if options.id == "" {
		return errors.New("id is required")
	}
	removeUser(options.id)
	err := persistUsers()
	if err != nil {
		return err
	}
	return nil
}

// git-su list
func handleListUsers() {
	for id, user := range users {
		fmt.Printf("%s: name=%s email=%s\n", id, user.Name, user.Email)
	}
}

// git-su <id>
func handleSwitchUser() error {
	if options.id == "" {
		return errors.New("id is required")
	}
	user, ok := users[options.id]
	if !ok {
		return errUserIdNotFound
	}
	var err error
	if user.Email != "" {
		cmd := exec.Command("git", "config", "--global", "user.email", user.Email)
		output, err := cmd.Output()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
	}
	if user.Name != "" && err == nil {
		cmd := exec.Command("git", "config", "--global", "user.name", user.Name)
		output, err := cmd.Output()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
	}
	return nil
}

// git-su which
func handleWhichUser() {
	cmd := exec.Command("git", "config", "--global", "--list")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}

// git-su -h
func handleHelp() {
	fmt.Println("Usage:")
	fmt.Println("git-su <id> : Switch to the user with the given id")
	fmt.Println("git-su which : Show the current user information")
	fmt.Println("git-su add -id=<id> -email=<email> -name=<name> : Add a new user")
	fmt.Println("git-su remove -id=<id> : or <rm> Remove the user with the given id")
	fmt.Println("git-su list : or <ls> List all users")
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Error: Not enough arguments provided")
		return
	}
	command := os.Args[1]
	err := parseArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = getExistedUsers()
	if err != nil {
		if err == errNotFound {
			users = UsersKeys{}
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	var eventialError error
	fileHanlder = fileio.NewFileIO()
	switch command {
	case "add":
		eventialError = handleAddNewUser()
	case "remove":
	case "rm":
		eventialError = handleRemoveUser()
	case "list":
	case "ls":
		handleListUsers()
	case "-h":
	case "--help":
		handleHelp()
	case "which":
		handleWhichUser()
	default:
		if len(os.Args) == 2 {
			options.id = command
			eventialError = handleSwitchUser()
			if eventialError == nil {
				fmt.Println("Switched to user with id:", command)
			}
		} else {
			fmt.Println("Error: Invalid command")
		}
	}
	if eventialError != nil {
		fmt.Println(eventialError)
	}
}
