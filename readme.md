#rest-api for user service

GET /users - list of users -                200(OK), 404(Not Found), 500(Server fail) 
GET /users/{id} - user by id -              200(OK), 404(Not Found), 500(Server fail)
POST /users/{id} - create new user -        204(No content), 4xx, Header Location: url
PUT /users/{id} - fully update user -       204/200, 404, 400, 500 Body: fully updated user
PATCH /users/{id} - partially update user - 204/200, 404, 400, 500 Body: partially updated user
DELETE /users/{id} - delete user -          200(OK), 404(Not found), 400