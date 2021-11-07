var ws
var token
var wrmodal = document.getElementById("wr-modal")
var detailmodal = document.getElementById("detail-modal")
var changepasswordmodal = document.getElementById("change-password-modal")
var detail = document.getElementById("detail")
var wrMap = new Map()
var lastID = 0
const APIHOST = "localhost:8080"

function initWS(){
   //TOBE
   fetch(`http://${APIHOST}/token`,null).then(res=>{return res.json()}).then(resJSON=>{
       token = resJSON.msg
       ws = new WebSocket(`ws://${APIHOST}/websocket/user`)

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
           let data = jsonData.data
           let msg = jsonData.msg
           switch (jsonData.type){
               case "changePasswordUserFromServer":
                   alert(msg)
                   break
               case "cancelWRUserFromServer":
                    alert(msg)
                    destroyWR(data)
                    break
               case "createWRUserFromServer":
                   alert(msg)
                   populateWR(data,"before")
                   break
               case "pagingUserFromServer":
                    if (data){
                        populateWR(data,"after")
                    }
                    break
               case "initUserFromServer":
                   //if data is exist then populate table
                   if (data){
                       populateWR(data,"after")
                   }
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

function destroyWR(id){
    wrMap.delete(+id)
    let wr = document.getElementById(`wr-${id}`)
    wr.parentNode.removeChild(wr)
}

//to populate or create new row of work request history
function populateWR(wrArray,appendType){
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
        newRow.id = `wr-${id}`
        let newRowInnerHTML = `<td>${id}</td><td>${priority}</td><td>${date_created}</td><td>${task}</td><td>${status}</td><td>${location}</td><td>${equipment}</td><td><button onclick="showDetail(${id})">Detail</button></td>`
        newRow.innerHTML = newRowInnerHTML
        wrMap.set(id,wrArray[i])
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

function showCreateWR(){
    wrmodal.style.display = "block"
}

function showChangePassword(){
    changepasswordmodal.style.display = "block"
}

function closeChangePassword(){
    changepasswordmodal.style.display = "none"
}

function closeModalWR(){
    wrmodal.style.display = "none"
}

function createWR(){
    let priority = document.getElementById("input-priority")
    let task = document.getElementById("input-task")
    let equipment = document.getElementById("input-equipment")
    let location = document.getElementById("input-location")
    let instruction = document.getElementById("input-instruction")
    let description = document.getElementById("input-description")
    //TOBE VALIDATION
    ws.send(JSON.stringify({
        type:"createWRUserFromClient",
        token:token,
        wrfromclient:{
            work_request_priority:priority,
            work_request_task:task,
            work_request_equipment:equipment,
            work_request_location:location,
            work_request_instruction:instruction,
            work_request_description:description,
        }
    }))
    closeModalWR()
}

function cancelWR(id){
    ws.send(JSON.stringify({
        type:"cancelWRUserFromClient",
        token:token,
        idfromclient:+id
    }))
    closeModalDetail()
}

function loadMore(){
    for (let [key,value] of wrMap.entries()){
        if (+key > lastID){
            lastID = +key
        }
    }
    ws.send(JSON.stringify({
        type:"pagingUserFromClient",
        token:token,
        last_id:lastID,
    }))
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
    let btnBatal = document.getElementById("button-cancel")
    var newHTML = `<h3>Work Request ID<h3><h4>${id}</h4><h3>Prioritas<h3><h4>${priority}</h4><h3>Tanggal Dibuat</h3><h3>${date_created}</h4><h3>Pekerjaan</h3><h4>${task}</h4><h3>Status</h3><h4>${status}</h4><h3>Lokasi</h3><h4>${location}</h4><h3>Nama / Tag Alat</h3><h4>${equipment}</h4><h3>Instruksi</h3><p>${instruction}</p><h3>Deskripsi</h3><p>${description}</p>`
    detail.innerHTML = newHTML
    detailmodal.style.display = "block"
    btnBatal.onclick = ()=>{
        cancelWR(id)
    }
}

function changePwd(){
    let oldPassword = document.getElementById("input-old-password").value
    let newPassword = document.getElementById("input-new-password").value
    //TOBE VALIDATION
    ws.send(JSON.stringify({
        type:"changePasswordUserFromClient",
        token:token,
        changepwdfromclient:{
            old:oldPassword,
            new:newPassword
        }
    }))
    closeChangePassword()
}

function closeModalDetail(){
    detailmodal.style.display = "none"
}

initWS()