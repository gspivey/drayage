 # Drayage

No straightforward method to deal with docker volume migration exists, therefore
[Drayage](https://en.wikipedia.org/wiki/Drayage) was created. 
Current solutions require Kubernetes, custom scripts, or a shared storage 
mechanism, however Drayage is a docker volume migration tool meant to 
dynamically move small volumes at container start time. This project shares many
of the deliverables of Flocker, however this project has a smaller scope of 
deployment.

## Manifesto
1. Generally, local storage is fast (SSD or multi-drive raid array) and availble.
   Drayage aims to leverage the throughput available with local storage.
2. Volume migration should be agnostic to container orchestration topologies.
3. Storage migrations are set to a limited size so in order to minimize
   container start times.
  - Use cases are limited: filesystems with large object counts should be 
    managed separately. For example, databases should be migrated and managed
    with more robust database management tools.
4. Manual volume moves work with any size volume

## Goal

* **Local Datacenter, Local Storage**: Dynamic volume migration is intended to 
  occur within a datacenter or low-latency network cluster. Higher latency
  conections should be clustered separately in other datacenters.

* **Simple Storage**: Drayage does not requre a complicated storage scheme - 
  storage local to the cluster is king.

* **Fast**: Using the recommended 5GB limit with concurrent copy support will
  allow volume migration to support quick deployments of newly running nodes. 

* **Encryption Made Easy**: All communication is encrypted and mutually
  authenticated. Drayage leverages built in certificate generation to aid in
  quick, secure deployments.

## Versioning
Drayage follows [semver](https://semver.org/spec/v2.0.0.html) (Semantic Versioning)
Public API's will not be stable until the 1.0.0 relesase

## Install Dependencies

1. gcc
2. protoc v3.7.1
3. Docker v18.09.5

## Roadmap

### Pre 1.0.0 Features
1. CLI volume copy of any size
2. CLI tls initialization for all configured nodes
3. Dynamic volume migration for volume's smaller then 5GB

### 1.X.X Maintinence Enhancements
1. Automatic Node Configuration
2. File Streaming optimization
3. /etc/hosts x.drayage "dns"
4. Support for consul/etcD volume state information storage
5. Support configurtion for dynamic migration of volume's larger than 5GB (not recommended)

### X.O.O Features
[TBD]

##  FAQ
1. I see you're using gRPC. Why did you use a separate TCP file streamer? 
  gRPC is handling the Control Protocol only. It could be used for managing the 
  data stream as well however there would be multiple SerDes
  (Serialize/Desearialize Marshall/Un-Marshall) events. Instead steaming read, 
  transport, and writes are all done in bytes followed by a checksum to ensure
  transfer integrity.

2. Why a 5GB dynamic transfer limit? 
  Drayage is focused on speed - by keeping the container storage migration to a 
  5GB limit, we can ensure speedy deployments. 5GB is configurable, however 
  it is at your own risk.

3. How do I get a new feature into the the next/future major release?
  Read the contribution guidelines. Interact with the development team.
  When the branch for the next version is created submit a pull request to get
  it in faster.
