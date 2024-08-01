Feature: Create Progress Notes
As a user,
I need to create progress Notes
So I can log and maintain the progress of patients I work with

# Happy path
Scenario:       User successfully creates a progress note
Given 			The user is authenticated
When            The user visits the note creation page
And				Enters valid values into the form fields
And             Submits the form
Then			The system saves the note to the database
And				Redirects the user to the dashboard

# Unhappy path
Scenario:       The user unsuccessfully creates a progress note due to invalid input
Given 			The user is authenticated
When            The user visits the note creation page
And				Enters invalid values into the form fields
And             Submits the form
Then			The system raises an error notifying the user which fields contain invalid input

# Unhappy path
Scenario:       The user is not authorized to create a note
Given           The user is not authenticated
When            The user visits the note creation page
Then            The system redirects the user to the login page

