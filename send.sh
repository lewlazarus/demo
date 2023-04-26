#!/bin/sh

nats request pudaf.math '{"operation": "Sum", "values": [1,2,3,4,5,6,7,8,9,10]}'
nats request pudaf.math '{"operation": "Multiply", "values": [2,1,1,1,1,1,1,1,1,1]}'
nats request pudaf.math '{"operation": "Sum", "values": [1,2,3,4,5,6,7,8,9,10]}'
nats request pudaf.math '{"operation": "Sum", "values": [1,2,3,4,5,6,7,8,9,10]}'
nats request pudaf.math '{"operation": "Sum", "values": [1,2,3,4,5,6,7,8,9,10]}'
nats request pudaf.math '{"operation": "Sum", "values": [1,2,3,4,5,6,7,8,9,10]}'
nats request pudaf.math '{"operation": "Multiply", "values": [2,2,1,1,1,1,1,1,1,1]}'
nats request pudaf.math '{"operation": "Sum", "values": [1,2,3,4,5,6,7,8,9,10]}'
nats request pudaf.math '{"operation": "Sum", "values": [1,2,3,4,5,6,7,8,9,10]}'
nats request pudaf.math '{"operation": "Sum", "values": [1,2,3,4,5,6,7,8,9,10]}'
nats request pudaf.math '{"operation": "Multiply", "values": [2,2,2,1,1,1,1,1,1,1]}'
nats request pudaf.math '{"operation": "Sum", "values": [1,2,3,4,5,6,7,8,9,10]}'
