# Migration [MySQL]
Simple migration script wrote in golang, inspired in laravel migration implementation

# How Use
* Rename config.env.example to config.env and edit with the DataBase info
* Create .sql script files with the next format: YYYY_MM_YY_PRIORITY_NAME.sql

# Example: 
* Create a folder for your sql migration files example: **sqlFiles/** and put all yours **.sql** files inside
```
0000_00_00_0000_create_migrations_table.sql   // This script should be the fist sql executed
2010_01_01_0000_squema.sql                    // Your initial squema
2020_11_03_0001_create_folder_table.sql       // A migration script file
2020_11_03_0002_alter_text_table.sql          // Another file with migration scripts
2021_01_03_0001_insert_in_users_table.sql     // Another file with migration scripts
```

* Rename **config.env.example** to **config.env** and set your project values
* Execute migration.exe (windows) or the specific binary for (linux, mac).
* The app will generate a log file with the scripts executed.
* Please, check out **example/** folder for more detals.

# Dev - Build Binaries
```
GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -o bin/migration-windows-amd64.exe src/app/main.go;
GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/migration-linux-amd64 src/app/main.go;
GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s" -o bin/migration-mac-amd64 src/app/main.go;
GOOS=darwin GOARCH=arm64 go build -ldflags="-w -s" -o bin/migration-mac-arm64 src/app/main.go;
lipo -create -output bin/migration-mac-universal bin/migration-mac-amd64 bin/migration-mac-arm64;
rm bin/migration-mac-amd64; 
rm bin/migration-mac-arm64;
```


