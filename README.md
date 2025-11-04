# XKCD Client
A companion program to the (XKCD Server)[https://github.com/Cole-Parsons/XKCD-Scraper-HTTP-Server.git] project.
This client interacts with the XKCD Server’s REST API to request, download, and manage comics automatically over HTTP.
---
## Overview
The XKCD Client is a Go Application designed to:
* Periodically generate random comic numbers
* Query the XKCD Server to check if a comic is downloaded
* Request the server to download missing comics
* Retrieve and store the downloaded images locally
This project demonstrates HTTP communication, RESTful client design, and docker integration in Go.
---
