## Overview
```bash

docker run -e docker_user=someuser -e docker_pass=somepass -e docker_email=some@test.com -v /var/run/docker.sock:/var/run/docker.sock -v /usr/bin/docker:/usr/bin/docker  -p 3000:3000 -it  maleck13/ose_shim:1.0

```

