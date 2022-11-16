package main

import (
	"bufio"
	"fmt"
	"github.com/cloudbit-interactive/cuppago"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

var DB cuppago.DataBase

func main() {
	configPath := cuppago.GetRootPath() + "/config.env"
	if godotenv.Load(configPath) != nil {
		cuppago.Error("Error loading [" + configPath + "] file")
	}

	DB = cuppago.NewDataBase(os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"))
	files, _ := ioutil.ReadDir(GetFilePath())

	if os.Getenv("FILES_SORT") == "UPDATE_DATE_ASC" {
		sort.Slice(files, func(i, j int) bool { return files[i].ModTime().After(files[j].ModTime()) })
	} else if os.Getenv("FILES_SORT") == "UPDATE_DATE_DESC" {
		sort.Slice(files, func(i, j int) bool { return files[i].ModTime().Before(files[j].ModTime()) })
	}

	for _, file := range files {
		if !strings.Contains(file.Name(), ".sql") {
			continue
		}
		data := DB.GetRow("migrations", "migration = '"+file.Name()+"'", "", "")
		if cuppago.Value(data, "migration", "") != "" {
			continue
		}
		ImportFile(file)
	}

	if strings.TrimSpace(os.Getenv("EXIT")) == "true" {
		return
	} else {
		fmt.Print("Press [Enter] to exit...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}

func ImportFile(file os.FileInfo) {
	fileData, err := ioutil.ReadFile(GetFilePath() + file.Name())
	if err != nil {
		cuppago.Error(err)
	}
	scripts := strings.Split(string(fileData), ";")
	cuppago.LogFile("--- " + file.Name() + " ---")
	for _, script := range scripts {
		if strings.TrimSpace(script) == "" {
			continue
		}
		parts := strings.Split(script, "\r\n")
		script := strings.TrimSpace(strings.Join(parts, " "))
		DB.SQL(script)
	}
	DB.SQL("INSERT INTO `migrations` (`migration`) VALUES ('" + file.Name() + "')")
}

func GetFilePath() string {
	filePath := cuppago.GetRootPath()
	cuppago.Log(filePath)
	if os.Getenv("FILES_PATH") != "" {
		filePath += "/" + os.Getenv("FILES_PATH")
	}
	return filePath
}
