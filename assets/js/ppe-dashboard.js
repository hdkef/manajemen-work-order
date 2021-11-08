var ws
var token
var inboxMap = new Map()
var detailModal = document.getElementById("detail-modal")
var respondModal = document.getElementById("respond-modal")

var detail = document.getElementById("detail")
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
                case "respondPPEFromServer":
                    destroyInbox(data)
                    break
                case "initPPEFromServer":
                    populateInbox(data)
                    break
                case "error":
                   alert(msg)
                   window.location.href = `http://${APIHOST}/login`
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
        let newRowInnerHTML = `<td>${id}</td><td>${priority}</td><td>${date_created}</td><td>${task}</td><td>${est_cost}</td><td>${location}</td><td>${equipment}</td><td><button onclick="showDetail(${id})">Detail</button><button onclick="showRespond(${id})">Respon</button></td>`
        newRow.innerHTML = newRowInnerHTML
        inboxMap.set(id,inboxArray[i])
        tableBody.appendChild(newRow)
    }
}

function closeModalDetail(){
    detailModal.style.display = "none"
}

function closeModalRespond(){
    respondModal.style.display = "none"
}

function showRespond(id){
    respondModal.style.display = "block"
    buttonRespond = document.getElementById("button-respond")
    buttonRespond.onclick = (e)=>{
        respond(id)
    }
}

function showDetail(id){
    let inboxTmp = inboxMap.get(id)
    let reqid = inboxTmp.WorkRequest.work_request_id
    let est_cost = inboxTmp.ppe_inbox_est_cost
    let priority = inboxTmp.WorkRequest.work_request_priority
    let date_created = inboxTmp.WorkRequest.work_request_date_created
    let task = inboxTmp.WorkRequest.work_request_task
    let location = inboxTmp.WorkRequest.work_request_location
    let equipment = inboxTmp.WorkRequest.work_request_equipment
    let instruction = inboxTmp.WorkRequest.work_request_instruction
    let description = inboxTmp.WorkRequest.work_request_description
    var newHTML = `<h3>Inbox ID<h3><h4>${id}</h4><h3>ID Pembuat Work Order<h3><h4>${reqid}</h4><h3>Prioritas<h3><h4>${priority}</h4><h3>Tanggal Dibuat</h3><h3>${date_created}</h4><h3>Pekerjaan</h3><h4>${task}</h4><h3>Lokasi</h3><h4>${location}</h4><h3>Nama / Tag Alat</h3><h4>${equipment}</h4><h3>Perkiraan Biaya</h3><h4>${est_cost}</h4><h3>Instruksi</h3><p>${instruction}</p><h3>Deskripsi</h3><p>${description}</p>`
    detail.innerHTML = newHTML
    detailModal.style.display = "block"
}

function respond(id){
    let estDate = new Date(document.getElementById("input-est-date").value).toISOString()
    let estLaborHour = document.getElementById("input-est-labor-hour").value
    let worker = document.getElementById("input-worker").value
    let workerEmail = document.getElementById("input-worker-email").value
    let cost = document.getElementById("input-cost").value

    ws.send(JSON.stringify({
        type:"respondPPEFromClient",
        token:token,
        pperespondfromclient:{
            id:id,
            est_date:estDate,
            est_labor_hour:+estLaborHour,
            worker:worker,
            worker_email:workerEmail,
            cost:+cost,
        }
    }))
    closeModalRespond()
}


function destroyInbox(id){
    inboxMap.delete(+id)
    let inbox = document.getElementById(`inbox-${id}`)
    inbox.parentNode.removeChild(inbox)
}

initWS()