var ws
var token
var inboxMap = new Map()
const APIHOST = "localhost:8080"

function initWS(){
    //TOBE
    fetch(`http://${APIHOST}/token`,null).then(res=>{return res.json()}).then(resJSON=>{
        token = resJSON.msg
        ws = new WebSocket(`ws://${APIHOST}/websocket/ppe`)

        ws.onopen = (e) => {
            ws.send(JSON.stringify({
                type:"initPPEFromClient",
                token:token
            }))
        }

        ws.onmessage = (e) => {
            let jsonData = JSON.parse(e.data)
            let data = jsonData.data
            let msg = jsonData.msg
            switch (jsonData.type){
                case "initPPEFromServer":
                    alert(JSON.stringify(data))
                    break
            }
        }
    })
}

initWS()