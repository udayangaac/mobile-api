
##### 1. Signup 
- Requset
```json
{
  "name": "Jayan",
  "email": "administrator@gmail.com",
  "password": "password",
  "mobile": 778990427,
  "Address": "Address",
  "DOB": "2010-11-12",
  "job_status": 1,
  "job_details": {
    "name": "Sampath Bank",
    "address": "Kirulapana"
  },
  "married": 1,
  "family": {
    "kids": 1
  },
  "location": {
    "lat": "",
    "lon": ""
  }
}
```
- Response
```json
{
  "message": "Successfully created the user account."
}
```

##### 2. Login Request
- Request
```json
{
  "email": "administrator@gmail.com",
  "password": "password",
  "location": {
    "lat": "",
    "lon": ""
  }
}
```
- Response
```json
{
  "id": 2,
  "name": "Client",
  "client_id": 1,
  "email": "client@clickapps.co",
  "avatar": "http://localhost:3000/default_image.png",
  "lbs_notification": {
    "id": 8704,
    "content": "Notification Body"
  },
  "token": "eyJhbGciOiJIUzI1NiJ9.eyJpZCI6MiwibmFtZSI6IkNsaWVudCIsImVtYWlsIjoiY2xpZW50QGNsaWNrYXBwcy5jbyIsIm1vYmlsZSI6IjEyMzY1NDc4OSIsImltYWdlIjoiL2RlZmF1bHRfaW1hZ2UucG5nIiwiYWRtaW4iOmZhbHNlLCJpYXQiOjE1NDc5MjU0MzIsImV4cCI6MTU1MDUxNzQzMn0.4Vyjd7BG7v8AFSmGKmIs4VM2FBw3gOLn97Qdf6U4jxU"
}
```