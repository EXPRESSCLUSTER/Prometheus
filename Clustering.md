# Prometheus Clustering

## Overview
- EXPRESSCLUSTER X controls Prometheus.
- On the active server, Prometheus is running.
- Prometheus data and configuratoin file are saved on the mirror disk of EXPRESSCLUSTER.
  ```
                  +---------+
                  | Grafana |
                  +----+----+
                       |
           +-----------+            
           |
  +--------+----------+ +-------------------+
  | server1 (Active)  | | server2 (Standby) |
  | Prometheus        | | Prometheus        |
  | EXPRESSCLUSTER X  | | EXPRESSCLUSTER X  |
  +--------+----------+ +-------------------+
           |
           +-----------+
                       |
          +------------+------------+   
          | Monitoring Target       |
          | e.g. Kubernetes Cluster |
          +-------------------------+   
  ```
## Evaluation Configuration
- Oralce Linux 8.2
- Prometheus 2.21.0
- EXPRESSCLUSTER X 4.2

## How to Create Cluster
1. Create a cluster that has the following resources.
   - Floating IP Address
     - IP Address: 192.168.1.11
   - Mirror Disk
     - Name: md1
     - Mount Point: /mnt/md1
1. Start the failover group on the server1.
1. Download Prometheus binary file and copy it to the server1.
   - https://prometheus.io/download/
1. Expand the archive file and rename the directory (e.g. prometheus).
1. Copy the directory to /mnt/md1 on the server1.
   ```sh
   # ll /mnt/md1
   total 20
   drwx------ 2 root root 16384 Sep 18 16:17 lost+found
   drwxr-xr-x 5 3434 3434  4096 Sep 19 07:08 prometheus
   ```
1. Start a web browser and access to Cluster WebUI (e.g. http://(IP address of server1):29003).
1. Add exec resource as below.
   |Paremeter   |Value   |
   |------------|--------|
   |Name        |exec-p8s|
   |Start script|start.sh|
   |Stop script |stop.sh |
   - start.sh
     ```sh
     #! /bin/sh
     #***************************************
     #*              start.sh               *
     #***************************************
     
     #ulimit -s unlimited
     
     cd /mnt/md1/prometheus
     ./prometheus &
     exit 0
     ```
   - stop.sh
     ```sh
     #! /bin/sh
     #***************************************
     #*               stop.sh               *
     #***************************************
     
     #ulimit -s unlimited
     
     kill -s 9 `pgrep prometheus`
     exit 0
     ```
1. Add HTTP monitor resource (*) as below.
   |Paremeter                 |Value|
   |--------------------------|-----|
   |Name                      |httpw|
   |Monitoring Timing         |Active|
   |Target Resource           |exec-p8s|
   |Connecting Destination    |localhost|
   |Port                      |9090|
   |Recovery Target           |exec-p8s|
   |Maximum Reactivation Count|1|
   |Maximum Failover Count    |1|
   - (*) You need to pruchase Internet Server Agent license.
1. Apply the configuration and start the cluster.
1. Start a web browser and access to Prometheus (e.g. http://(floating IP address):9090).