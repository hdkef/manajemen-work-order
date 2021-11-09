var ws
var token
var inboxMap = new Map()
var detailModal = document.getElementById("detail-modal")
var respondModal = document.getElementById("respond-modal")
var changepasswordmodal = document.getElementById("change-password-modal")
var historyMap = new Map()
var historyLastID = 0

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
                case "loadHistoryPPEFromServer":
                    populateHistory(data)
                    break
                case "changePasswordFromServer":
                   alert(msg)
                   break
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
        let priority = inboxArray[i].WorkRequest.priority
        let est_cost = inboxArray[i].ppe_inbox_est_cost
        let date_created = inboxArray[i].WorkRequest.date_created
        let task = inboxArray[i].WorkRequest.task
        let location = inboxArray[i].WorkRequest.location
        let equipment = inboxArray[i].WorkRequest.equipment
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
    let wrid = inboxTmp.WorkRequest.id
    let estCost = inboxTmp.ppe_inbox_estCost
    let priority = inboxTmp.WorkRequest.priority
    let dateCreated = inboxTmp.WorkRequest.date_created
    let task = inboxTmp.WorkRequest.task
    let location = inboxTmp.WorkRequest.location
    let equipment = inboxTmp.WorkRequest.equipment
    let instruction = inboxTmp.WorkRequest.instruction
    let description = inboxTmp.WorkRequest.description
    var newHTML = `<h3>Inbox ID<h3><h4>${id}</h4><h3>ID Pembuat Work Order<h3><h4>${wrid}</h4><h3>Prioritas<h3><h4>${priority}</h4><h3>Tanggal Dibuat</h3><h3>${dateCreated}</h4><h3>Pekerjaan</h3><h4>${task}</h4><h3>Lokasi</h3><h4>${location}</h4><h3>Nama / Tag Alat</h3><h4>${equipment}</h4><h3>Perkiraan Biaya</h3><h4>${estCost}</h4><h3>Instruksi</h3><p>${instruction}</p><h3>Deskripsi</h3><p>${description}</p>`
    detail.innerHTML = newHTML
    detailModal.style.display = "block"
}

function showHistoryDetail(id){
    let obj = historyMap.get(id)
    let wrid = obj.WorkRequest.id
    let priority = obj.WorkRequest.priority
    let dateCreated = obj.WorkRequest.date_created
    let task = obj.WorkRequest.task
    let estCost = obj.ppe_inbox_est_cost
    let location = obj.WorkRequest.location
    let equipment = obj.WorkRequest.equipment
    let status = obj.WorkRequest.status
    let instruction = obj.WorkRequest.instruction
    let description = obj.WorkRequest.description
    var newHTML = `<h3>Inbox ID<h3><h4>${id}</h4><h3>ID Pembuat Work Order<h3><h4>${wrid}</h4><h3>Status</h3><h4>${status}</h4><h3>Prioritas<h3><h4>${priority}</h4><h3>Tanggal Dibuat</h3><h3>${dateCreated}</h4><h3>Pekerjaan</h3><h4>${task}</h4><h3>Lokasi</h3><h4>${location}</h4><h3>Nama / Tag Alat</h3><h4>${equipment}</h4><h3>Perkiraan Biaya</h3><h4>${estCost}</h4><h3>Instruksi</h3><p>${instruction}</p><h3>Deskripsi</h3><p>${description}</p>`
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

function showChangePassword(){
    changepasswordmodal.style.display = "block"
}

function closeChangePassword(){
    changepasswordmodal.style.display = "none"
}

function changePwd(){
    let oldPassword = document.getElementById("input-old-password").value
    let newPassword = document.getElementById("input-new-password").value
    //TOBE VALIDATION
    ws.send(JSON.stringify({
        type:"changePasswordFromClient",
        token:token,
        changepwdfromclient:{
            old:oldPassword,
            new:newPassword
        }
    }))
    closeChangePassword()
}

function loadHistory(){
    for (let [key,_] of historyMap.entries()){
        if (+key > historyLastID){
            historyLastID = +key
        }
    }
    ws.send(JSON.stringify({
        type:"loadHistoryPPEFromClient",
        token:token,
        last_id:historyLastID,
    }))
}

function populateHistory(historyArray){
    let tableBody = document.getElementById("table-body-history")
    for (let i =0;i < historyArray.length;i++){
        let id = historyArray[i].ppe_inbox_id
        let priority = historyArray[i].WorkRequest.priority
        let date_created = historyArray[i].WorkRequest.date_created
        let task = historyArray[i].WorkRequest.task
        let estCost = historyArray[i].ppe_inbox_est_cost
        let location = historyArray[i].WorkRequest.location
        let equipment = historyArray[i].WorkRequest.equipment
        let status = historyArray[i].WorkRequest.status
        let newRow = document.createElement("tr")
        newRow.id = `history-${id}`
        //tobe
        let newRowInnerHTML = `<td>${id}</td><td>${status}</td><td>${priority}</td><td>${date_created}</td><td>${task}</td><td>${estCost}</td><td>${location}</td><td>${equipment}</td><td><button onclick="showHistoryDetail(${id})">Detail</button></td>`
        newRow.innerHTML = newRowInnerHTML
        historyMap.set(id,historyArray[i])
        tableBody.appendChild(newRow)
    }
}

initWS()