CONNECTIONS=$1
REPLICAS=$2
IP=$3
for (( c=0; c<${REPLICAS}; c++ ))
do
   sudo docker run -v $(pwd)/client:/client --name gopher_$c -d alpine /client  -conn=${CONNECTIONS} -ip=${IP}
done