# Hadock
Hadock is an open source project meant to provide everyday apps trought a docker Ecosystem. This allow you to manage which app do you want and easely share app environements.

# Documentation
* [Api](server/api/README.md)
* [Properties](server/properties/README.md)

# Install
You need to clone the projet like this:
`git clone https://github.com/hadock-project/hadock-server`<br/>
and run the app like this: `go run main.go`
Go will automatically download all dependencies and run the project

# Troubleshooting
## Common issues
### Making a request to /api/app return me an error
If your error is à`Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock [...] permission denied`

Haddock failed to access the docker deamon socket, verify that docker is installed or try running go as `sudo` <br/><br/>
It can also happen because the docker socket path is different in your environment, please submit an issue ^^

## My issue isn't covered
Please open a [GitHub issue](https://github.com/haddock-project/haddock-server/issues), we'll try to assist you!
