#!/bin/bash

echo "Creating redis cluster custom, a master with one replica for each.."
sleep 1.5

echo "Let's continue to the steps.."
sleep 1

create_redis_cluster() {
    numb=$1
    port_master=$2
    port_replica=$3

    echo "Cluster $numb : Master redis with port $port_master with redis replica port $port_replica"
    echo "Checking port $port_master..."
    result=$(lsof -i -P -n | grep redis | grep LISTEN | grep IPv4 | grep :$port_master)

    # if result is empty then we can use the port
    if [[ -z $result ]]; then
        echo "Port $port_master is available, starting redis master $port_master with config/redis-master-$port_master.conf"
        config_master="config/redis-master-$port_master.conf"
        redis-server $config_master
        result=$(lsof -i -P -n | grep redis | grep LISTEN | grep IPv4 | grep :$port_master)
        if [[ -z $result ]]; then
            echo "Failed to start redis.."
            exit
        else
            echo "Succeed!!! continue..."
        fi
    else
        echo "Port is in use, please kill application that using the port and start this script again.."
        exit
    fi

    echo "Checking port $port_replica..."
    result=$(lsof -i -P -n | grep redis | grep LISTEN | grep IPv4 | grep :$port_replica)
    # if result is empty then we can use the port
    if [[ -z $result ]]; then
        echo "Port $port_replica is available, starting redis master $port_replica with config/redis-replica-$port_replica.conf"
        config_replica="config/redis-master-$port_replica.conf"
        redis-server $config_replica
        result=$(lsof -i -P -n | grep redis | grep LISTEN | grep IPv4 | grep :$port_replica)
        if [[ -z $result ]]; then
            echo "Failed to start redis.."
            exit
        else
            echo "Succeed!!! continue..."
        fi
    else
        echo "Port is in use, please kill application that using the port and start this script again.."
        exit
    fi
}

create_redis_cluster "1" "7000" "7001"
create_redis_cluster "2" "8000" "8001"
create_redis_cluster "3" "9000" "9001"