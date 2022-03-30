# Programmer Heaven (PH)
Monorepo software to make a monorepo where apps are written for multiple programming languages. It is written in Go (Golang).

# How to run/build
1. Download the code
2. Have Go installed
3. Run go build
4. Use one of the commands

# Commands
Create a new project

`./ph init`

Create a new project without filling in the information.

`./ph init -y`

Add an app to the project (Requires to use `./ph init` first)

`./ph add-app --name APP_NAME`

Add an app to the project with no name (Requires to use `./ph init` first)

`./ph add-app`