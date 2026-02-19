# Aggregator in go
## Description
This is a CLI tool that lets local users aggregate news feeds they have subscribed to. MacOS compatible only.

## Prerequisites

Before running this project, make sure you have the following installed:

- **Go** (1.25+ recommended)  
  https://go.dev/dl/

- **Postgresql**  
  https://docs.sqlc.dev/en/latest/overview/install.html](https://www.postgresql.org/download/macosx/)

## Installation

### 1. Install Go
https://go.dev/doc/install
### 2. Install Postgresql 
 install: `brew install postgresql@16`
 run services in background: `brew services start postgresql@16`
### 3. Install the program
  `go install github.com/tzutzush/aggreGator-go`
### 4. Create database
  Enter the psql shell: `psql postgres`
  Create a database: `CREATE DATABASE <db_name>`;
### 5. Config file
  create a config file called `.gatorconfig.json` and fill as follows:
  
  `{
  "db_url":"postgres://<username>:@localhost:5432/<db_name>?sslmode=disable",
  "current_user_name":"<username>"
  }`
### 6. Available commands
First you have to run commands `with go run . `
`login <name>`: Logs in if user is found in the db, changes the current user name to given name in config file (only one word accepted).
`register <name>`: Puts user into the db.
`reset`: Clears the db.
`users`: Lists users found in db.
`agg <time in format of 1m, 10s, etc>`: Starts to scrape feeds found in the db in intervals corresponding to the given argument, scraping creats posts in db.
`addfeed <url>` Adds the feed to the db.
`feeds`: List all feeds in the db.
`follow <title> <url>`: Current user starts to follow feed with title and url.
`following`: Prints the users followed feeds.
`unfollow <url>`: Current user stops to follow feed with url.
`browse`: List the two latest posts from feeds the user follows.
  

  
  
