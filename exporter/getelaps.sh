#! /bin/sh

# set monitor name
mon="genw"

tail -n 1 /opt/nec/clusterpro/perf/cluster/monitor/$mon.cur |awk -F',' '{print $8}'|tr -d '"'

