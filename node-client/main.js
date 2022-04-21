const {HelloServiceClient} = require("../node-client/hello/hello_grpc_web_pb")
const {LoginRequest} = require("../node-client/hello/hello_pb")
const hello_service_client = new HelloServiceClient("http://localhost:8080")
global.XMLHttpRequest = require('xhr2');

var lr = new LoginRequest();
lr.setUsername("shane");
lr.setPassword("mossmoss");
hello_service_client.login(lr, {}, (err, res)=>{
    if(err){
        console.log(err)
        return 
    }else{
        console.log(res.getToken())
    }
})