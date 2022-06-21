# Haddock server config
Haddock server generate a server.properties file in the data directory (which can be pointed virtually anywhere on the host machine).

| Property            | Default              | Description                                                                                                         |
|---------------------|----------------------|---------------------------------------------------------------------------------------------------------------------|
| allowAnonymousUsers | false                | Define if haddock should allow anonymous users to use your haddock server                                           |
| debugMode           | false                | Recommended for development, show error details (Fatal errors will always be shown independently from this setting) |
| dockerSocketPath    | /var/run/docker.sock | Path to the docker socket                                                                                           |