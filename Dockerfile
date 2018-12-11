FROM scratch
EXPOSE 8080
ENTRYPOINT ["/medbo-backend"]
COPY ./bin/ /