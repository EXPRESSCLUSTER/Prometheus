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
## Reference
- Prometheus
  - https://prometheus.io/docs/instrumenting/writing_exporters/
- Qiita (Japanese Web Site)
  - https://qiita.com/ryojsb/items/256f1d205a83ae772f39#comment-620e627d064f1f19bcd8
