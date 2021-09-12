# The challenge

The objetive of this challenge is to create an API that transfers money from one
account to the other in a digital bank.

# Build & Usage

To use the API defined in this service you can either use docker or build it
yourself using a local golang toolchain. When using docker you can build the
image using the following command:

`docker build -t stone .`

To run/stop the service you can use the following commands:

`docker run -d --network=host stone`
`docker stop <container-pid>`

For simplicity there's a Makefile in the root of the project that can be used to
start the service.