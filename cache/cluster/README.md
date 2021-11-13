My Experiment with redis-cluster library


redis-cli --cluster create 127.0.0.1:7000 127.0.0.1:8000 127.0.0.1:9000

`
Master[0] -> Slots 0 - 5460
Master[1] -> Slots 5461 - 10922
Master[2] -> Slots 10923 - 16383
M: 51d1de9c6f4da983ffb53089651e7fd113226b4d 127.0.0.1:7000
   slots:[0-5460] (5461 slots) master
M: 7e91b31997f19c9929b4a742a4bc3526e77fc9bb 127.0.0.1:8000
   slots:[5461-10922] (5462 slots) master
M: 7d43b492587383efe2332cc6a3b8e995a7597026 127.0.0.1:9000
   slots:[10923-16383] (5461 slots) master
`

----
Another way
$ redis-cli --cluster create 127.0.0.1:7000 127.0.0.1:8000 127.0.0.1:9000 127.0.0.1:7001 127.0.0.1:8001 127.0.0.1:9001 --cluster-replicas 1

----

redis-cli cluster nodes
`
cluster id | address | description | slots
7d43b492587383efe2332cc6a3b8e995a7597026 127.0.0.1:9000@19000 master - 0 1636795962750 3 connected 10923-16383
51d1de9c6f4da983ffb53089651e7fd113226b4d 127.0.0.1:7000@17000 myself,master - 0 1636795960000 1 connected 0-5460
7e91b31997f19c9929b4a742a4bc3526e77fc9bb 127.0.0.1:8000@18000 master - 0 1636795961725 2 connected 5461-10922
`

If you want to add new nodes first
`
redis-cli --cluster add-node 127.0.0.1:10000 127.0.0.1:11000
`

Resharded
`
redis-cli --cluster reshard 127.0.0.1:7000
`

Replica 7001
`
redis-cli --cluster add-node 127.0.0.1:7001 127.0.0.1:7000 --cluster-slave --cluster-master-id 51d1de9c6f4da983ffb53089651e7fd113226b4d
`

Replica 8001
`
redis-cli --cluster add-node 127.0.0.1:8001 127.0.0.1:8000 --cluster-slave --cluster-master-id 7e91b31997f19c9929b4a742a4bc3526e77fc9bb
`

Replica 9001
`
redis-cli --cluster add-node 127.0.0.1:9001 127.0.0.1:9000 --cluster-slave --cluster-master-id 7d43b492587383efe2332cc6a3b8e995a7597026
`