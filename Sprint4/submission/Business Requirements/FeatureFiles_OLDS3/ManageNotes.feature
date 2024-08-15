Feature: Manage Progress Notes
As a user,
I need to manage progress notes
So I can log and review therapy sessions I am authorized for.

# Happy path
Scenario:       User successfully creates a progress note
Given 			The user has been authenticated
And             on the note creation page
And			    Enters valid values into the form fields
When            The user submits the form
Then			The system will save the note to the database
And				Redirects the user to the dashboard

# Unhappy path
Scenario:       User unsuccessfully creates a progress note due to invalid input
Given 			The user has been authenticated
And             On the note creation page
When			The user enters invalid values into the form fields
And             Submits the form
Then			The system raises an error notifying the user which fields contain invalid input

# Happy path
Scenario:       User successfully views the list of progress notes they are authorized to view
Given           The user has been authenticated
When            The user selects the option to view the dashboard page
Then            The system will display a table of all the progress notes the user is authorized to view with a link to each note, sorted in chronological order, from newest to oldest.

# Unhappy Path
Scenario:       User views an empty table since no progress notes they are authorized to view exist
Given           The user has been authenticated
When            The user selects the option to view the dashboard page
Then            The system will display an empty table.

# Happy path
Scenario:       User successfully views a specific progress note they are authorized to view
Given           The user has been authenticated
And             Authorized to view a specific progress note
When            The user selects the option to view said progress note       
Then            The system will display the details for the progress note

# Unhappy path, only occurs when a user enters a note ID they are unauthorized to view in the URL, user is not able to "select" an unauthorized note
Scenario:       User unsuccessfully views a specific progress note they are not authorized to view
Given           The user has been authenticated
And             Not authorized to view a specific progress note
When            The user selects the option to view said progress note
Then            The system will display a "Forbidden" error message

# Unhappy path, only occurs when a user enters a note ID for a note that does not exist in the URL, user is not able to "select" an invalid note
Given           The user has been authenticated
When            The user visits the page for a progress note that does not exist
Then            The system will display a "Bad Request" error message