# Historinhas

Historinhas is a web application designed to create and manage short stories. It provides a platform where users can write, save, update, and delete their stories.

## Features

- **Create**: Users can create new stories by providing a title and content.
- **Update**: They can update existing stories by modifying their title or content.
- **Delete**: Users can delete stories they no longer want to keep.
- **Read**: They can view all stories or read individual stories.

## Technologies Used

- **Go (Golang)**: Backend language used for server-side development.
- **Gin**: Web framework used to handle HTTP requests and responses.
- **GORM**: Object-relational mapping (ORM) library for database interactions.
- **SQLite**: Lightweight database used for storing user data.
- **dotenv**: Library for managing environment variables.
- **Swagger**: API documentation tool for documenting endpoints.


## API Documentation

- The API documentation is available at `http://localhost:42069/docs`

## Usage

1. **Create a Story**:
   - Endpoint: `POST /histories`
   - Request body:
     ```json
     {
         "name": "Story Title",
         "history": "Story Content",
         "userID": 1
     }
     ```

2. **Update a Story**:
   - Endpoint: `PUT /histories/:historyId`
   - Request body:
     ```json
     {
         "name": "Updated Story Title",
         "history": "Updated Story Content",
         "userID": 1
     }
     ```

3. **Delete a Story**:
   - Endpoint: `DELETE /histories/:historyId`

4. **Get All Stories**:
   - Endpoint: `GET /histories`

5. **Get a Story by ID**:
   - Endpoint: `GET /histories/:historyId`

# assets 
![swag](assets/swag.png)

![tests](assets/test.png)

# log

archlinux:historinha toga$ git log    
commit 58f9106ed142d930c44044853c3b6788a9d4cb4d (HEAD -> main, origin/main)
Author: Rafael Coutinho <70056727+c0utin@users.noreply.github.com>
Date:   Sat Jun 1 13:43:25 2024 -0300

    Update README.md

commit 159618c74ace9cfae61528b73d7d768ac6200c04
Author: Rafael Coutinho <70056727+c0utin@users.noreply.github.com>
Date:   Sat Jun 1 13:41:10 2024 -0300

    Create README.md

commit f019a020538a3ece254c2cbc0eeac1e50511cdca
Author: Rafael Coutinho <70056727+c0utin@users.noreply.github.com>
Date:   Sat Jun 1 13:37:28 2024 -0300

    Update user_controller_test.go

commit 1f78300e1814c81051eb53741e0b9774c4798674
Author: Rafael Coutinho <70056727+c0utin@users.noreply.github.com>
Date:   Sat Jun 1 13:37:16 2024 -0300

    Update history_controller_test.go

commit 75fb529572ea448c1abf84b806fac9ac6ace3f90
Author: c0utin <rafaelcouto111@gmail.com>
Date:   Sat Jun 1 13:36:15 2024 -0300

    feat: add historinhas and gemini integration

commit 7c7aff84778c269fe53d904b87c5997b8c40e08d
Author: c0utin <rafaelcouto111@gmail.com>
Date:   Sat Jun 1 12:07:11 2024 -0300

    feat: initial structure
    
    Add CRUD for user and visual documentation




## License

This project is licensed under the [MIT License](LICENSE).
