FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY inspire-middleware /app

ENTRYPOINT ["/app/inspire-middleware"]

EXPOSE 50400 50401 50402 50403

CMD ["run"]
