var ws
var token
var modal = document.getElementById("wr-modal")

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
        let priority = wrArray[i].work_request_priority
        let date_created = wrArray[i].work_request_date_created
        let task = wrArray[i].work_request_task
        let status = wrArray[i].work_request_status
        let location = wrArray[i].work_request_location
        let equipment = wrArray[i].work_request_equipment
        let newRow = document.createElement("tr")
        let newRowInnerHTML = `<td>${priority}</td><td>${date_created}</td><td>${task}</td><td>${status}</td><td>${location}</td><td>${equipment}</td><td><button>Detail</button><button>Cancel</button></td>`
        newRow.innerHTML = newRowInnerHTML
        tableBody.appendChild(newRow)
    }
}

function showCreateWR(){
    modal.style.display = "block"
}

function closeModalWR(){
    modal.style.display = "none"
}

initWS()