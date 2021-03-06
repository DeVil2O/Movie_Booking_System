

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/DeVil2O/Movie_Booking_System">
    <img src="cinema.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Movie Theatre Booking System</h3>

</p>

## Table of Contents

* [About the Project](#about-the-project)
  * [Built With](#built-with)
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
* [Usage](#usage)
* [Roadmap](#roadmap)
* [Contributing](#contributing)
* [Contact](#contact)

<!-- ABOUT THE PROJECT -->
## About The Project

<p align="center">
  <a href="https://github.com/DeVil2O/Movie_Booking_System">
    <img src="image.png">
  </a>
</p>
" Movie Theatre Booking System " the goal of this project is to develop the backend service for the ticket booking system for the theatre counter. I have   implemented the services in "GO" programming language which can handle upto 1 million requests per second without use of any framework and the database used is "Mongo DB" which is very efficient working database for the management of any type of requests. According to the project, there is admin window which will be provide every service related to the ticket booking. The services including in the project are as follows : 

* Book a ticket using a user’s name, phone number, and timings
* Can also update the ticket timings
* View all the tickets for a particular time
* Delete a particular ticket
* View the user’s details based on the ticket id
* Mark a ticket as expired if there is a diff of 8 hours between the ticket timing and current time


### Built With

* [GO](https://golang.org/)<code><img height="20" src="https://raw.githubusercontent.com/github/explore/80688e429a7d4ef2fca1e82350fe8e3517d3494d/topics/go/go.png"></code>

* [Mongo DB](https://www.mongodb.com/)<code><img height="20" src="https://raw.githubusercontent.com/github/explore/80688e429a7d4ef2fca1e82350fe8e3517d3494d/topics/mongodb/mongodb.png"></code>


## :gear: Prerequisites

You need to have `Go >= 1.13` and `mongo DB` other necessary dependencies before you build it yourself.

### Installation
Make sure you have installed go language and mongoDB

1. Clone The Repository
```shell
git clone https://github.com/DeVil2O/Movie_Booking_System.git
```

2. Install The Go Packages
```shell
github.com/dgrijalva/jwt-go
```
```shell
github.com/globalsign/mgo
```
```shell
github.com/gorilla/mux
```
```shell
github.com/gorilla/securecookie
```
```shell
github.com/sony/sonyflake
```
```shell
go.mongodb.org/mongo-driver
```
```shell
golang.org/x/crypto
```
3. Setup the mongodb connection url as " mongodb://localhost:27017 " for testing purposes.

## Usage

<p align="center">
    <img src="Postman Images/Adminregistration.png" alt="Logo">
  <p align="center">*******Admin Registration using Routing through gorilla/mux*******</p>
    <img src="Postman Images/adminaccount.png" alt="Logo">
  <p align="center">*******Admin Account shown in postman using Routing through gorilla/mux*******</p>
    <img src="Postman Images/adminlogin.png" alt="Logo">
  <p align="center">*******Admin Login using Routing through gorilla/mux*******</p>
    <img src="Postman Images/housefullon20tickets.png" alt="Logo">
  <p align="center">*******Showing Housefull on selling maximum of 20 tickets*******</p>
    <img src="Postman Images/mongodbdatabase.png" alt="Logo">
  <p align="center">*******Database Structure on MongoDB*******</p>
    <img src="Postman Images/ticketcreate.png" alt="Logo">
  <p align="center">*******Ticket Creation using Routing through gorilla/mux*******</p>
    <img src="Postman Images/ticketdeletion.png" alt="Logo">
  <p align="center">*******Ticket Deletion using Routing through gorilla/mux*******</p>
    <img src="Postman Images/ticketsattime.png" alt="Logo">
  <p align="center">*******Tickets shown at a particular time using Routing through gorilla/mux*******</p>
    <img src="Postman Images/timingupdate.png" alt="Logo">
  <p align="center">*******Updating the timing on the tickets using Routing through gorilla/mux*******</p>
    <img src="Postman Images/userdetailswithid.png" alt="Logo">
  <p align="center">*******Getting the user details using the ticket id using Routing through gorilla/mux*******</p>
</p>


## RoadMap
If you want to see a new feature feel free to [create a new Issue](https://github.com/DeVil2O/Movie_Booking_System/issues/new)

## Contributing

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Author

Chirag Garg - [@DeVil2O](https://github.com/DeVil2O)

Project Link: [https://github.com/DeVil2O/Movie_Booking_System](https://github.com/DeVil2O/Movie_Booking_System)


