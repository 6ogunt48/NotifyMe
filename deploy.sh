#!/bin/bash

doctl registry login

docker buildx build -t registry.digitalocean.com/apprentice/website:latest .

docker push registry.digitalocean.com/apprentice/website:latest

ssh -i ~/.ssh/id_rsa root@138.68.181.207 << EOF

EOF