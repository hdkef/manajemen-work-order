var ws
var token
var wrmodal = document.getElementById("wr-modal")
var detailmodal = document.getElementById("detail-modal")
var detail = document.getElementById("detail")
var wrMap = new Map()

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
           switch (jsonData.type){
               case "initUserFromServer":
                   let workRequests = jsonData.Data
                   //if data is exist then populate table
                   if (workRequests){
                       populateWR(workRequests)
                   }
                   break
               case "error":
                   alert(jsonData.msg)
                   window.location.href = "http://localhost:8080/login"
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

//to populate or create new row of work request history
function populateWR(wrArray){
    let tableBody = document.getElementById("table-body")
    for (let i =0;i < wrArray.length;i++){
        let id = wrArray[i].work_request_id
        let priority = wrArray[i].work_request_priority
        let date_created = wrArray[i].work_request_date_created
        let task = wrArray[i].work_request_task
        let status = wrArray[i].work_request_status
        let location = wrArray[i].work_request_location
        let equipment = wrArray[i].work_request_equipment
        let newRow = document.createElement("tr")
        let newRowInnerHTML = `<td>${id}</td><td>${priority}</td><td>${date_created}</td><td>${task}</td><td>${status}</td><td>${location}</td><td>${equipment}</td><td><button onclick="showDetail(${id})">Detail</button><button>Cancel</button></td>`
        newRow.innerHTML = newRowInnerHTML
        tableBody.appendChild(newRow)
        wrMap.set(id,wrArray[i])
    }
}

function showCreateWR(){
    wrmodal.style.display = "block"
}

function closeModalWR(){
    wrmodal.style.display = "none"
}

function showDetail(id){
    let wr = wrMap.get(id)
    let priority = wr.work_request_priority
    let date_created = wr.work_request_date_created
    let task = wr.work_request_task
    let status = wr.work_request_status
    let location = wr.work_request_location
    let equipment = wr.work_request_equipment
    let instruction = wr.work_request_instruction
    let description = wr.work_request_description
    var newHTML = `<h3>Work Request ID<h3><h4>${id}</h4><h3>Prioritas<h3><h4>${priority}</h4><h3>Tanggal Dibuat</h3><h3>${date_created}</h4><h3>Pekerjaan</h3><h4>${task}</h4><h3>Status</h3><h4>${status}</h4><h3>Lokasi</h3><h4>${location}</h4><h3>Nama / Tag Alat</h3><h4>${equipment}</h4><h3>Instruksi</h3><p>${instruction}</p><h3>Deskripsi</h3><p>${description}</p>`
    detail.innerHTML = newHTML
    detailmodal.style.display = "block"
}

function closeModalDetail(){
    detailmodal.style.display = "none"
}

initWS()