# usermanager

Database details:
    host     : localhost,
    port     : 5432,
	user     : postgres,
	password : 1234,
	database   : factly

 you can change database data in config/local.go

 once database part is done. run usermanager.go

port = 8015  can be change in usermanager.go

http://localhost:8015/api/user - POST  create user

payload :

{
   "name": "ross geller",       //string | text
    "age": 31,                  //int    | int
    "phone": "+1 91892334690",  //string | text
    "city": "ny"                //string | text
}


http://localhost:8015/api/user - GET  List user

no payload

http://localhost:8015/api/user/{id} - GET  Read a user

no payload                       

http://localhost:8015/api/user/{id} - DELETE  delete user

http://localhost:8015/api/user/{id} - PUT  Update user

Payload :   

{
   "name": "ross geller",       //string | text
    "age": 31,                  //int    | int
    "phone": "+1 91892334690",  //string | text
    "city": "ny"                //string | text
}

//no need to provide all field just the ones to be changed
sample:

{               
    "phone": "+1 91892334690",  //string | text
    "city": "ny"                //string | text
}



if any queries,contact me through nirmal68b@gmail.com