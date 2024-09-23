# Two_Factor_Authentication
Login and Register system integrated with two factor authentication (Password and Email), where user is required to input the 4 digit OTP code sent by system to their email which will be expired in 4 minutes. This project using Golang net/http standard library of version 1.22, PostgreSQL, REDIS, and Resend for email notification.
# Feature Showcase
1. Login Form
   
![img](https://drive.google.com/uc?export=view&id=1-o3SNnIB_0UrIEJTNMNnAq_wrQ970fcM)
   
2. OTP Code Sent by Email
    
![img](https://drive.google.com/uc?export=view&id=1ATMJqv0aGiVRi2UM-3ia6XqzH8ku6v0n)
   
3. Valid OTP indicate the login process is Successful
   
![img](https://drive.google.com/uc?export=view&id=11KeX7g5MHOQ1taP7rueTQWGH6-F47Q-g)   
   
# Prerequisite
1. Have PostgreSQL & Redis installed and run on your machine
2. Create new database 
3. Edit ``.env`` by applying your configuration
4. For ``RESEND_API_KEY`` can be obtained by sing in into https://resend.com/
# How To Use
1. Clone this repo by run following script
```
git clone https://github.com/W-ptra/Golang_CRUD_API.git
```
2. Change Directory to main folder & Download all dependency
```
cd Golang_CRUD_API
go mod download
```
3. Run the main.go & api.go
```
go run main.go api.go
```
