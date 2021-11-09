var ws
var token
var pumMap = new Map()
var ppeMap = new Map()
var workerMap = new Map()
var detailModal = document.getElementById("detail-modal")
var detail = document.getElementById("detail")

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
        let newRowInnerHTML = `<td>${id}</td><td>${priority}</td><td>${dateCreated}</td><td>${task}</td><td>${location}</td><td>${equipment}</td><td><button onclick="showPUMDetail(${id})">Detail</button><button>Respon</button></td>`
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
        let newRowInnerHTML = `<td>${id}</td><td>${task}</td><td>${estDate}</td><td>${worker}</td><td>${workerEmail}</td><td>${cost}</td><td><button onclick="showPPEDetail(${id})">Detail</button><button>Buat WO</button></td>`
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
        let newRowInnerHTML = `<td>${id}</td><td>${dateCreated}</td><td>${task}</td><td>${equipment}</td><td>${location}</td><td>${worker}</td><td>${workerEmail}</td><td><button onclick="showWorkerDetail(${id})">Detail</button><button>Respon</buton></td>`
        newRow.innerHTML = newRowInnerHTML
        workerMap.set(id,workerArray[i])
        tableBody.appendChild(newRow)
    }
}

function showPUMDetail(id){
    let obj = pumMap.get(id)
    let reqid = obj.WorkRequest.work_request_id
    let dateCreated = obj.ppk_inbox_from_pum_date_created
    let priority = obj.WorkRequest.work_request_priority
    let task = obj.WorkRequest.work_request_task
    let location = obj.WorkRequest.work_request_location
    let equipment = obj.WorkRequest.work_request_equipment
    let instruction = obj.WorkRequest.work_request_instruction
    let description = obj.WorkRequest.work_request_description
    let newHTML =  `<h3>Work Request ID</h3><h4>${reqid}</h4><h3>Tanggal Dibuat</h3><h4>${dateCreated}</h4><h3>Prioritas</h3><h4>${priority}</h4><h3>Pekerjaan</h3><h4>${task}</h4><h3>Location</h3><h4>${location}</h4><h3>Nama / Tag Alat</h3><h4>${equipment}</h4><h3>Instruksi</h3><p>${instruction}</p><h3>Deskripsi</h3><p>${description}</p>`
    detail.innerHTML = newHTML
    detailModal.style.display = "block"
}

function showPPEDetail(id){
    let obj = ppeMap.get(id)
    let wrid = obj.WorkRequest.work_request_id
    let task = obj.WorkRequest.work_request_task
    let dateCreated = obj.ppk_inbox_from_ppe_date_created
    let estDate = obj.ppk_inbox_from_ppe_est_date
    let estLaborHour = obj.ppk_inbox_from_ppe_est_labor_hour
    let worker = obj.ppk_inbox_from_ppe_worker
    let workerEmail = obj.ppk_inbox_from_ppe_worker_email
    let cost = obj.ppk_inbox_from_ppe_cost
    let priority = obj.WorkRequest.work_request_priority
    let location = obj.WorkRequest.work_request_location
    let equipment = obj.WorkRequest.work_request_equipment
    let instruction = obj.WorkRequest.work_request_instruction
    let description = obj.WorkRequest.work_request_description
    let newHTML = `<h3>Work Request ID</h3><h4>${wrid}</h4><h3>Pekerjaan</h3><h4>${task}</h4><h3>Tanggal Dibuat</h3><h4>${dateCreated}</h4><h3>Tanggal Mulai Pekerjaan</h3><h4>${estDate}</h4><h3>Banyak Jam Kerja</h3><h4>${estLaborHour}</h4><h3>Pekerja</h3><h4>${worker}</h4><h3>Email Pekerja</h3><h4>${workerEmail}</h4><h3>Biaya</h3><h4>${cost}</h4><h3>Prioritas</h3><h4>${priority}</h4><h3>Lokasi</h3><h4>${location}</h4><h3>Alat / Nama Alat</h3><h4>${equipment}</h4><h3>Instruksi</h3><p>${instruction}</p><h3>Description</h3><p>${description}</p>`
    detail.innerHTML = newHTML
    detailModal.style.display = "block"
}

function showWorkerDetail(id){
    let obj = workerMap.get(id)
    let dateCreated = obj.ppk_inbox_from_worker_date_created
    let woid = obj.WorkOrder.work_order_id
    let estDate = obj.WorkOrder.work_order_est_date
    let estLaborHour = obj.WorkOrder.work_order_est_labor_hour
    let worker = obj.WorkOrder.work_order_worker
    let workerEmail = obj.WorkOrder.work_order_worker_email
    let cost = obj.WorkOrder.work_order_cost
    let task = obj.WorkOrder.WorkRequest.work_request_task
    let wrid = obj.WorkOrder.WorkRequest.work_request_id
    let priority = obj.WorkOrder.WorkRequest.work_request_priority
    let location = obj.WorkOrder.WorkRequest.work_request_location
    let equipment = obj.WorkOrder.WorkRequest.work_request_equipment
    let instruction = obj.WorkOrder.WorkRequest.work_request_instruction
    let description = obj.WorkOrder.WorkRequest.work_request_description
    let newHTML = `<h3>Work Request ID</h3><h4>${wrid}</h4><h3>Work Order ID</h3><h4>${woid}</h4><h3>Pekerjaan</h3><h4>${task}</h4><h3>Tanggal Dibuat</h3><h4>${dateCreated}</h4><h3>Tanggal Mulai Pekerjaan</h3><h4>${estDate}</h4><h3>Banyak Jam Kerja</h3><h4>${estLaborHour}</h4><h3>Pekerja</h3><h4>${worker}</h4><h3>Email Pekerja</h3><h4>${workerEmail}</h4><h3>Biaya</h3><h4>${cost}</h4><h3>Prioritas</h3><h4>${priority}</h4><h3>Lokasi</h3><h4>${location}</h4><h3>Alat / Nama Alat</h3><h4>${equipment}</h4><h3>Instruksi</h3><p>${instruction}</p><h3>Description</h3><p>${description}</p>`
    detail.innerHTML = newHTML
    detailModal.style.display = "block"
}

function closeModalDetail(){
    detailModal.style.display = "none"
}

initWS()