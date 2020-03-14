# Prometheus

## Exporter for EXPRESSCLUSTER
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
1. Run the following command with root account.
   ```sh
   # go run main.go
   ```
