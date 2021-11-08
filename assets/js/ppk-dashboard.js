var ws
var token
var pumMap = new Map()
var ppeMap = new Map()
var workerMap = new Map()

const APIHOST = "localhost:8080"

function initWS(){
    fetch(`http://${APIHOST}/token`,null).then(res=>{return res.json()}).then(resJSON=>{
        token = resJSON.msg

        ws = new WebSocket(`ws://${APIHOST}/websocket/ppk`)

        ws.onopen = ()=>{
            console.log("opening ws conn")
            ws.send(JSON.stringify({
                type:"initPPKFromClient",
                token:token
            }))

            ws.send(JSON.stringify({
                type:"inboxPUMFromClient",
                token:token
            }))

            ws.send(JSON.stringify({
                type:"inboxPPEFromClient",
                token:token
            }))
        
            ws.send(JSON.stringify({
                type:"inboxWorkerFromClient",
                token:token
            }))
        }

        ws.onmessage = (e)=>{
            let jsonData = JSON.parse(e.data)
            let data = jsonData.data
            let msg = jsonData.msg

            switch(jsonData.type){
                case "inboxWorkerFromServer":
                    populateInboxWorker(data)
                    break
                case "inboxPPEFromServer":
                    populateInboxPPE(data)
                    break
                case "inboxPUMFromServer":
                    populateInboxPUM(data)
                    break
            }
        }

        ws.onerror = (e) => {
            console.log(e)
        }
    
        ws.onclose = (e) => {
            alert("connection closed.")
        }
    })
}

function populateInboxPUM(pumArray){
    let tableBody = document.getElementById("table-body-pum")
    for (let i=0;i < pumArray.length;i++){
        let id = pumArray[i].ppk_inbox_from_pum_id
        let dateCreated = pumArray[i].ppk_inbox_from_pum_date_created
        let priority = pumArray[i].WorkRequest.work_request_priority
        let task = pumArray[i].WorkRequest.work_request_task
        let location = pumArray[i].WorkRequest.work_request_location
        let equipment = pumArray[i].WorkRequest.work_request_equipment
        let newRow = document.createElement("tr")
        newRow.id = `pum-${id}`
        let newRowInnerHTML = `<td>${id}</td><td>${priority}</td><td>${dateCreated}</td><td>${task}</td><td>${location}</td><td>${equipment}</td><td><button>Detail</button><button>Respon</button></td>`
        newRow.innerHTML = newRowInnerHTML
        pumMap.set(id,pumArray[i])
        tableBody.appendChild(newRow)
    }
}

function populateInboxPPE(ppeArray){
    let tableBody = document.getElementById("table-body-ppe")
    for (let i=0;i<ppeArray.length;i++){
        let id = ppeArray[i].ppk_inbox_from_ppe_id
        let task = ppeArray[i].WorkRequest.work_request_task
        let estDate = ppeArray[i].ppk_inbox_from_ppe_est_date
        let worker = ppeArray[i].ppk_inbox_from_ppe_worker
        let workerEmail = ppeArray[i].ppk_inbox_from_ppe_worker_email
        let cost = ppeArray[i].ppk_inbox_from_ppe_cost
        let newRow = document.createElement("tr")
        newRow.id = `pum-${id}`
        let newRowInnerHTML = `<td>${id}</td><td>${task}</td><td>${estDate}</td><td>${worker}</td><td>${workerEmail}</td><td>${cost}</td><td><button>Detail</button><button>Buat WO</button></td>`
        newRow.innerHTML = newRowInnerHTML
        ppeMap.set(id,ppeArray[i])
        tableBody.appendChild(newRow)
    }
}

function populateInboxWorker(workerArray){
    let tableBody = document.getElementById("table-body-worker")
    for (let i=0;i<workerArray.length;i++){
        let id = workerArray[i].ppk_inbox_from_worker_id
        let dateCreated = workerArray[i].ppk_inbox_from_worker_date_created
        let task = workerArray[i].WorkOrder.WorkRequest.work_request_task
        let equipment = workerArray[i].WorkOrder.WorkRequest.work_request_equipment
        let location = workerArray[i].WorkOrder.WorkRequest.work_request_location
        let worker = workerArray[i].WorkOrder.work_order_worker
        let workerEmail = workerArray[i].WorkOrder.work_order_worker_email
        let newRow = document.createElement("tr")
        newRow.id = `pum-${id}`
        let newRowInnerHTML = `<td>${id}</td><td>${dateCreated}</td><td>${task}</td><td>${equipment}</td><td>${location}</td><td>${worker}</td><td>${workerEmail}</td><td><button>Terima</button><button>Revisi</button></td>`
        newRow.innerHTML = newRowInnerHTML
        workerMap.set(id,workerArray[i])
        tableBody.appendChild(newRow)
    }
}

initWS()