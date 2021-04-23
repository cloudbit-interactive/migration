package main

import (
	"com/cuppa/utils"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)
var DB utils.DataBase

func main(){
	if godotenv.Load(utils.GetRootPath()+"./config.env") != nil { utils.Error("Error loading config.env file") }
	DB = utils.NewDataBase(os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"))
	filePath := utils.GetRootPath()
	if os.Getenv("FILES_PATH") != "" { filePath += "/"+os.Getenv("FILES_PATH") }
	files, _ := ioutil.ReadDir(filePath)

	if os.Getenv("FILES_SORT") == "UPDATE_DATE_ASC" {
		sort.Slice(files, func(i,j int) bool{ return files[i].ModTime().After(files[j].ModTime()) })
	}else if os.Getenv("FILES_SORT") == "UPDATE_DATE_DESC" {
		sort.Slice(files, func(i,j int) bool{ return files[i].ModTime().Before(files[j].ModTime()) })
	}

	for _, file := range files {
		if !strings.Contains(file.Name(), ".sql") { continue }
		data := DB.GetRow("migrations", "migration = '"+file.Name()+"'", "", "")
		if utils.Value(data, "migration", "") != "" { continue }
		ImportFile(file)
	}
}

func ImportFile(file os.FileInfo){
	fileData, _ := ioutil.ReadFile(file.Name())
	scripts := strings.Split(string(fileData), ";")
	utils.LogFile("--- "+file.Name()+" ---")
	for _, script := range scripts {
		if strings.TrimSpace(script) == "" { continue }
		parts := strings.Split(script, "\r\n")
		script := strings.TrimSpace(strings.Join(parts, " "))
		DB.SQL(script)
		utils.LogFile(script)
	}
	DB.SQL("INSERT INTO `migrations` (`migration`) VALUES ('"+file.Name()+"')")
}

