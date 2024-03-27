# Endpoints
- [] GET /events // Get a list of available events
- [] GET /events/{id} // Show an especific event
- [] POST /event // Creates a new event *
- [] PUT /event/{id} // Updates an event *
- [] DELETE /event/{id} // Deletes an event *
- [] POST /signup // Creates a new user
- [] POST /login // Authenticate user
- [] POST /events/{id}/register // Register user for event *
- [] DELETE /events/{id}/register // Cancel registration *

// * Requires authentication

# Next goals
- [] Improve error messages for 404 (it is currently 500) FIND BY ID, UPDATE, DELETE
- [] Improve error messages (do not display internal error messages raw)
