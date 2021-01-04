# Hexgo

The backend that powers [HextechDocs](https://hextechdocs.dev).

---

## Features

The software is written in Go and acts both as the importer (that imports the contents of its working directory into the system) and the forward facing GraphQL server.

The GraphQL schema is generated programmatically and it's well commented so any GraphQL client should be able to provide documentation via introspection.

---

## Deployment

### Using the HashiCorp stack

If your system happens to be running Consul and Nomad, you can deploy the system using automatic configuration and dynamic ports.
All you have to do is set the environment variable `CONSUL_KV_PATH` to point to the configuration file you'd like to use from Consul KV.
As for nomad, you have to name the port `http` for the environment variable `NOMAD_HOST_PORT_http` to be correctly populated.
*(we use the standard API libraries provided by HashiCorp so don't forget to set the environment variables for ACL to make sure the software can authenticate against Consul)*

### Without the HashiCorp stack

Configuration via file can still be done without ever touching any other service. All you have to do is set the value of `ANU_DISABLE` to anything but an empty string and on the first run, the software will generate an example configuration file in its working directory.

### Process arguments

| #Flag    | #Example  | #Description                                                 |
| -------- | --------- | ------------------------------------------------------------ |
| importer | -importer | Runs the software in importer mode as opposed to webserver mode. If not set, the software will run in webserver mode. |
| port     | -port=80  | Changes the default server port of 8080 to something else. Can be overwritten when using the HashiCorp stack. |

