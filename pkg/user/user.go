package user

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/nabadeep25/gswitch/pkg/helper"
)

func AddUser(username, email string) {
	config, err := helper.LoadConfig()
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		os.Exit(1)
	}
	config.Users = append(config.Users, helper.GitUser{Username: username, Email: email})
	helper.SaveConfig(config)

	fmt.Printf("Git user added: %s <%s>\n", username, email)
}

func ListUsers() helper.UserConfig {
	config, err := helper.LoadConfig()
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		os.Exit(1)
	}

	if len(config.Users) == 0 {
		fmt.Println("No Git users found. Add a user using 'gswitch add' first.")
		os.Exit(1)
	}

	fmt.Println("Git users:")
	for i, user := range config.Users {
		fmt.Printf("%d: %s <%s>\n", i+1, user.Username, user.Email)
	}
	return config
}
func SelectUser() {
	config := ListUsers()
	var selectedIndex int
	fmt.Print("Enter the user key to select: ")
	fmt.Scanln(&selectedIndex)
	if selectedIndex < 1 || selectedIndex > len(config.Users) {
		fmt.Println("Invalid user index.")
		os.Exit(1)
	}
	user := config.Users[selectedIndex-1]
	cmd := exec.Command("git", "config", "--global", "user.name", user.Username)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error setting active Git user:", err)
		os.Exit(1)
	}
	cmd = exec.Command("git", "config", "--global", "user.email", user.Email)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error setting active Git user:", err)
		os.Exit(1)
	}
	fmt.Printf("Active Git user set to: %s <%s>\n", user.Username, user.Email)
}

func DeleteUser() {
	config := ListUsers()
	var selectedIndex int
	fmt.Print("Enter the user key to remove: ")
	fmt.Scanln(&selectedIndex)
	if selectedIndex < 1 || selectedIndex > len(config.Users) {
		fmt.Println("Invalid user index.")
		os.Exit(1)
	}
	deletedUser := config.Users[selectedIndex-1]
	config.Users = append(config.Users[:selectedIndex-1], config.Users[selectedIndex:]...)
	helper.SaveConfig(config)
	fmt.Printf("Removed user: %s <%s>\n", deletedUser.Username, deletedUser.Email)
}
