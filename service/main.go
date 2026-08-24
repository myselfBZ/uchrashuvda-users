package service

// import (
// 	"fmt"
// 	"net"
//
// 	_ "github.com/joho/godotenv/autoload"
// 	_ "github.com/lib/pq"
// 	pb "github.com/myselfBZ/uchrashuvda-isc/users"
// 	"github.com/myselfBZ/user-service/env"
// 	"github.com/myselfBZ/user-service/store"
// 	"google.golang.org/grpc"
// )
//
// type cfg struct {
// 	port int
// 	db   store.DBConfig
// }
//
// func (c *cfg) Load() {
// 	c.port 	  = env.GetInt("SERVICE_PORT", 6767)
// 	c.db.Addr = env.MustGetString("DB")
// 	c.db.MaxIdleConns = env.GetInt("DB_MAX_IDLE_CONNS", 30)
// 	c.db.MaxOpenConns = env.GetInt("DB_MAX_OPEN", 30)
// 	c.db.MaxIdleTime = env.GetString("DB_MAX_IDLE_TIME", "15m")
// }
//
// func main() {
// 	server := grpc.NewServer()
// 	c := &cfg{}
// 	c.Load()
// 	db, err := store.NewDBConn(c.db)
// 	if err != nil {
// 		panic(err)
// 	}
// 	store := store.NewPostgreStore(db)
// 	pb.RegisterUserServiceServer(server, &service{store: store})
// 	ln, err := net.Listen("tcp", fmt.Sprintf(":%d",c.port))
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Let's fuck...")
// 	if err := server.Serve(ln); err != nil {
// 		panic(err)
// 	}
// }
