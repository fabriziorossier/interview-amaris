# Amaris
Technical Interview Project for Amaris Consulting

# Requirements
## Functional
- Make a local Go service available
- Use a database to store pokemon data
- Create an API/pokemonname endpoint to access pokemon data from the https://pokeapi.co/ API
- Check if the pokemon data is already stored in the database, in that case respond accordingly, if it's not in the database, store the pokemon data with 3 attributes plus the ID in the database 
- Obtain the name of the pokemon from the pokemon.name field, suppose that sometimes the name arrives as a string in the pokemon.pokemon field. The service should handle both cases

## Non functional
- The project should be up and running just by executing a "docker compose up -d" command
- Create corresponding Unit Tests
- Provide a UML Component Diagram