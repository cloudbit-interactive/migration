package main

import (
	"bufio"
	"fmt"
	"github.com/cloudbit-interactive/cuppago"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
	"regexp"
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

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	if os.Getenv("FILES_SORT") == "UPDATE_DATE_ASC" {
		sort.Strings(fileNames)
	} else if os.Getenv("FILES_SORT") == "UPDATE_DATE_DESC" {
		sort.Sort(sort.Reverse(sort.StringSlice(fileNames)))
	}

	for _, fileName := range fileNames {
		if !strings.Contains(fileName, ".sql") {
			continue
		}
		data := DB.GetRow("migrations", "migration = '"+fileName+"'", "", "")
		if cuppago.Value(data, "migration", "") != "" {
			continue
		}
		ImportFile(fileName)
	}

	if strings.TrimSpace(os.Getenv("EXIT")) == "true" {
		return
	} else {
		fmt.Print("Press [Enter] to exit...")
		_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
		if err != nil {
			return
		}
	}
}

func ImportFile(fileName string) {
	fileData, err := ioutil.ReadFile(GetFilePath() + fileName)
	if err != nil {
		cuppago.Error(err)
	}
	var scripts []string
	if strings.Contains(fileName, "no_split") {
		scripts = []string{RemoveDelimiters(string(fileData))}
	} else {
		scripts = strings.Split(string(fileData), ";")
	}

	cuppago.LogFile("--- " + fileName + " ---")
	for _, script := range scripts {
		if strings.TrimSpace(script) == "" {
			continue
		}
		parts := strings.Split(script, "\r\n")
		script := strings.TrimSpace(strings.Join(parts, " "))
		DB.SQL(script)
	}
	DB.SQL("INSERT INTO `migrations` (`migration`) VALUES ('" + fileName + "')")
}

func GetFilePath() string {
	filePath := cuppago.GetRootPath()
	cuppago.Log(filePath)
	if os.Getenv("FILES_PATH") != "" {
		filePath += "/" + os.Getenv("FILES_PATH")
	}
	return filePath
}

func RemoveDelimiters(sql string) string {
	delimiterPattern := `(?i)DELIMITER\s+\S+`
	re := regexp.MustCompile(delimiterPattern)
	sql = re.ReplaceAllString(sql, "")
	sql = strings.ReplaceAll(sql, "//", "")
	return sql
}
