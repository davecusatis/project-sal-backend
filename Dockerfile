FROM alpine:3.4

RUN apk -U add ca-certificates

EXPOSE 3030

ADD ./project-sal /bin/project-sal

CMD ["project-sal"]
