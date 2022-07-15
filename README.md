# Class Management API

Back-end for my [Class Management](https://class-management.vercel.app)  website. This is just my second option if prisma still doesn't work.

Tech used:
- Go
- Echo
- Sqlc
- Postgres
- DB diagram

## API Endpoints
| REQUEST METHOD    | Response Type    | Endpoint    |
|---------------- | --------------- | --------------- |
| GET    | []User    | /api/v1/users    |
| POST    | String    | /api/v1/users    |
| GET   | User   | /api/v1/users/:id   |
| UPDATE    | User    | /api/v1/users/:id    |
| DELETE    | ID   | /api/v1/users/:id    |
| GET    | []Class    | /api/v1/classrooms    |
| POST    | String    | /api/v1/classrooms    |
| GET   | Class   | /api/v1/classrooms/:id   |
| UPDATE    | Class    | /api/v1/classrooms/:id    |
| DELETE    | ID   | /api/v1/classrooms/:id    |

## Response Types
- User = { username, email, role, created_at, updated_at }
- ID = UUID
- String

## TODO
- [x] Get all user
- [x] Get one user
- [x] Create new user
- [x] Update a user
- [x] Delete a user
- [x] Get all classroom
- [ ] Get one classroom
- [ ] Create new classroom
- [ ] Update a classroom
- [ ] Delete a classroom
- [x] Implement JWT
- [x] Implement Logger
- [x] Implement Rate Limit
- [x] Implement CORS
- [ ] Integrate storage provider for file uploads 
- [ ] Implement file upload for class works
- [ ] Implement charts for teachers
- [ ] Implement role base access

