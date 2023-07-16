To begin testing you need to run docker-compose up in the root directory to build the project

Once all the images have connected you can login with the initial test account 
Username : "george" 
password: "password" 

send a POST request to localhost:8080/login 

with the username and password in the body to login and get an auth token

after that you can test the application by using the following requests:

[Title](https://www.postman.com/irvyntree/workspace/takehome)