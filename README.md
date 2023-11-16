# ports-exercise

This application implements a simple solution to read ports data from a JSON and store it in an in-memory database.

## Architecture
I opted to use hexagonal architecture, as recommended. All business logic is isolated under `ports` directory, that 
represents the domain. The ports (the ones from the architecture's concept) are represented by repository and service 
interfaces. The adapters are represented by the concrete implementation for the in-memory ports repository and the 
json handler.

## How to run
You can use Docker to build and run the application with
``` shell
docker run $(docker build -q .)
```
Or test it with 
```
docker run $(docker build --target test -q .)
```

I did not follow the recommendation to not compile the code using Docker (sorry). 

## Decisions and Future Enhancements

### Database
I used the most simple form of in-memory store possible: a map. So the storage is just map in the repository.
This obviously does not cover some basic and desirable features from databases, like persistence or transactions.
If I continued this implementation, my choice would be some key-value database, like [Redis](https://redis.com/nosql/key-value-databases/).

### JSON decoding

By far the part I spent most time with. To meet the memory constraints as the filesize is unknown and can be really large, 
I decided to implement a decoder that goes through line by line. 
It is a slow approach, specially if we introduce a write operation to a real database. Ideally, with more time, I would try to build 
a decoder that can read data chunks from the JSON to optimize the memory consumption vs. speed tradeoff. 


