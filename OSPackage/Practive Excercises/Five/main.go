package main

import (
	"fmt"
	"os"
)

func CheckDirectories(path string) (result bool, er error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	if !info.IsDir() {
		return false, fmt.Errorf("%s is not a directory", path)
	}
	return true, nil
}

func CheckFileExistenceAndAccessibility(filePath string) (
	isExist bool,
	accessMode os.FileMode,
	er error,
) {
	info, err := os.Stat(filePath)
	if err != nil {
		isExist = false
		accessMode = 0
		er = err
		return isExist, accessMode, er
	}

	isExist = true
	accessMode = info.Mode()
	er = nil
	return isExist, accessMode, er
}

func AskUserForNextStep(question string) (answer string, result bool) {
	fmt.Println(question)

	_, _ = fmt.Scanln(&answer)

	if answer == "yes" || answer == "y" ||
		answer == "YES" || answer == "Y" ||
		answer == "Yes" {
		return answer, true
	}

	return answer, false
}

func main() {
	const (
		workDir = "workspace"
		dataDir = "workspace/data"
		srcPath = "workspace/data/settings.json"
		bkpPath = "workspace/data/settings.backup.json"
	)

	ok, err := CheckDirectories(workDir)
	if !ok {
		fmt.Println("workspace not found:", err)
		if _, create := AskUserForNextStep("Create 'workspace' directory?"); create {
			if mkErr := os.MkdirAll(workDir, 0755); mkErr != nil {
				fmt.Println("Failed to create workspace:", mkErr)
				return
			}
			fmt.Println("Created:", workDir)
		} else {
			fmt.Println("Aborted: workspace directory is required.")
			return
		}
	}

	ok, err = CheckDirectories(dataDir)
	if !ok {
		fmt.Println("data dir not found:", err)
		if _, create := AskUserForNextStep("Create 'workspace/data' directory?"); create {
			if mkErr := os.MkdirAll(dataDir, 0755); mkErr != nil {
				fmt.Println("Failed to create data dir:", mkErr)
				return
			}
			fmt.Println("Created:", dataDir)
		} else {
			fmt.Println("Aborted: data directory is required.")
			return
		}
	}

	srcExists, _, srcErr := CheckFileExistenceAndAccessibility(srcPath)
	if !srcExists {
		fmt.Println("settings.json not found:", srcErr)
		if _, create := AskUserForNextStep("Create a default 'settings.json'?"); create {
			content := []byte("{\n  \"setting\": \"default\"\n}\n")
			if wErr := os.WriteFile(srcPath, content, 0644); wErr != nil {
				fmt.Println("Failed to create settings.json:", wErr)
				return
			}
			fmt.Println("Created:", srcPath)
		} else {
			fmt.Println("Nothing to back up. Continuing with temp dir demo.")
		}
	}

	if info, statErr := os.Stat(srcPath); statErr == nil && !info.IsDir() {
		f, openErr := os.Open(srcPath)
		if openErr != nil {
			fmt.Println("settings.json exists but cannot be read:", openErr)
		} else {
			f.Close()
			fmt.Println("settings.json is readable.")

			doBackup := true
			if _, bErr := os.Stat(bkpPath); bErr == nil {
				if _, overwrite := AskUserForNextStep(
					"Backup 'settings.backup.json' already exists. Overwrite it?",
				); overwrite {
					if rmErr := os.Remove(bkpPath); rmErr != nil {
						fmt.Println("Failed to remove existing backup:", rmErr)
						doBackup = false
					}
				} else {
					fmt.Println("Error: backup already exists and was not overwritten.")
					doBackup = false
				}
			}

			if doBackup {
				if rnErr := os.Rename(srcPath, bkpPath); rnErr != nil {
					fmt.Println("Failed to rename settings.json:", rnErr)
				} else {
					fmt.Println("Renamed:", srcPath, "->", bkpPath)
				}
			}
		}
	}

	tempDir, mkErr := os.MkdirTemp("", "settings-work-")
	if mkErr != nil {
		fmt.Println("Failed to create temp dir:", mkErr)
		return
	}
	defer os.RemoveAll(tempDir)

	fmt.Println("Temporary working directory:", tempDir)
	fmt.Println("Done.")
}
