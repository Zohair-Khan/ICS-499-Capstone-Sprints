Feature: Login
As a user
I need the system to authenticate who I am
So that I can perform actions I am authorized to do.

# Happy path
Scenario:		User successfully logs in with valid credentials
Given 			The user is not logged in
When            The user visits the login page
And				submits a valid username and password
Then			The system verifies the submitted username and password
And				Logs the user in
And				Redirects the user to the dashboard

# Unhappy path
Scenario		User unsuccessfully logs in with invalid credentials
Given 			The user is not logged in
When            The user visits the login page
And				submits an invalid username and password
Then			The system fails to verify the submitted username and password
And				Displays an "invalid credentials" error message

# Un/Happy path? not sure
Scenario		User tries to visit the login page while already logged in
Given 			The user is logged in
When            The user visits the login page
Then			The system redirects the user to the dashboard




