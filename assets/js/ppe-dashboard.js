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
                    populateInbox(data)
                    break
            }
        }
    })
}

//to populate or create new row of inbox
function populateInbox(inboxArray){
    let tableBody = document.getElementById("table-body-inbox")
    for (let i =0;i < inboxArray.length;i++){
        let id = inboxArray[i].ppe_inbox_id
        let priority = inboxArray[i].WorkRequest.work_request_priority
        let est_cost = inboxArray[i].ppe_inbox_est_cost
        let date_created = inboxArray[i].WorkRequest.work_request_date_created
        let task = inboxArray[i].WorkRequest.work_request_task
        let location = inboxArray[i].WorkRequest.work_request_location
        let equipment = inboxArray[i].WorkRequest.work_request_equipment
        let newRow = document.createElement("tr")
        newRow.id = `inbox-${id}`
        let newRowInnerHTML = `<td>${id}</td><td>${priority}</td><td>${date_created}</td><td>${task}</td><td>${est_cost}</td><td>${location}</td><td>${equipment}</td><td><button>Detail</button></td>`
        newRow.innerHTML = newRowInnerHTML
        inboxMap.set(id,inboxArray[i])
        tableBody.appendChild(newRow)
    }
}

initWS()