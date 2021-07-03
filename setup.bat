@echo OFF
ECHO Setting up your project

COPY "env.example" ".env"

TYPE .env

ECHO "\n"
ECHO Running migration
go run cmd/database/migrate.go

CHOICE /M "Do you want to run seeding users?"
IF %ERRORLEVEL% EQU 2 goto items

ECHO Seeding users
go run cmd/database/migrate.go seed users

:items
CHOICE /M "Do you want to run seeding items?"
IF %ERRORLEVEL% EQU 2 goto end
ECHO Seeding items
go run cmd/database/migrate.go seed items

:end
ECHO Done ! Run go run main.go to start