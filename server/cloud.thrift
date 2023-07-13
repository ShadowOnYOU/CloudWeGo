namespace go server

struct College {
    1: required string name(go.tag = 'json:"name"'),
    2: string address(go.tag = 'json:"address"'),
}

struct Student {
    1: required i32 id,
    2: required string name,
    3: required College college,
    4: optional list<string> email,
}

struct RegisterResp {
    1: bool success,
    2: string message,
}

struct QueryReq {
    1: required i32 id
}


//----------------------service-------------------
service StudentService {
    RegisterResp Register(1: Student info)
    Student Query(1: QueryReq req)
}