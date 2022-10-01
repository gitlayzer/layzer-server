FROM busybox

COPY layzer-server  /root/
WORKDIR /root/

ENTRYPOINT ["./layzer-server"]

EXPOSE 8080
