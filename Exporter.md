# How to Deploy Exporter for EXPRESSCLUSTER X

## Overview
- EXPRESSCLUSTER X and clpexporter run on the same server.
  ```
                  +--------------+
                  | CentOS       |
                  | - Prometheus |
                  +------+-------+
                         |
          +--------------+---------+
          |                        |
  +-------+------------+   +-------+------------+
  | CentOS             |   | CentOS             |
  | - EXPRESSCLUSTER X |   | - EXPRESSCLUSTER X |
  | - clpexporter      |   | - clpexporter      | 
  +--------------------+   +--------------------+
  ```

## Run Exporter for EXPRESSCLUSTER
1. Install EXPRESSCLUSTER and create a cluster.
1. Download clpexporter.
1. Run clpexporter.
1. Edit prometheus.yml file asbelow and run Prometheus.
   ```yaml
       static_configs:
               - targets: ['<IP address of the primary server>:29090', '<IP address of the secondary server>:29090']   
   ```
1. Start web browser and access the following address.
   ```
   htt://<IP Address of Prometheus>:9090
   ```
1. Select **clp_monitor_(monitor resoruce name)** and click execute. You can get the elapsed time [msec] of the monitor resource.

## Reference
- Prometheus Exporter
  - https://prometheus.io/docs/instrumenting/writing_exporters/
- Sample Exporter
  - https://github.com/hirano00o/sample-exporter