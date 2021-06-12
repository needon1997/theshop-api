package global

import "google.golang.org/grpc"

var OrderSvcConn *grpc.ClientConn
var GoodsSvcConn *grpc.ClientConn
var InventoryConn *grpc.ClientConn
var PaymentSvcConn *grpc.ClientConn
