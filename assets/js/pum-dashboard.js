var ws
var token
var inboxMap = new Map()
var detailmodal = document.getElementById("detail-modal")

var APIHOST = "localhost:8080"

function initWS(){
    fetch(`http://${APIHOST}/token`,null).then(res=>{return res.json()}).then(resJSON=>{
        token = resJSON.msg
        
        ws = new WebSocket(`ws://${APIHOST}/websocket/pum`)

        ws.onopen = (e) => {
            ws.send(
                JSON.stringify({
                    type:"initPUMFromClient",
                    token:token
                })
            )
        }

        ws.onmessage = (e) => {
            let jsonData = JSON.parse(e.data)
            let data = jsonData.data
            let msg = jsonData.msg
            switch (jsonData.type){
                case "resWRPUMFromServer":
                    destroyInbox(data)
                    alert(msg)
                    break
                case "initPUMFromServer":
                    populateInbox(data,"after")
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
            console.log(e,"onclose")
        }
    })
    
}

function destroyInbox(id){
    inboxMap.delete(+id)
    let inbox = document.getElementById(`inbox-${id}`)
    inbox.parentNode.removeChild(inbox)
}

function closeModalDetail(){
    detailmodal.style.display = "none"
}

function acceptWR(id){
    ws.send(JSON.stringify({
        type:"acceptWRPUMFromClient",
        token:token,
        idfromclient:+id
    }))
    closeModalDetail()
}

function declineWR(id){
    ws.send(JSON.stringify({
        type:"declineWRPUMFromClient",
        token:token,
        idfromclient:+id
    }))
    closeModalDetail()
}

//to populate or create new row of inbox
function populateInbox(inboxArray,appendType){
    let tableBody = document.getElementById("table-body")
    for (let i =0;i < inboxArray.length;i++){
        let id = inboxArray[i].pum_inbox_id
        let priority = inboxArray[i].WorkRequest.work_request_priority
        let date_created = inboxArray[i].WorkRequest.work_request_date_created
        let task = inboxArray[i].WorkRequest.work_request_task
        let location = inboxArray[i].WorkRequest.work_request_location
        let equipment = inboxArray[i].WorkRequest.work_request_equipment
        let newRow = document.createElement("tr")
        newRow.id = `inbox-${id}`
        let newRowInnerHTML = `<td>${id}</td><td>${priority}</td><td>${date_created}</td><td>${task}</td><td>${location}</td><td>${equipment}</td><td><button onclick="showDetail(${id})">Detail</button></td>`
        newRow.innerHTML = newRowInnerHTML
        inboxMap.set(id,inboxArray[i])
        switch(appendType){
            case "after":
                tableBody.appendChild(newRow)
                break
            case "before":
                tableBody.prepend(newRow)
                break
            default:
                tableBody.appendChild(newRow)
        }
    }
}

function showDetail(id){
    let inboxTmp = inboxMap.get(id)
    let reqid = inboxTmp.WorkRequest.work_request_id
    let priority = inboxTmp.WorkRequest.work_request_priority
    let date_created = inboxTmp.WorkRequest.work_request_date_created
    let task = inboxTmp.WorkRequest.work_request_task
    let location = inboxTmp.WorkRequest.work_request_location
    let equipment = inboxTmp.WorkRequest.work_request_equipment
    let instruction = inboxTmp.WorkRequest.work_request_instruction
    let description = inboxTmp.WorkRequest.work_request_description
    let btnAccept = document.getElementById("button-accept")
    var btnDecline = document.getElementById("button-decline")
    var newHTML = `<h3>Inbox ID<h3><h4>${id}</h4><h3>ID Pembuat Work Order<h3><h4>${reqid}</h4><h3>Prioritas<h3><h4>${priority}</h4><h3>Tanggal Dibuat</h3><h3>${date_created}</h4><h3>Pekerjaan</h3><h4>${task}</h4><h3>Lokasi</h3><h4>${location}</h4><h3>Nama / Tag Alat</h3><h4>${equipment}</h4><h3>Instruksi</h3><p>${instruction}</p><h3>Deskripsi</h3><p>${description}</p>`
    detail.innerHTML = newHTML
    detailmodal.style.display = "block"
    btnAccept.onclick = ()=>{
        acceptWR(+id)
    }
    btnDecline.onclick = ()=>{
        declineWR(+id)
    }
}

initWS()