To begin testing, you need to run `docker-compose up` in the root directory to build the project.

Once all the images have connected, you can log in with the initial test account:
Username: "george"
Password: "password"

Send a POST request to `localhost:8080/login` with the username and password in the body to log in. Once done, you can copy the value in the authorization header to send subsequent requests.

After that, you can test the application by using the following requests:

[Title](https://www.postman.com/irvyntree/workspace/takehome)
