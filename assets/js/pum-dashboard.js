var ws
var token
var inboxMap = new Map()
var historyMap = new Map()
var detailModal = document.getElementById("detail-modal")
var historyModal = document.getElementById("history-modal")
var detail = document.getElementById("detail")
var historyDetail = document.getElementById("history-detail")
var changepasswordmodal = document.getElementById("change-password-modal")
var historyLastID = 0

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
                case "loadHistoryPUMFromServer":
                    populateHistory(data)
                    break
                case "changePasswordFromServer":
                   alert(msg)
                   break
                case "resWRPUMFromServer":
                    destroyInbox(data)
                    alert(msg)
                    break
                case "initPUMFromServer":
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

function loadHistory(){
    for (let [key,_] of historyMap.entries()){
        if (+key > historyLastID){
            historyLastID = +key
        }
    }
    ws.send(JSON.stringify({
        type:"loadHistoryPUMFromClient",
        token:token,
        last_id:historyLastID,
    }))
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

function destroyInbox(id){
    inboxMap.delete(+id)
    let inbox = document.getElementById(`inbox-${id}`)
    inbox.parentNode.removeChild(inbox)
}

function closeModalDetail(){
    detailModal.style.display = "none"
}

function acceptWR(id){
    ws.send(JSON.stringify({
        type:"acceptWRPUMFromClient",
        token:token,
        pumrespondfromclient:{id:+id}
    }))
    closeModalDetail()
}

function declineWR(id){
    ws.send(JSON.stringify({
        type:"declineWRPUMFromClient",
        token:token,
        pumrespondfromclient:{id:+id}
    }))
    closeModalDetail()
}

function populateHistory(historyArray){
    let tableBody = document.getElementById("table-body-history")
    for (let i =0;i < historyArray.length;i++){
        let id = historyArray[i].pum_inbox_id
        let priority = historyArray[i].WorkRequest.priority
        let date_created = historyArray[i].WorkRequest.date_created
        let task = historyArray[i].WorkRequest.task
        let location = historyArray[i].WorkRequest.location
        let equipment = historyArray[i].WorkRequest.equipment
        let status = historyArray[i].WorkRequest.status
        let newRow = document.createElement("tr")
        newRow.id = `history-${id}`
        let newRowInnerHTML = `<td>${id}</td><td>${status}</td><td>${priority}</td><td>${date_created}</td><td>${task}</td><td>${location}</td><td>${equipment}</td><td><button onclick="showHistoryDetail(${id})">Detail</button></td>`
        newRow.innerHTML = newRowInnerHTML
        historyMap.set(id,historyArray[i])
        tableBody.appendChild(newRow)
    }
}

//to populate or create new row of inbox
function populateInbox(inboxArray){
    let tableBody = document.getElementById("table-body-inbox")
    for (let i =0;i < inboxArray.length;i++){
        let id = inboxArray[i].pum_inbox_id
        let priority = inboxArray[i].WorkRequest.priority
        let date_created = inboxArray[i].WorkRequest.date_created
        let task = inboxArray[i].WorkRequest.task
        let location = inboxArray[i].WorkRequest.location
        let equipment = inboxArray[i].WorkRequest.equipment
        let newRow = document.createElement("tr")
        newRow.id = `inbox-${id}`
        let newRowInnerHTML = `<td>${id}</td><td>${priority}</td><td>${date_created}</td><td>${task}</td><td>${location}</td><td>${equipment}</td><td><button onclick="showDetail(${id})">Detail</button></td>`
        newRow.innerHTML = newRowInnerHTML
        inboxMap.set(id,inboxArray[i])
        tableBody.appendChild(newRow)
    }
}

function showHistoryDetail(id){
    let obj = inboxMap.get(id)
    let priority = obj.WorkRequest.priority
    let wrid = obj.WorkRequest.id
    let dateCreated = obj.WorkRequest.date_created
    let task = obj.WorkRequest.task
    let location = obj.WorkRequest.location
    let equipment = obj.WorkRequest.equipment
    let status = obj.WorkRequest.status
    let instruction = obj.WorkRequest.instruction
    let description = obj.WorkRequest.description
    var newHTML = `<h3>Inbox ID<h3><h4>${id}</h4><h3>ID Pembuat Work Order<h3><h4>${wrid}</h4><h3>Status</h3><h4>${status}</h4><h3>Prioritas<h3><h4>${priority}</h4><h3>Tanggal Dibuat</h3><h3>${dateCreated}</h4><h3>Pekerjaan</h3><h4>${task}</h4><h3>Lokasi</h3><h4>${location}</h4><h3>Nama / Tag Alat</h3><h4>${equipment}</h4><h3>Instruksi</h3><p>${instruction}</p><h3>Deskripsi</h3><p>${description}</p>`
    historyDetail.innerHTML = newHTML
    historyModal.style.display = "block"
}

function closeHistoryDetail(){
    historyModal.style.display = "none"
}

function showDetail(id){
    let inboxTmp = inboxMap.get(id)
    let wrid = inboxTmp.WorkRequest.id
    let priority = inboxTmp.WorkRequest.priority
    let date_created = inboxTmp.WorkRequest.date_created
    let task = inboxTmp.WorkRequest.task
    let location = inboxTmp.WorkRequest.location
    let equipment = inboxTmp.WorkRequest.equipment
    let instruction = inboxTmp.WorkRequest.instruction
    let description = inboxTmp.WorkRequest.description
    let btnAccept = document.getElementById("button-accept")
    var btnDecline = document.getElementById("button-decline")
    var newHTML = `<h3>Inbox ID<h3><h4>${id}</h4><h3>ID Pembuat Work Order<h3><h4>${wrid}</h4><h3>Prioritas<h3><h4>${priority}</h4><h3>Tanggal Dibuat</h3><h3>${date_created}</h4><h3>Pekerjaan</h3><h4>${task}</h4><h3>Lokasi</h3><h4>${location}</h4><h3>Nama / Tag Alat</h3><h4>${equipment}</h4><h3>Instruksi</h3><p>${instruction}</p><h3>Deskripsi</h3><p>${description}</p>`
    detail.innerHTML = newHTML
    detailModal.style.display = "block"
    btnAccept.onclick = ()=>{
        acceptWR(+id)
    }
    btnDecline.onclick = ()=>{
        declineWR(+id)
    }
}

initWS()