FROM scratch
EXPOSE 8080
ENTRYPOINT ["/hcbc-backend"]
COPY ./bin/ /