## Overview

docker run -e docker_user=someuser -e docker_pass=somepass -v /var/run/docker.sock:/var/run/docker.sock -v /usr/bin/docker:/usr/bin/docker -v /usr/local/bin/docker:/usr/local/bin/docker  -p 3000:3000 -it  maleck13/ose_shim

