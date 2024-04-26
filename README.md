# scalex.club
Create library management system for admin and regular user uncluding jwt authentication


## Backend Developer Assignment: Library Management System
## Backend Developer Assignment: Library Management System

### Objective
The goal of this assignment is to build a backend server for managing a library of books. The system will differentiate between two types of users: Admin and Regular Users. You will be implementing authentication, authorization, and data manipulation operations based on user roles.

### Getting Started

#### Requirements
- Go programming language
- Gin framework for building web applications
- JWT for authentication
- Any text editor like VSCode
- Postman for testing the API endpoints

#### Setup Instructions
1. Clone the repository from GitHub.
2. Navigate into the project directory.
3. Install dependencies using `go mod tidy`.
4. Start the server using `go run main.go`.
5. The server will be running on `http://localhost:8080`.

### API Endpoints

#### 1. `/login` - User Login
- **Method:** POST
- **Description:** Authenticates users and provides a JWT for authorization.
- **Request Body:** 
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
- **Responses:**
  - **200 OK:** Returns a JWT token.
  - **401 Unauthorized:** If credentials are invalid.

#### 2. `/home` - List Books
- **Method:** GET
- **Description:** Displays books available to the user based on their role.
- **Authorization:** Bearer Token (JWT)
- **Responses:**
  - **200 OK:** Returns a list of books.
  - **401 Unauthorized:** If token is missing or invalid.

#### 3. `/addBook` - Add a Book
- **Method:** POST
- **Description:** Allows admin users to add a new book to the library.
- **Authorization:** Bearer Token (JWT)
- **Request Body:** 
  ```json
  {
    "bookName": "string",
    "author": "string",
    "publicationYear": "number"
  }
  ```
- **Responses:**
  - **201 Created:** If the book is successfully added.
  - **400 Bad Request:** If input validation fails.
  - **403 Forbidden:** If user is not an admin.

#### 4. `/deleteBook` - Delete a Book
- **Method:** DELETE
- **Description:** Allows admin users to delete a book from the library.
- **Authorization:** Bearer Token (JWT)
- **Request Params:**
  - **bookName:** `string` - The name of the book to delete.
- **Responses:**
  - **200 OK:** If the book is successfully deleted.
  - **400 Bad Request:** If validation fails.
  - **403 Forbidden:** If user is not an admin.

### How to Test the API
1. Use Postman or any API testing tool to send requests to the server.
2. Start by logging in to receive a JWT.
3. Use the JWT to authenticate other requests based on the required permissions.
4. Test `/home`, `/addBook`, and `/deleteBook` endpoints by following the above specifications.

### Submission
1. Ensure your code is well-commented and follows coding standards.
2. Zip your project directory.
3. Upload the zip file to Google Drive and ensure the file has view access enabled.
4. Submit the Google Drive link as per the submission instructions provided.

### Important Notes
- Ensure all data manipulations are done safely and consider edge cases for data handling.
- Handle all possible errors gracefully and return appropriate error messages to the user.

Good luck with your assignment!