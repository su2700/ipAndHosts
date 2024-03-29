package main

import (
        "bufio"
        "fmt"
        "os"
        "strings"
)

func main() {
        // Get user inputs
        fmt.Print("Enter IP address: ")
        var ipAddress string
        fmt.Scanln(&ipAddress)

        fmt.Print("Enter hostname: ")
        var hostname string
        fmt.Scanln(&hostname)

        // Read the existing /etc/hosts file
        hostsFile, err := os.OpenFile("/etc/hosts", os.O_RDWR|os.O_APPEND, 0644)
        if err != nil {
                fmt.Println("Error opening /etc/hosts file:", err)
                return
        }
        defer hostsFile.Close()

        // Check if the entry already exists
        scanner := bufio.NewScanner(hostsFile)
        found := false
        for scanner.Scan() {
                line := scanner.Text()
                if strings.Contains(line, ipAddress) && strings.Contains(line, hostname) {
                        found = true
                        break
                }
        }

        // If the entry doesn't exist, append it to the file
        if !found {
                entry := fmt.Sprintf("%s\t%s\n", ipAddress, hostname)
                _, err = hostsFile.WriteString(entry)
                if err != nil {
                        fmt.Println("Error writing to /etc/hosts file:", err)
                        return
                }
                fmt.Println("Entry added to /etc/hosts file successfully.")
        } else {
                fmt.Println("Entry already exists in /etc/hosts file.")
        }
}