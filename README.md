# Serieall API
The project aims to provide an api to interface with NATS and the BDD for the SerieAll website.

## Development
Before starting to contribute to this project, you need to setup a NATS Streaming Server. A docker-compose file is present at the root of the repo for this purpose.
Just launch it : 
```bash
make dev
``` 

Yo can test sending a message to NATS with the bin/nats_pub.py script.

## Tests
The project has unit tests and you need to pass all of them before your PR is validated.
Just run :
```bash 
make unittests
```

## Contributing

Thanks for thinking about contributing to worker_images! The success of an open source project is entirely down to the efforts of its contributors, so thank you for even thinking of contributing.

Before you do so, you should check out our contributing guidelines in the [CONTRIBUTING.md](CONTRIBUTING.md) file, to make sure it's as easy as possible for us to accept your contribution.