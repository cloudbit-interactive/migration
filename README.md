# migration [MySQL]
Simple migration script wrote in golang, inspired in laravel migration implementation

HOW USE
* Open and edit config.env with DataBase info
* Create .sql script files with the next format: YYYY_MM_YY_PRIORITY_NAME.sql
````
EXAMPLE: 

Files
0000_00_00_0000_create_migrations_tablet.sql // This script should be the fist sql executed
2020_05_01_0002_squema.sql
2020_05_01_0002_create_table1.sql
2020_07_12_0000_update_table1.sql
````
* Execute migration.exe (windows) or migration (linux)
* The app will generate a log file with the scripts executed