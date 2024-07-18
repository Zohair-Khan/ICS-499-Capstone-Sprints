Feature: View Specific Progress Note
As a user
I need to be able to view any specific progress note I am authorized to view
So that I can review past therapy sessions I am authorized to view

# Happy path
Scenario:       User successfully views a specific progress note they are authorized to view
Given           The user is authenticated
And             Is authorized to view a specific proogress note
When            The user visits the page to view said progress note
Then            The system displays the details for the progress note

# Unhappy path
Scenario:       User unsuccessfully views a specific progress note due to lacking authentication
Given           The user is not authenticated
When            The user visits the page to view a specific progress note
Then            The system redirects the user to the login page

# Unhappy path
Scenario:       User unsuccessfully views a specific progress note they are not authorized to view
Given           The user is authenticated
And             Is not authorized to view a specific proogress note
When            The user visits the page to view said progress note
Then            The system redirects the user to the dashboard

# Unhappy path
Scenario:       User unsuccessfully tries to view a note that does not exist
Given           The user is authenticated
When            The user visits the page to view a progress note that does not exist
Then            The system redirects the user to the dashboard


