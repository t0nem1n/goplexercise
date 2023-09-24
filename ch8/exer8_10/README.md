
Chat server that only have one room, when someone join, he/she cannot see the history of chat before, the server will only broadcast new message.

To be able to broadcase new message, we use a map to store clients channel. map[string]chan string.
A broadcaster in main function will broadcast the message to channel.

There will be 3 events:
- entering
- sending massage
- leaving
