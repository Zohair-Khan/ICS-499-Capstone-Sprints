# Purple Hippos Capstone Project
## Running the code
- download the zip file
- use MySQL Workbench or the terminal to run the db-creation.sql file in MySQL
- in the terminal, navigate to the project folder
- to run, type `go run ./cmd/web` and press enter
- a message should appear with the content "starting server"
- navigate to localhost:8080 in your browser to see it in action

## Alternate
- go to [our website](https://ph-notes.com/)

## TODO

- Feedback for filling in forms
- Authorization levels
- User friendly error response
- Admin functionality

## Things to do
- Update links on pages
- Plan
    - use integer of auth level
    - display changes based on auth level
    - partial html uses auth level to display different things
    - for others, it should probably be separate pages to make things not so complicated 
        (have one note view page for admin, one for providers)
    - have middleware that universally checks token for level, passes it as request
