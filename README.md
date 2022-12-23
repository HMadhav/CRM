# CRM

### Requirements 

1. GO
2. MongoDB

### Steps to run the project

1. Clone the repo
2. Make sure you have Go and MongoDB installed 
3. install all the pkg (`go mod tidy`)
4. create a **.env** file in the root of project
5. Past the below code in **.env** file and provide the Username and password.
   > Example `MONGOURI=mongodb+srv://USERNAME:PASSWORD@crm.rimlr8b.mongodb.net/?retryWrites=true&w=majority`.
6. Run the below command in root of your project.
   > `go run main.go`
7. You can import the postman collection(`CRM.postman_collection.json`) attached in this project to test the API endpoints.
8. Not using `io/ioutil` It looks like it's depricated.
