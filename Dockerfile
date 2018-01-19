FROM debian:stable-slim
COPY ./who-is-there /who-is-there
RUN chmod +x /who-is-there
EXPOSE 8080
ENTRYPOINT ["/who-is-there"]
