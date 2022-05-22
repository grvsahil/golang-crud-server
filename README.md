Manual to use different routes

#Register user

route :- /user
method :- post

{   
    "fname": "firstname",
    "lname": "lastname",
    "email": "user@gmail.com",
    "dob": "2000-01-01",
    "password": "mysecretkey"
}


#Login user

route :- /login
method :- post

{
    "email":"user@gmail.com",
    "password":"mysecretkey"
}


#Logout user

route :- /logout
method :- get


#Update user

id -> user's id to update details

route :- /users/id
method :- patch


#Delete user

id -> user's id to delete

route :- /users/id
method :- delete


#Get all user details

route :- /users
method :- get


#Search user (Partial search also supported)

Search by id
route :- /users?id=7

Search by name
route :- /users?name=username

Search by archived
route :- /users?archived=true

Search by email
route :- /users?email=useremail@gmail.com


#Sort users (specifying order also supported)

Sort by name
route :- /users?sortby=name

Sort by email
route :- /users?sortby=email

Sort by id
route :- /users?sortby=id

Sort by dob
route :- /users?sortby=dob

Specify order for sorting (ascending by default)
route :- /users?order=desc


#Control pagination (Default 1st page & 3 items per page)

Specify page number
route :- /users?page=2

Specify items per page
route :- /users?items=7
















