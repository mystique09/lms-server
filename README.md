# Class Management API

Back-end for my [Class Management](https://class-management.vercel.app)  website. This is just my second option if prisma still doesn't work.

Tech used:
- Go
- Echo
- Sqlc
- Postgres
- DB diagram

## API Endpoints

```diff
// Auth
POST /api/v1/login
POST /api/v1/signup

// User
GET /api/v1/users
GET /api/v1/users/:id
UPDATE /api/v1/users/:id
DELETE /api/v1/users/:id
GET /api/v1/users/:id/classrooms
GET /api/v1/users/:id/followers
GET /api/v1/users/:id/following
POST /api/v1/users/:id/following
DELETE /api/v1/users/:id/following/:id

// Clasroom
POST /api/v1/classrooms
GET /api/v1/classrooms/:id
UPDATE /api/v1/classrooms/:id
DELETE /api/v1/classrooms/:id
GET /api/v1/classrooms/:id/users
GET /api/v1/classrooms/:id/posts

// Post
POST /api/v1/posts
GET  /api/v1/posts/:id
UPDATE /api/v1/posts/:id
DELETE /api/v1/posts/:id
GET  /api/v1/posts/:id/likes
GET /api/v1/posts/:id/comments
POST /api/v1/posts/:id/likes
DELETE /api/v1/posts/:id/likes

// Comments
POST /api/v1/comments
GET /api/v1/comments/:id
UPDATE /api/v1/comments/:id
DELETE /api/v1/comments/:id
GET  /api/v1/comments/:id/likes
POST /api/v1/comments/:id/likes
DELETE /api/v1/comments/:id/likes
```

## TODO
- [x] Get all user
- [x] Get one user
- [x] Create new user
- [x] Update a user
- [x] Delete a user
- [x] Get all public classroom
- [x] Get all classrooms joined/owned by a user
- [x] Leave a classroom
- [x] Follow fellow user
- [x] Unfollow fellow user
- [x] Get all followers
- [x] Get all users you've followed(followings)
- [x] Get one classroom
- [x] Create new classroom
- [x] Update a classroom
- [x] Delete a classroom
- [x] Implement JWT
- [x] Implement Bcrypt
- [x] Implement Logger
- [x] Implement Rate Limit
- [x] Implement CORS
- [x] Create the initial routers of posts
- [x] Create the initial routers of comments
- [x] Create the initial routers of class_works/assignments
- [x] Create a post inside a classroom
- [ ] Udpate a post
- [ ] Delete a post
- [x] Get all posts
- [x] Get one posts
- [ ] Like a post
- [ ] Unlike a post
- [ ] Get all likes of a post
- [ ] Get all comments of a post inside a classroom
- [ ] Get one comment of a post
- [ ] Like a comment of a post
- [ ] Unlike a comment of a post
- [ ] Delete a comment
- [ ] Update a comment
- [ ] Integrate storage provider for file uploads 
- [ ] Implement file upload for class works
- [ ] Implement charts for teachers
- [ ] Implement role base access

