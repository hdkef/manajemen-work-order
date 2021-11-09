var ws
var token
var pumMap = new Map()
var ppeMap = new Map()
var workerMap = new Map()
var detailModal = document.getElementById("detail-modal")
var detail = document.getElementById("detail")
var respondPUMModal = document.getElementById("respond-pum-modal")
var createWOModal = document.getElementById("create-wo-modal")
var respondWorkerModal = document.getElementById("respond-worker-modal")

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
                case "ppkRespondWorkerFromServer":
                    destroyWorker(data)
                    break
                case "createWOFromServer":
                    destroyPPE(data)
                    break
                case "ppkRespondPUMFromServer":
                    destroyPUM(data)
                    break
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

function destroyPPE(id){
    ppeMap.delete(+id)
    let inbox = document.getElementById(`ppe-${id}`)
    inbox.parentNode.removeChild(inbox)
}

function createWO(id){
    let estDate = new Date(document.getElementById("wo-input-est-date").value).toISOString()
    let estLaborHour = document.getElementById("wo-input-est-labor-hour").value
    let worker = document.getElementById("wo-input-worker").value
    let workerEmail = document.getElementById("wo-input-worker-email").value
    let cost = document.getElementById("wo-input-cost").value
    ws.send(JSON.stringify({
        type:"createWOFromClient",
        token:token,
        wofromclient:{
            est_date:estDate,
            est_labor_hour:+estLaborHour,
            worker:worker,
            worker_email:workerEmail,
            cost:+cost
        },
        ppkrespondppefromclient:{id:+id},
    }))
    closeCreateWO()
}

function destroyWorker(id){
    workerMap.delete(+id)
    let inbox = document.getElementById(`worker-${id}`)
    inbox.parentNode.removeChild(inbox)
}

function respondRevWorker(id){
    let msg = document.getElementById("worker-input-msg").value
    ws.send(JSON.stringify({
        type:"ppkRespondWorkerFromClient",
        token:token,
        ppkrespondworkerfromclient:{
            id:+id,
            msg:msg,
            type:false
        }
    }))
    closeWorkerRespond()
}

function respondAccWorker(id){
    ws.send(JSON.stringify({
        type:"ppkRespondWorkerFromClient",
        token:token,
        ppkrespondworkerfromclient:{
            id:+id,
            type:true
        }
    }))
    closeWorkerRespond()
}

function destroyPUM(id){
    pumMap.delete(+id)
    let inbox = document.getElementById(`pum-${id}`)
    inbox.parentNode.removeChild(inbox)
}

function respondPUM(id){
    let estCost = document.getElementById("pum-input-est-cost")
    ws.send(JSON.stringify({
        type:"ppkRespondPUMFromClient",
        token:token,
        ppkrespondpumfromclient:{
            id:+id,
            est_cost:+estCost
        }
    }))
    closePUMRespond()
}

function showPUMRespond(id){
    let btn = document.getElementById("pum-button-submit")
    btn.onclick = (e)=>{
        respondPUM(id)
    }
    respondPUMModal.style.display = "block"
}

function closePUMRespond(){
    respondPUMModal.style.display = "none"
}

function showCreateWO(id){
    let btn = document.getElementById("wo-button-submit")
    btn.onclick = (e)=>{
        createWO(id)
    }
    createWOModal.style.display = "block"
}

function closeCreateWO(){
    createWOModal.style.display = "none"
}

function showWorkerRespond(id){
    let buttonAcc = document.getElementById("worker-button-acc")
    let buttonRev = document.getElementById("worker-button-rev")
    buttonAcc.onclick = ()=>{
        respondAccWorker(id)
    }
    buttonRev.onclick = ()=>{
        respondRevWorker(id)
    }
    respondWorkerModal.style.display = "block"
}

function closeWorkerRespond(){
    respondWorkerModal.style.display = "none"
}

function populateInboxPUM(pumArray){
    let tableBody = document.getElementById("table-body-pum")
    for (let i=0;i < pumArray.length;i++){
        let id = pumArray[i].ppk_inbox_from_pum_id
        let dateCreated = pumArray[i].ppk_inbox_from_pum_date_created
        let priority = pumArray[i].WorkRequest.priority
        let task = pumArray[i].WorkRequest.task
        let location = pumArray[i].WorkRequest.location
        let equipment = pumArray[i].WorkRequest.equipment
        let newRow = document.createElement("tr")
        newRow.id = `pum-${id}`
        let newRowInnerHTML = `<td>${id}</td><td>${priority}</td><td>${dateCreated}</td><td>${task}</td><td>${location}</td><td>${equipment}</td><td><button onclick="showPUMDetail(${id})">Detail</button><button onclick="showPUMRespond(${id})">Respon</button></td>`
        newRow.innerHTML = newRowInnerHTML
        pumMap.set(id,pumArray[i])
        tableBody.appendChild(newRow)
    }
}

function populateInboxPPE(ppeArray){
    let tableBody = document.getElementById("table-body-ppe")
    for (let i=0;i<ppeArray.length;i++){
        let id = ppeArray[i].ppk_inbox_from_ppe_id
        let task = ppeArray[i].WorkRequest.task
        let estDate = ppeArray[i].ppk_inbox_from_ppe_est_date
        let worker = ppeArray[i].ppk_inbox_from_ppe_worker
        let workerEmail = ppeArray[i].ppk_inbox_from_ppe_worker_email
        let cost = ppeArray[i].ppk_inbox_from_ppe_cost
        let newRow = document.createElement("tr")
        newRow.id = `ppe-${id}`
        let newRowInnerHTML = `<td>${id}</td><td>${task}</td><td>${estDate}</td><td>${worker}</td><td>${workerEmail}</td><td>${cost}</td><td><button onclick="showPPEDetail(${id})">Detail</button><button onclick="showCreateWO(${id})">Buat WO</button></td>`
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
        let task = workerArray[i].WorkOrder.WorkRequest.task
        let equipment = workerArray[i].WorkOrder.WorkRequest.equipment
        let location = workerArray[i].WorkOrder.WorkRequest.location
        let worker = workerArray[i].WorkOrder.worker
        let workerEmail = workerArray[i].WorkOrder.worker_email
        let newRow = document.createElement("tr")
        newRow.id = `worker-${id}`
        let newRowInnerHTML = `<td>${id}</td><td>${dateCreated}</td><td>${task}</td><td>${equipment}</td><td>${location}</td><td>${worker}</td><td>${workerEmail}</td><td><button onclick="showWorkerDetail(${id})">Detail</button><button onclick="showWorkerRespond(${id})">Respon</buton></td>`
        newRow.innerHTML = newRowInnerHTML
        workerMap.set(id,workerArray[i])
        tableBody.appendChild(newRow)
    }
}

function showPUMDetail(id){
    let obj = pumMap.get(id)
    let wrid = obj.WorkRequest.id
    let dateCreated = obj.ppk_inbox_from_pum_date_created
    let priority = obj.WorkRequest.priority
    let task = obj.WorkRequest.task
    let location = obj.WorkRequest.location
    let equipment = obj.WorkRequest.equipment
    let instruction = obj.WorkRequest.instruction
    let description = obj.WorkRequest.description
    let newHTML =  `<h3>Work Request ID</h3><h4>${wrid}</h4><h3>Tanggal Dibuat</h3><h4>${dateCreated}</h4><h3>Prioritas</h3><h4>${priority}</h4><h3>Pekerjaan</h3><h4>${task}</h4><h3>Location</h3><h4>${location}</h4><h3>Nama / Tag Alat</h3><h4>${equipment}</h4><h3>Instruksi</h3><p>${instruction}</p><h3>Deskripsi</h3><p>${description}</p>`
    detail.innerHTML = newHTML
    detailModal.style.display = "block"
}

function showPPEDetail(id){
    let obj = ppeMap.get(id)
    let wrid = obj.WorkRequest.id
    let task = obj.WorkRequest.task
    let dateCreated = obj.ppk_inbox_from_ppe_date_created
    let estDate = obj.ppk_inbox_from_ppe_est_date
    let estLaborHour = obj.ppk_inbox_from_ppe_est_labor_hour
    let worker = obj.ppk_inbox_from_ppe_worker
    let workerEmail = obj.ppk_inbox_from_ppe_worker_email
    let cost = obj.ppk_inbox_from_ppe_cost
    let priority = obj.WorkRequest.priority
    let location = obj.WorkRequest.location
    let equipment = obj.WorkRequest.equipment
    let instruction = obj.WorkRequest.instruction
    let description = obj.WorkRequest.description
    let newHTML = `<h3>Work Request ID</h3><h4>${wrid}</h4><h3>Pekerjaan</h3><h4>${task}</h4><h3>Tanggal Dibuat</h3><h4>${dateCreated}</h4><h3>Tanggal Mulai Pekerjaan</h3><h4>${estDate}</h4><h3>Banyak Jam Kerja</h3><h4>${estLaborHour}</h4><h3>Pekerja</h3><h4>${worker}</h4><h3>Email Pekerja</h3><h4>${workerEmail}</h4><h3>Biaya</h3><h4>${cost}</h4><h3>Prioritas</h3><h4>${priority}</h4><h3>Lokasi</h3><h4>${location}</h4><h3>Alat / Nama Alat</h3><h4>${equipment}</h4><h3>Instruksi</h3><p>${instruction}</p><h3>Description</h3><p>${description}</p>`
    detail.innerHTML = newHTML
    detailModal.style.display = "block"
}

function showWorkerDetail(id){
    let obj = workerMap.get(id)
    let dateCreated = obj.ppk_inbox_from_worker_date_created
    let woid = obj.WorkOrder.id
    let estDate = obj.WorkOrder.est_date
    let estLaborHour = obj.WorkOrder.est_labor_hour
    let worker = obj.WorkOrder.worker
    let workerEmail = obj.WorkOrder.worker_email
    let cost = obj.WorkOrder.cost
    let task = obj.WorkOrder.WorkRequest.task
    let wrid = obj.WorkOrder.WorkRequest.id
    let priority = obj.WorkOrder.WorkRequest.priority
    let location = obj.WorkOrder.WorkRequest.location
    let equipment = obj.WorkOrder.WorkRequest.equipment
    let instruction = obj.WorkOrder.WorkRequest.instruction
    let description = obj.WorkOrder.WorkRequest.description
    let newHTML = `<h3>Work Request ID</h3><h4>${wrid}</h4><h3>Work Order ID</h3><h4>${woid}</h4><h3>Pekerjaan</h3><h4>${task}</h4><h3>Tanggal Dibuat</h3><h4>${dateCreated}</h4><h3>Tanggal Mulai Pekerjaan</h3><h4>${estDate}</h4><h3>Banyak Jam Kerja</h3><h4>${estLaborHour}</h4><h3>Pekerja</h3><h4>${worker}</h4><h3>Email Pekerja</h3><h4>${workerEmail}</h4><h3>Biaya</h3><h4>${cost}</h4><h3>Prioritas</h3><h4>${priority}</h4><h3>Lokasi</h3><h4>${location}</h4><h3>Alat / Nama Alat</h3><h4>${equipment}</h4><h3>Instruksi</h3><p>${instruction}</p><h3>Description</h3><p>${description}</p>`
    detail.innerHTML = newHTML
    detailModal.style.display = "block"
}

function closeModalDetail(){
    detailModal.style.display = "none"
}

initWS()