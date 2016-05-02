## Overview

Simple http shim for pull images without need to ssh into the box

```bash

docker run -e auth=somesharedsecret -e docker_user=someuser -e docker_pass=somepass -e docker_email=some@test.com -v /var/run/docker.sock:/var/run/docker.sock  -p 3000:3000 -it  maleck13/ose_shim:1.0

```

This exposes a single endpoint

```
POST /docker/images

```

It accepts an array of images in the post body example

```json

["maleck13/authentication:latest"]

```

There is a header auth required, it matches the env var set when executing docker run

```

x-auth:somesharedsecret

```

