Feature: Manage Progress Notes
As a user,
I need to manage progress notes
So I can log and review therapy sessions I am authorized for.

# Happy path
Scenario:       User successfully creates a progress note
Given 			The user has been authenticated
And             Is on the note creation page
And			    Has entered valid values into the form fields
When            The user submits the form
Then			The system verifies the entered values are valid
And				Saves the note to the database
And				Redirects the user to the dashboard

# Unhappy path
Scenario:       User unsuccessfully creates a progress note due to invalid input
Given 			The user has been authenticated
And             Is on the note creation page
And				Has entered invalid values into the form fields
When			The user submits the form
Then			The system determines the values are invalid
And				Displays the appropriate "invalid value" error message

# Happy path
Scenario:       User successfully views the list of progress notes they are authorized to view
Given           The user has been authenticated
And				Is on any page of the website
When            The user clicks the dashboard button
Then            The system determines the notes the user is authorized to view
And				Displays the appropriate notes in chronological order

# Unhappy Path
Scenario:       User views an empty list since no progress notes they are authorized to view exist
Given           The user has been authenticated
And				Is on any page of the website
When            The user clicks the dashboard button
Then            The system determines the user is not authorized to view any notes
And				Displays an empty list

# Happy path
Scenario:       User successfully views a specific progress note they are authorized to view
Given           The user has been authenticated
And				Is on the dashboard page
When            The user clicks on a progress note     
Then            The system verifies they are authorized to view the note
And				Displays the note

# Unhappy path, only occurs when a user enters a note ID they are unauthorized to view in the URL, user is not able to "select" an unauthorized note
Scenario:       User unsuccessfully views a specific progress note they are not authorized to view
Given           The user has been authenticated
When            The user opens a URL to a note they are not authorized to view
Then            The system determines they are not authorized to view that note
And				Displays a "Forbidden" error message

# Unhappy path, only occurs when a user enters a note ID for a note that does not exist in the URL, user is not able to "select" an invalid note
Scenario:		The user tries to open a nonexistent note through URL
Given           The user has been authenticated
When            The user opens a URL to a note that does not exist
Then            The system determines the note does not exist
And				Displays a "Bad Request" error message

# Happy path
Scenario:		Admin changes note status
Given			The user has been authenticated
And				Has admin privileges
And				Is viewing a note
When			The user clicks approve or deny buttons
Then			The system saves the updated note status

# Happy path
Scenario:		User filters the list of notes
Given			The user has been authenticated
And				Is on the dashboard page
When			The user modifies a filter fields
Then			The system displays only the notes matching that filter