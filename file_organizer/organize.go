package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	numOfFiles := map[string]int{
		"audio":     0,
		"image":    0,
		"video":    0,
		"document": 0,
		"other":    0,
	}
	folderPath := os.Args[1]

	organizeFiles(folderPath, numOfFiles)
	displayNumberOfFiles(numOfFiles)
}

func getFolderItems(path string) []fs.DirEntry {
	info, err := os.Stat(path)
	if err != nil {
		log.Fatalln(err)
	}

	if !info.IsDir() {
		log.Fatalln("Please provide a path to a folder")
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}

	return entries
}

func createFolder(path string) {
	if err := os.Mkdir(path, 07555); err != nil {
		log.Fatalln(err)
	}
}

func moveFile(fileInfo fs.FileInfo, folderPath string) {
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		createFolder(folderPath)
	}

	oldPath := filepath.Join(filepath.Dir(folderPath), fileInfo.Name())
	newPath := filepath.Join(folderPath, fileInfo.Name())
	if err := os.Rename(oldPath, newPath); err != nil {
		log.Fatalln(err)
	}
}

func organizeFiles(folderPath string, numOfFiles map[string]int) {
	folderItems := getFolderItems(folderPath)

	for _, folderItem := range folderItems {
		if !folderItem.IsDir() {
			fileInfo, err := folderItem.Info()
			if err != nil {
				log.Fatalln(err)
			}

			fileExt := filepath.Ext(fileInfo.Name())

			switch fileExt {
			case ".wav", ".mp3", "aac":
				subFolderPath := filepath.Join(folderPath, "audio")
				moveFile(fileInfo, subFolderPath)
				numOfFiles["audio"] += 1
			case ".jpg", ".jpeg", ".png", ".svg", ".gif", ".webp", ".bmp":
				subFolderPath := filepath.Join(folderPath, "image")
				moveFile(fileInfo, subFolderPath)
				numOfFiles["image"] += 1
			case ".mp4", ".avi", ".mov", ".mkv", ".flv", ".amv":
				subFolderPath := filepath.Join(folderPath, "video")
				moveFile(fileInfo, subFolderPath)
				numOfFiles["video"] += 1
			case ".doc", ".docx", ".xlsx", ".csv", ".txt", ".pdf", ".epub", ".odt":
				subFolderPath := filepath.Join(folderPath, "document")
				moveFile(fileInfo, subFolderPath)
				numOfFiles["document"] += 1
			default:
				subFolderPath := filepath.Join(folderPath, "other")
				moveFile(fileInfo, subFolderPath)
				numOfFiles["other"] += 1
			}
		}
	}
}

func displayNumberOfFiles(numOfFiles map[string]int) {
	for category, count := range numOfFiles {
		if count > 0 {
			fmt.Printf("%d %s files moved\n", count, category)
		} else {
			fmt.Printf("No %s files moved\n", category)
		}
	}
}
