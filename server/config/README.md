# Haddock server config
Haddock server generate a server.properties file in the data directory (which can be pointed virtually anywhere on the host machine).

| Property                  | Default              | Description                                                                                                         |
|---------------------------|----------------------|---------------------------------------------------------------------------------------------------------------------|
| allowAnonymousUsers       | false                | Define if haddock should allow anonymous users to use your haddock server                                           |
| debugMode                 | false                | Recommended for development, show error details (Fatal errors will always be shown independently from this setting) |
| dockerSocketPath          | /var/run/docker.sock | Path to the docker socket                                                                                           |
| host                      | ":8080"              | Host to listen on, should contain a port                                                                            |
| tokenExpiration           | 2h                   | How long the token should be valid for (supports "ns", "us" (or "µs"), "ms", "s", "m", "h"; set 0 to disable).      |
| rememberMeTokenExpiration | 240h (10days)        | How long the token should be valid for if the user clicked 'remember me' (works the same as 'tokenExpiration')      |