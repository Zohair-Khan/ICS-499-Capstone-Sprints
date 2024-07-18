Feature: View List of Progress Notes
As a user
I need to be able to view a list of all the notes I am authorized to view
So that I can maintain a record of the past therapy sessions I am authorized to view

# Happy path
Scenario:       User successfully views the list of progress notes they are authorized to view
Given           The user is authenticated
When            The user visits the dashboard page
Then            The system displays a table of all the progress notes the user is authorized to view with a link to each note

# Unhappy path
Scenario:       User unsuccessfully views list of progress notes due to lacking authentication
Given           The user is not authenticated
When            The user visits the dashboard page
Then            The system redirects the user to the login page


