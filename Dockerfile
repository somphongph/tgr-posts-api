FROM alpine:3.19
COPY ./out/goapp /app/goapp
# COPY ./configs ./configs

RUN adduser -D -u 1000 appuser
USER appuser

CMD ["/app/goapp"]
