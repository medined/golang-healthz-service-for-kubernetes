# Golang Healthz Service For Kubernetes

The smallest image possible (?) for a Kubernetes Healthz Service.

# Why

When `kubeadm` is initialized, it looks for a local server as a health check. The image created by this project will respond to those health checks. It is written using Go to make a 27mb image.

# Prequisite

* Define the image and tag.

```bash
export DHUSER="medined"
export INAME="healthz-service"
export ITAG="0.0.1"
export IMAGE="$INAME:$ITAG"
```

# Using With Fedora CoreOS

Different part of Kubernetes use different ports for their health check. My knowledge is an inch deep so don't on information here being 100% accurate.

* 10248 - used by kubeadm during initialization.

Use the following as a template for your `fcc` file. It will create healthz service that responds on port 10248.

```bash
cat <<EOF > example.fcc
variant: fcos
version: 1.0.0
- name: healthz.service for kubeadm
  enabled: true
  contents: |
    [Unit]
    Description=A healthz unit!
    After=network-online.target
    Wants=network-online.target
    [Service]
    Type=forking
    KillMode=none
    Restart=on-failure
    RemainAfterExit=yes
    ExecStartPre=podman pull $IMAGE
    ExecStart=podman run -d --name healthz-server -e PORT=10248 -p 10248:10248 $IMAGE
    ExecStop=podman stop -t 10 healthz-server
    ExecStopPost=podman rm healthz-server
    [Install]
    WantedBy=multi-user.target
```

# What To Do

* Download the go binary.

```bash
mkdir -p $HOME/bin
curl -o go.tgz -L https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz
tar xf go.tgz -C $HOME/bin
rm go.tgz

cat <<EOF | sudo tee /etc/profile.d/go.sh
export PATH=$PATH:$HOME/bin/go/bin
EOF
```

* Build the image.

```bash
docker build -t $IMAGE .
```

* Run inside container.

```
PORT=10258
docker run -it --rm=true -e PORT=$PORT -p $PORT:$PORT $IMAGE
```

* Push Image

```bash
docker login
docker tag $IMAGE $DHUSER/$IMAGE
docker push $DHUSER/$IMAGE
```
