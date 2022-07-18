
What Is This? 

Hangman-go is a work in progress asynchronous and multiplayer application
that implements the well known hangman game [Wikipedia Entry](https://en.wikipedia.org/wiki/Hangman_(game)) 

This application will ultimately be  used to drive a front end implementation

The application can be accessed on `{$host}:8080/` (Will soon allow the port to be configured)

# So What does it currently do ? #

### Create a game ### 
`GET /create`
* Proqvides a REST API to create a game, this will return a game payload thats below

```
{
"id": "e853cdd8-061e-11ed-9342-acde48001122",
"word": "cat", // This will soon be retured observate
"usedCharacters": [],
"attemptsLeft": 10,
"numberOfPlayers": 0,
"status": 0
}
```

* Currently, a game can have 3 states IN-PROGRESS(0), FINISHED-WIN(1), FINISHED-LOSE(2)


### Make a guess for a game ###
`PUT /guess/{gameId}`
* Provides a REST API to make a guess on what letters the word contains by trying a single letter request payload example below
```
{
 "id":  "6e0a6f26-05c6-11ed-b36b-0aa4f8902d12",
 "letter" :"c"
}
```

### Connect to a game ###
`GET /connect/{id}` <- If everything is valid the connection will be upgraded to a websocket 
* Provides a websocket to connect to a game with a given id.
* The server will continuously broadcast to the game state to all connected clients 





