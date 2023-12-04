package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	folderPath := os.Args[1]

	organizeFiles(folderPath)
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

func organizeFiles(folderPath string) {
	folderItems := getFolderItems(folderPath)
	numOfFiles := map[string]int{
		"audio":     0,
		"image":    0,
		"video":    0,
		"document": 0,
		"compressed": 0,
		"other":    0,
	}

	for _, folderItem := range folderItems {
		if !folderItem.IsDir() {
			fileInfo, err := folderItem.Info()
			if err != nil {
				log.Fatalln(err)
			}

			fileExt := filepath.Ext(fileInfo.Name())

			switch fileExt {
			case ".wav", ".mp3", "aac", ".wma", ".m4a":
				subFolderPath := filepath.Join(folderPath, "audio")
				moveFile(fileInfo, subFolderPath)
				numOfFiles["audio"] += 1
			case ".jpg", ".jpeg", ".png", ".svg", ".gif", ".webp", ".bmp":
				subFolderPath := filepath.Join(folderPath, "image")
				moveFile(fileInfo, subFolderPath)
				numOfFiles["image"] += 1
			case ".mp4", ".avi", ".mov", ".mkv", ".flv", ".amv", ".wmv", ".m4v":
				subFolderPath := filepath.Join(folderPath, "video")
				moveFile(fileInfo, subFolderPath)
				numOfFiles["video"] += 1
			case ".doc", ".docx", ".ppt", ".pptx", ".xlsx", ".csv", ".txt", ".rtf", ".pdf", ".epub", ".odt":
				subFolderPath := filepath.Join(folderPath, "document")
				moveFile(fileInfo, subFolderPath)
				numOfFiles["document"] += 1
			case ".zip", ".rar", ".pkg":
				subFolderPath := filepath.Join(folderPath, "compressed")
				moveFile(fileInfo, subFolderPath)
				numOfFiles["compressed"] += 1
			default:
				subFolderPath := filepath.Join(folderPath, "other")
				moveFile(fileInfo, subFolderPath)
				numOfFiles["other"] += 1
			}
		}
	}

	displayNumOfFiles(numOfFiles)
}

func displayNumOfFiles(numOfFiles map[string]int) {
	for category, count := range numOfFiles {
		if count > 0 {
			fmt.Printf("%d %s files moved\n", count, category)
		} else {
			fmt.Printf("No %s files moved\n", category)
		}
	}
}
