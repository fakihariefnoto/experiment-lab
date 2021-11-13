Just experiment with pgx go driver for postgres but i'm using it for yugabytedb

## Run yugabyte

`
docker run -d --name yugabyte \
         -p7000:7000 -p9000:9000 -p5433:5433 -p9042:9042 \
         -v ~/yb_data:$HOME/AppData/yugabytedata \
         yugabytedb/yugabyte:latest bin/yugabyted start \
         --base_dir=$HOME/AppData/yugabytedata --daemon=false
`

## UI

The yb-master Admin UI is available at http://localhost:7000 and the yb-tserver Admin UI is available at http://localhost:9000. To avoid port conflicts, you should make sure other processes on your machine do not have these ports mapped to localhost

## 
