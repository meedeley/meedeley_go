# Go Code StarterPack

## Copyright
For any open source project, there must be a LICENSE file in the repository root to claim the rights.

Here are two examples of using Apache License, Version 2.0 and MIT License.

### Apache License, Version 2.0
This license requires to put following content at the beginning of every file:
```
// Copyright [yyyy] [name of copyright owner]
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
```
Replace [yyyy] with the creation year of the file. Then use personal name for personal projects, or organization name for team projects to replace [name of copyright owner].
### MIT License
This license requires to put following content at the beginning of every file:
```
// Copyright [yyyy] [name of copyright owner]. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
```
Replace [yyyy] with the creation year of the file.
### Other Note
- Other types of license can follow the template of above two examples.

- If a file has been modified by different individuals and/or organizations, and multiple licenses are compatiable, then change the first line to multiple lines and sort them in the order of time:
## Project structure
```
bin/        # The directory for compiled binaries or executables
cmd/        # Contains the main applications for the project
db/         # Database migrations, seeds, or configurations
docs/       # Documentation files for the project
internal/   # Private application code
pkg/        # Code that's safe to be shared between different parts of the project
storage/    # Directory for file storage or static assets
tests/      # Test files and test utilities
```
### Additional directories and files:
- `.air.toml`: Configuration file for live-reloading with `air`.
- `.env`: Environment variables for the project.
- `.env-example`: A sample file to guide the creation of `.env`.
- `.gitignore`: Specifies files and directories to ignore in Git.
- `go.mod`: Go module dependencies file.
- `go.sum`: Checksums for module dependencies.
- `main`: Entry point of the application.
- `Makefile`: Contains common commands to build or manage the project.
- `README.md`: Project documentation and guidance.

## Requirement Before You Run This Starter
| Package Name |  IsNeeded  |
|:-----|:--------:|
| Golang >= v1.16   | ✅ |
| PostgreSQL  |  ✅  |
| Golang Migrate   | ✅ |
| Makefile   | ✅ |
| Air Hot Reload🔥   | ✅ | 

## Instalation
Clone this repository correctly
```
git clone http://github.com/meedeley/meedeley_code.git
```
🔥🔥🔥
```
cd meedeley_code
go mod tidy
```
💧 and voilla...
```
make run //or if u want run manually
go run cmd/app/main.go
```
## Makefile documentation
This Makefile provides a set of commands to build, run, and manage the application, as well as handle database migrations. Below is a detailed explanation of each section and command.

#### Application Settings
- BINARY_NAME:      # Name of the application binary (default: app).
- BINARY_DIR:       # Directory where the binary file will be generated (default: bin).
- MAIN_FILE:        # Path to the main Go application file (default: cmd/app/main.go).

#### Database Settings
- MIGRATIONS_DIR:   # Directory for database migration files (default: db/migrations).
- DB_USER:          # Database username (default: postgres).
- DB_PASSWORD:      # Database password (default: postgres).
- DB_HOST:          # Database host (default: localhost).
- DB_PORT:          # Database port (default: 5432).
- DB_NAME:          # Database name (default: meedeley).
- DATABASE_URL:     # Full PostgreSQL connection URL.

#### Build Commands

```
make build
```
Builds the application binary and places it in the specified BINARY_DIR.
```
make build
```
Removes all build files from the BINARY_DIR.
make clean

##### Migration Commands
Creates a new migration file in the MIGRATIONS_DIR. Requires a name argument.
```
make migration name=<migration_name>
```
Runs all pending migrations.
```
make migrate-up
```
Runs the next pending migration.
```
make migrate-up-one
```
Rolls back the last migration.
```
make migrate-down
```
Rolls back all migrations.
```
make migrate-down-all
```
Forces the migration version to a specific value. Requires a version argument.
```
make migrate-force version=<version_number>
```
Displays the current migration version in the database.
```
make migrate-version
```
###### Helper Command
Tests the database connection using the configured DATABASE_URL.
````
make db-test
````
