FROM    alphine:3.5
COPY    bugu-linux-amd64 /bugu/bugu
COPY    conf /bugu/conf/
COPY    static /bugu/static/
COPY    views /bugu/views/
COPY    swagger /bugu/swagger/

RUN apk add --update ca-certificates

EXPOSE 19881
WORKDIR /bugu
CMD ["/bugu/bugu"]

