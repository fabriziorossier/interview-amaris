# Amaris
Technical Interview Project for Amaris Consulting

## Prerequisites
- **Docker Engine - Version 27.3.1 and up:** https://docs.docker.com/engine/install/
- **Docker Compose - Version 2.29.7 and up:** https://docs.docker.com/compose/install/

## Running the project
- **From the project root folder:** docker compose up -d
- **Check if the containers are running:** docker ps -a
- *If any container is not running, rerun* **docker compose up -d**

## Accessing the project
- http://localhost:8080/pokemon/pokemonname
- Replace pokemonname with the name of the pokemon you want to access

## Accessing the database (to check the saved Pokemon data)
- **From the project root folder:** docker exec -it amaris-postgres-db bash
- **In the amaris-postgres-db container bash shell:** psql -h localhost -d developmentdatabase -U developmentuser
- **In the developmentdatabase database:** SELECT * FROM pokemon;

# Use of this software
This software is provided solely for the technical evaluation process of Fabrizio Alexander Rossier González for Amaris Consulting and is not intended for commercial or non-commercial use without appropriate permission by the former.

# License
Copyright Fabrizio Alexander Rossier González. All rights reserved.

Licensed under the Creative Commons Attribution-NonCommercial-NoDerivs 3.0 Chile (CC BY-NC-ND 3.0 CL) license available in https://creativecommons.org/licenses/by-nc-nd/3.0/cl/legalcode and in the adjunt LICENSE file.