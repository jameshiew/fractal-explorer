FROM golang:1.13.7-buster

RUN DEBIAN_FRONTEND=noninteractive \
    apt-get update && \
    apt-get install -y --no-install-recommends \
        xorg-dev=1:7.7+19 \
        libglu1-mesa-dev=9.0.0-2.1+b3 && \
    rm -rf /var/lib/apt/lists/*
