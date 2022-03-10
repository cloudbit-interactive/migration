# migration [MySQL]
Simple migration script wrote in golang, inspired in laravel migration implementation

HOW USE
* Rename config.env.example to config.env and edit with the DataBase info
* Create .sql script files with the next format: YYYY_MM_YY_PRIORITY_NAME.sql
````
EXAMPLE: 

Files
0000_00_00_0000_create_migrations_tablet.sql // This script should be the fist sql executed
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
> cmd /C "SET GOOS=windows&&SET GOARCH=amd64&& go build -ldflags="-w -s" -o main.exe src/app/main.go"
FOR LINUX
> cmd /C "SET GOOS=linux&&SET GOARCH=amd64&& go build -ldflags="-w -s" -o src/app/main main.go"
````