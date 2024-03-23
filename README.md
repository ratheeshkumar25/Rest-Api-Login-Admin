
>>>>>>Secure REST API with JWT Authentication and RBAC

This project implements a secure REST API with features for user authentication, role-based access control (RBAC), and student management. JWT (JSON Web Token) authentication and stateless sessions ensure robust security and scalability.
Endpoint	Description	Roles Allowed


>>>>>>students/signup	Allows students to register for accounts.	Public
>>>>>>students/login	Enables students to log in with credentials and receive a JWT upon successful login.	Public
>>>>>>students/logout	Ends an authenticated student's session by invalidating their JWT.	Authorized
>>>>>>students/*	Additional endpoints for student-specific operations (implement as needed).	Authorized
>>>>>>admin/login	Grants admins access to the API through login with credentials and generates a JWT.	Public
>>>>>>admin/logout	Terminates an admin's authenticated session by invalidating their JWT.	Authorized
>>>>>>admin/*	CRUD (Create, Read, Update, Delete) operations on resources (implement as needed).	Admin

Authentication Workflow:

User Login: Users (students or admins) send login credentials to the /students/login or /admin/login endpoint.
Server Verification: The server validates the credentials.
JWT Generation (Success): Upon successful validation, the server generates a JWT containing user information and a set expiration time.
JWT Response: The server sends the JWT back to the user.
Authorization in Subsequent Requests: For protected endpoints, the user includes the JWT in the Authorization header of subsequent requests.
JWT Validation: The server verifies the JWT's signature and expiration before processing the request.
RBAC Enforcement: The server checks the user's role (encoded in the JWT) against allowed roles for the requested operation.
Authorized Access: If the JWT is valid and the user has the necessary role, the server processes the request.
Unauthorized Access (Error): If the JWT is invalid, expired, or the user lacks the required role, the server returns an error response.
Client-Side JWT Storage:

For optimal security, store JWTs securely on the client-side (e.g., HttpOnly cookies with Secure attribute, Local Storage with appropriate security settings).
Running the API:

Prerequisites: Ensure you have necessary dependencies installed (e.g., Go environment, database drivers).
Configuration: Edit configuration files (if any) to set database connection details or other API settings.
Run the Server: Execute the appropriate command to start the API server (e.g., go run main.go).


Replace placeholders with actual implementation details (e.g., dependencies, environment variables).
Enhance security measures as needed (consider refresh tokens for longer-lived sessions, HTTPS for communication).
