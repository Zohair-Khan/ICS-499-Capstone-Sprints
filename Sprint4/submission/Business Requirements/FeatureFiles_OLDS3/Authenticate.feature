Feature: Create Progress Notes
As a user,
I need to authenticate myself
So I can use features of the system I am authorized to use.

# Happy path
Scenario:		User successfully logs in with valid credentials
Given 			The user has not logged in
And             On the login page
When			The user submits a valid username and password
Then			The system will verify the submitted username and password
And				Log the user in
And             Save a cookie containing authentication information to the user's browser session
And				Redirect the user to the dashboard

# Unhappy path
Scenario:		User unsuccessfully logs in with invalid credentials
Given 			The user has not logged in
And             On the login page
When			The user submits an invalid username and password
Then			The system will fail to verify the submitted username and password
And				Display an "invalid credentials" error message

# Happy path
Scenario:       User successfully logs out
Given           The user has logged in
When            The user clicks "Log Out"
Then            The system will log the user out
And             Delete the cookie containing authentication information