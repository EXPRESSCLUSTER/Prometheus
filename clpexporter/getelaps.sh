#! /bin/sh

# set monitor name
mon="mdw1"

tail -n 1 /opt/nec/clusterpro/perf/cluster/monitor/$mon.cur |awk -F',' '{print $8}'|tr -d '"'

