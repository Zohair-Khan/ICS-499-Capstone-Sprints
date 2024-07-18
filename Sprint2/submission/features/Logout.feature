Feature: View Specific Progress Note
As a user
I need to be able to log out of my current session
So that I ensure others that use the same machine as me do not have unauthorized access to privileged information

# Happy path
Scenario:       User successfully logs out
Given           The user is authenticated
And             On the dashboard page
When            The user clicks "Log Out"
Then            The system logs the user out
And             deletes the cookie containing authentication information



