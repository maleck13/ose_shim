## Overview

Simple http shim for pull images without need to ssh into the box

```bash

docker run -e docker_user=someuser -e docker_pass=somepass -e docker_email=some@test.com -v /var/run/docker.sock:/var/run/docker.sock  -p 3000:3000 -it  maleck13/ose_shim:1.0

```

