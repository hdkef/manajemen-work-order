var ws
var token

function initWS(){
   //TOBE
   fetch("http://localhost:8080/before-ws",null).then(res=>{return res.json()}).then(resJSON=>{
       token = resJSON.msg
       ws = new WebSocket("ws://localhost:8080/websocket")

       ws.onopen = (e) => {
           ws.send(
               JSON.stringify({
                   type:"initUserFromClient",
                   token:token
               })
           )
       }
   
       ws.onmessage = (e) => {
           let jsonData = JSON.parse(e.data)
           console.log("jsonData",jsonData)
           switch (jsonData.type){
               case "error":
                   alert(jsonData.data)
                   break
           }
       }
   
       ws.onerror = (e) => {
           console.log(e)
       }
   
       ws.onclose = (e) => {
           console.log(e,"onclose")
       }
   })
}

initWS()