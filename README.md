# XKCD Client  

The XKCD Client is a companion program to the [XKCD-Server](https://github.com/Cole-Parsons/XKCD-Scraper-HTTP-Server.git),   written in Golang, that automates comic requests through HTTP. It periodically checks whether a random XKCD comic has been   downloaded by the server. If not, it instructs the server to download it and waits until the comic becomes available.  

This client demonstrates HTTP communication, JSON handling, environment configuration, and timed task execution in Go.  
---
## Overview  
The XKCD Client is a Go Application designed to:  
* Periodically generate random comic numbers  
* Query the XKCD Server to check if a comic is downloaded  
* Request the server to download missing comics  
* Retrieve and store the downloaded images locally  
### This project demonstrates HTTP communication, RESTful client design, and docker integration in Go.  
---
## Prerequisites  
* Go 1.22+  
<mark>* A running instance of XKCD Server (from the companion project)</mark>  
* Optional: Docker and Docker Compose for containerizad deployment  
---
# Set Up  
```bash
git clone https://github.com/Cole-Parsons/XKCD-Client.git
cd XKCD-Client
go run XKCD-Client.go
```
# Related Projects
[XKCD-Server](https://github.com/Cole-Parsons/XKCD-Scraper-HTTP-Server.git)  
[XKCD-frontend](https://github.com/Cole-Parsons/XKCD-frontend.git)
---
## To run the entirety of the Scraper see:  
[Docker-compose](https://github.com/Cole-Parsons/XKCD-docker-compose.git)
