package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const configFileName = "gswitch_config.json"

type UserConfig struct {
	Users []GitUser `json:"users"`
}

type GitUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func LoadConfig() (UserConfig, error) {
	config := UserConfig{}
	configPath := getConfigFilePath()
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return config, nil
		}
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func SaveConfig(config UserConfig) {
	configPath := getConfigFilePath()

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling config data:", err)
		os.Exit(1)
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		fmt.Println("Error writing to config file:", err)
		os.Exit(1)
	}
}

func getConfigFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		os.Exit(1)
	}

	return filepath.Join(homeDir, configFileName)
}

func GetActiveGitUser() (string, string, error) {
	getUsernameCmd := exec.Command("git", "config", "--get", "user.name")
	username, err := getUsernameCmd.Output()
	if err != nil {
		return "", "", err
	}
	getEmailCmd := exec.Command("git", "config", "--get", "user.email")
	email, err := getEmailCmd.Output()
	if err != nil {
		return "", "", err
	}
	return strings.TrimSpace(string(username)), strings.TrimSpace(string(email)), nil
}
