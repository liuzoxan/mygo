namespace go echo_thrift

struct Req {
    1: string msg;
}

struct Res {
    1: string msg;
}

service Echo {
    Res echo(1: Req req);
}
