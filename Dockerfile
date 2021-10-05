FROM alpine:3
COPY bin/larvis /bin/larvis
ENTRYPOINT [ "/bin/larvis" ]
