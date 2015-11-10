FROM gliderlabs/alpine:3.2
ENTRYPOINT ["/bin/go_server_linux_amd64"]

COPY bin/go_server_linux_amd64 /bin/
COPY public /public
