# User Notes App written in Golang

-- Guide --

Demo -- https://www.youtube.com/watch?v=ZsmGPgLCgNU

1. Run Xampp Server.
2. open phpmyadmin Mysql dashboard.
3. Create Database "accuknox" or give any name.

4. Open code in code editor, change database settings according to your inputs and run it.
	database settings: models --> setup.go

5. Open Postman and check the API Endpoints.

6. Endpoints for Postman : 
	1. User Signup : [POST]   localhost:3000/signup
	2. User Login  : [POST]   localhost:3000/login
	3. User Logout : [GET]    localhost:3000/logout
	4. Create Notes: [POST]   localhost:3000/notes
	5. List Notes  : [GET]    localhost:3000/notes
	6. Delete Note : [DELETE] localhost:3000/notes	

That's it!
