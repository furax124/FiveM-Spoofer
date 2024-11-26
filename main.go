package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var (
	fiveMPath      string
	subprocessPath string
)

func main() {
	cleanHWIDTraces()
	getPathsFromUser()
	printWelcomeMessage()
	monitorFiveM()
}

func getPathsFromUser() {
	fmt.Println("Enter the FiveM installation path (e.g., C:\\Program Files\\FiveM):")
	fmt.Scanln(&fiveMPath)
	fmt.Println("Enter the path to the FiveM subprocesses (e.g., C:\\Program Files\\FiveM\\FiveM.app\\data\\cache\\subprocess):")
	fmt.Scanln(&subprocessPath)
}

func printWelcomeMessage() {
	fmt.Println("*********************************")
	fmt.Println("* FiveM HWID Bypass - v1.0      *")
	fmt.Println("*********************************")
}

func isProcessRunning(processName string) bool {
	cmd := exec.Command("tasklist", "/FI", "IMAGENAME eq "+processName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing tasklist command:", err)
		return false
	}
	return strings.Contains(string(output), processName)
}

func detectGTAVersions() []string {
	var versions []string
	err := filepath.Walk(subprocessPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error reading path:", path, err)
			return err
		}
		if strings.Contains(info.Name(), "FiveM_") && strings.Contains(info.Name(), "GTAProcess.exe") {
			version := strings.Split(info.Name(), "_")[1]
			if !contains(versions, version) {
				versions = append(versions, version)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error detecting GTA versions:", err)
	}
	return versions
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func blockConnection(process string) {
	executeCommand(fmt.Sprintf("netsh advfirewall firewall add rule name=\"%s\" dir=out program=\"%s\" action=block", process, process))
	executeCommand(fmt.Sprintf("netsh advfirewall firewall add rule name=\"%s\" dir=in program=\"%s\" action=block", process, process))
}

func unblockConnection(process string) {
	executeCommand(fmt.Sprintf("netsh advfirewall firewall delete rule name=\"%s\" dir=out program=\"%s\"", process, process))
	executeCommand(fmt.Sprintf("netsh advfirewall firewall delete rule name=\"%s\" dir=in program=\"%s\"", process, process))
}

func executeCommand(command string) {
	fmt.Println("Executing:", command)
	err := exec.Command("cmd", "/C", command).Run()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
	}
}

func setupBypass() {
	blockConnection(filepath.Join(fiveMPath, "FiveM.exe"))
	for _, version := range detectGTAVersions() {
		blockConnection(filepath.Join(subprocessPath, fmt.Sprintf("FiveM_%sGTAProcess.exe", version)))
		blockConnection(filepath.Join(subprocessPath, fmt.Sprintf("FiveM_%sSteamChild.exe", version)))
	}
}

func removeBypass() {
	unblockConnection(filepath.Join(fiveMPath, "FiveM.exe"))
	for _, version := range detectGTAVersions() {
		unblockConnection(filepath.Join(subprocessPath, fmt.Sprintf("FiveM_%sGTAProcess.exe", version)))
		unblockConnection(filepath.Join(subprocessPath, fmt.Sprintf("FiveM_%sSteamChild.exe", version)))
	}
}

func cleanHWIDTraces() {
	directories := []string{
		filepath.Join(os.Getenv("LOCALAPPDATA"), "DigitalEntitlements"),
		filepath.Join(os.Getenv("APPDATA"), "CitizenFX"),
		filepath.Join(os.Getenv("LOCALAPPDATA"), "CrashDumps"),
	}

	registryKeys := []string{
		"HKCU\\Software\\CitizenFX",
		"HKCU\\Software\\Rockstar Games",
		"HKLM\\Software\\CitizenFX",
		"HKEY_LOCAL_MACHINE\SOFTWARE\FiveM",
	}

	fmt.Println("Starting HWID cleaning process...")

	for _, dir := range directories {
		err := os.RemoveAll(dir)
		if err != nil {
			fmt.Printf("Couldn't Find or Remove: %s\n", dir)
		} else {
			fmt.Printf("Removed: %s\n", dir)
		}
	}

	for _, key := range registryKeys {
		cmdCheck := exec.Command("cmd", "/C", fmt.Sprintf("reg query %s", key))
		if err := cmdCheck.Run(); err == nil {
			cmdDelete := exec.Command("cmd", "/C", fmt.Sprintf("reg delete %s /f", key))
			err := cmdDelete.Run()
			if err != nil {
				fmt.Printf("Error deleting registry key: %s, %v\n", key, err)
			} else {
				fmt.Printf("Deleted registry key: %s\n", key)
			}
		} else {
			fmt.Printf("Registry key not found: %s\n", key)
		}
	}

	fmt.Println("Resetting network settings...")
	commands := []string{
		"ipconfig /flushdns",
		"netsh int ip reset",
		"netsh advfirewall reset",
	}

	for _, command := range commands {
		cmd := exec.Command("cmd", "/C", command)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error executing command: %s, %v\n", command, err)
		} else {
			fmt.Printf("Executed: %s\n", command)
		}
	}

	fmt.Println("HWID cleaning process completed.")
}

func monitorFiveM() {
	for {
		if isProcessRunning("FiveM.exe") {
			fmt.Println("FiveM detected, starting bypass...")
			setupBypass()
			for isProcessRunning("FiveM.exe") {
				time.Sleep(1 * time.Second)
			}
			fmt.Println("FiveM closed - Removing network bypass...")
			removeBypass()
			fmt.Println("Cleaning up firewall rules...")
			cleanHWIDTraces()
		} else {
			fmt.Println("FiveM is not running. Waiting for it to start...")
			time.Sleep(1 * time.Second)
		}
	}
}
