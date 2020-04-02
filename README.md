# Prometheus

## Overview
- Prometheus and EXPRESSCLUSTER run on different servers.
  ```
                +--------------+
                | CentOS       |
                | - Prometheus |
                +------+-------+
                       |
          +------------+---------+
          |                      |
  +-------+----------+   +-------+----------+
  | CentOS           |   | CentOS           |
  | - EXPRESSCLUSTER |   | - EXPRESSCLUSTER |
  +------------------+   +------------------+
  ```

## Run Exporter for EXPRESSCLUSTER
1. Install EXPRESSCLUSTER and create a cluster in advance.
1. Clone go source file and shell script.
   ```sh
   git clone https://github.com/EXPRESSCLUSTER/Prometheus.git
   ```
1. Change the following parameter in getelaps.sh to match your environment.
   ```sh
   # set monitor name
   mon="genw" 
   ```
1. 
1. Run the following command with root account.
   ```sh
   # go run main.go
   ```

## Run Prometheus
1. Download Prometheus binary file from https://prometheus.io/download/.
1. Expand the file.
   ```sh
   $ tar xzvf prometheus-2.17.1.linux-amd64.tar.gz
   ```
1. Move to prometheus directory.
   ```sh
   $ cd prometheus-2.17.1.linux-amd64
   ```
1. Open prometheus.yml and add exporter IP address and port number.
   ```yaml
       static_configs:
       - targets: ['<IP Address of the Primary Server>:9090', '<IP Address of the Secondary Server>:9090']
   ```
1. Run Prometheus.
   ```sh
   $ ./prometheus --config.file=prometheus.yml
   ```
1. Start web browser and access the following address.
   ```
   htt://<IP Address of Prometheus>:9090
   ```
1. Select **ecx_monitor_elapsed_time** and click execute. You can get the elapsed time [msec] of the monitor resource.

## Reference
- Prometheus
  - https://prometheus.io/docs/instrumenting/writing_exporters/
- Qiita (written in Japanese)
  - https://qiita.com/ryojsb/items/256f1d205a83ae772f39
