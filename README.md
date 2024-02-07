📝 Documentation 📝

### Description:
This Go code provides a basic API for managing advertisements. It includes endpoints for creating, reading, updating, and deleting advertisement entries in a simple in-memory database.

### Structures:
- **advertisement**: Represents an advertisement entry with fields such as ID, Title, Photo URL, Description, Price, and CreatedAt.
- **allAdvertisements**: A slice containing all advertisement entries.
- **Advertisements**: A variable holding all advertisement entries.

### Endpoints:
1. **GET /advertisement**:
   - Description: Displays a welcome message.
   - Method: GET

2. **POST /advertisements**:
   - Description: Creates a new advertisement entry.
   - Method: POST
   - Body: JSON data containing Title and Description of the advertisement.

3. **GET /advertisements**:
   - Description: Retrieves all advertisement entries.
   - Method: GET

4. **GET /advertisements/{id}**:
   - Description: Retrieves a specific advertisement entry by ID.
   - Method: GET

5. **PATCH /advertisements/{id}**:
   - Description: Updates a specific advertisement entry by ID.
   - Method: PATCH
   - Body: JSON data containing updated Title and Description.

6. **DELETE /advertisements/{id}**:
   - Description: Deletes a specific advertisement entry by ID.
   - Method: DELETE

### Running the Application:
1. Ensure you have Go installed on your system.
2. Install Gorilla Mux package: `go get -u github.com/gorilla/mux`
3. Run the main application file.
4. The application will start a server at `localhost:8080`.

### Examples:
1. To create a new advertisement:
   - Endpoint: POST `http://localhost:8080/advertisement`
   - Request Body:
     ```json
     {
         "Title": "New Advertisement",
         "Description": "A new ad description"
     }
     ```

2. To update an advertisement by ID:
   - Endpoint: PATCH `http://localhost:8080/advertisements/{id}`
   - Request Body:
     ```json
     {
         "Title": "Updated Title",
         "Description": "Updated Description"
     }
     ```

3. To delete an advertisement by ID:
   - Endpoint: DELETE `http://localhost:8080/advertisements/{id}`

### Note:
- This code uses an in-memory database, so all data will be lost upon restarting the server.
- For a production environment, consider integrating a persistent data storage solution.

Happy coding! 🚀
