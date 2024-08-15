Feature: Create Progress Notes
As a user,
I need to authenticate myself
So I can use features of the system I am authorized to use.

# Happy path
Scenario:		User successfully logs in with valid credentials
Given 			The user has opened the login page
And				Has provided a valid username and password
When			The user clicks the login button
Then			The system verifies the submitted credentials
And				Logs the user in
And             Saves the authentication cookie
And				Redirects the user to the dashboard

# Unhappy path
Scenario:		User unsuccessfully logs in with invalid credentials
Given 			The user has opened the login page
And             Has provided an invalid password or username
When			The user clicks the login button
Then			The system determines the credentials are invalid
And				Displays an "invalid credentials" error message

# Happy path
Scenario:       User successfully logs out
Given           The user is logged in
When            The user clicks "Log Out"
Then            The system logs the user out
And             Deletes the authentication cookie