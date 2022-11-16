# migration [MySQL]
Simple migration script wrote in golang, inspired in laravel migration implementation

HOW USE
* Rename config.env.example to config.env and edit with the DataBase info
* Create .sql script files with the next format: YYYY_MM_YY_PRIORITY_NAME.sql
````
EXAMPLE: 

Files
0000_00_00_0000_create_migrations_table.sql // This script should be the fist sql executed
2010_01_01_0000_squema.sql
2020_11_03_0001_create_folder_table.sql
2020_11_03_0002_alter_text_table.sql
2021_01_03_0001_insert_in_users_table.sql
````
* Execute migration.exe (windows) or migration (linux).
* The app will generate a log file with the scripts executed.
* Please, check our example folder.

#DEV - BUILD FROM WINDOWS
````
FOR WIN
WINDOW
  GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -o bin/migration-win-amd64.exe src/app/main.go
LINUX
  GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/migration-linux-amd64 src/app/main.go
MAC
  GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s" -o bin/migration-mac-amd64 src/app/main.go
````


