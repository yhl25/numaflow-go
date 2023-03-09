####################################################################################################
# base
####################################################################################################
FROM alpine:3.12.3 as base
RUN apk update && apk upgrade && \
    apk add ca-certificates && \
    apk --no-cache add tzdata

COPY dist/generate-ten-example /bin/generate-ten-example
RUN chmod +x /bin/generate-ten-example

####################################################################################################
# generate-ten
####################################################################################################
FROM scratch as generate-ten
ARG ARCH
COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=base /bin/generate-ten-example /bin/generate-ten-example
ENTRYPOINT [ "/bin/generate-ten-example" ]
