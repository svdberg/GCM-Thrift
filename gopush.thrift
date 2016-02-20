namespace go gopush

struct Payload {
  1:string message
}

struct GcmRequest {
  1:string to
  2:Payload data
}

service Gopush {
  void message(1:GcmRequest m)
}
