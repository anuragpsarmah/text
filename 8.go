package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type Info struct {
	Name       string `json:"Name"`
	Course     string `json:"Course"`
	University string `json:"University"`
}

func main() {
	for {
		fmt.Println("\nSelect an option:")
		fmt.Println("1. Marshaling (Save user input as JSON)")
		fmt.Println("2. Unmarshaling (Fetch JSON data from URL)")
		fmt.Println("3. Exit")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			marshalJSON()
		case "2":
			unmarshalJSON()
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func marshalJSON() {
	info := Info{}
	fmt.Print("Enter Name: ")
	fmt.Scanln(&info.Name)
	fmt.Print("Enter Course: ")
	fmt.Scanln(&info.Course)
	fmt.Print("Enter University: ")
	fmt.Scanln(&info.University)

	jsonData, err := json.MarshalIndent(info, "", " ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	file, err := os.Create("info.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("JSON data saved to info.txt")

	// Navigate to the repository directory
	repoDir := "C:/Users/anura/OneDrive/Desktop/Code/GO-Lang/"
	err = os.Chdir(repoDir)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}

	// Add the modified file to the staging area
	addCmd := exec.Command("git", "add", "info.txt")
	err = addCmd.Run()
	if err != nil {
		fmt.Println("Error adding file to staging area:", err)
		return
	}

	// Commit the changes
	commitMsg := "Updated info.txt"
	commitCmd := exec.Command("git", "commit", "-m", commitMsg)
	err = commitCmd.Run()
	if err != nil {
		fmt.Println("Error committing changes:", err)
		return
	}

	// Add the remote repository URL
	remoteCmd := exec.Command("git", "remote", "add", "origin", "https://github.com/anuragpsarmah/text.git")
	err = remoteCmd.Run()
	if err != nil {
		fmt.Println("Error adding remote repository:", err)
		return
	}

	// Push the changes to the remote repository
	pushCmd := exec.Command("git", "push", "origin", "main")
	err = pushCmd.Run()
	if err != nil {
		fmt.Println("Error pushing changes to remote repository:", err)
		return
	}

	fmt.Println("JSON data saved to info.txt and pushed to the remote repository.")
}

func unmarshalJSON() {
	url := "https://raw.githubusercontent.com/anuragpsarmah/text/main/info.txt"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var info Info
	err = json.Unmarshal(body, &info)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("\nExtracted information:")
	fmt.Println("Name:", info.Name)
	fmt.Println("Course:", info.Course)
	fmt.Println("University:", info.University)
}
