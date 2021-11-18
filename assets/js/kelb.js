const API_HOST = window.location.hostname + ":8080"

let getKELBPPP = ()=>{
    hitEndpoint(`http://${API_HOST}/api/v1/kelb/ppp`,null,"get").then((resJSON)=>{
        populateTablePPP(resJSON)
    }).catch((err)=>{
        handleError()
    })
}

let PPPDetail = (id)=>{
    window.open(`/ppp/${id}`)
}

let createRP = (id,pppID)=>{
    window.open(`/create-rp/${id}/${pppID}`)
}

let populateTablePPP = (slice)=>{
    if (!slice){
        return
    }
    let tablePPP = document.getElementById("table-ppp")
    for (let i=0;i<slice.length;i++){
        let id = +slice[i].id
        let pppid = +slice[i].ppp_id
        let dateCreated = slice[i].date_created
        let nota = slice[i].ppp.nota
        let perihal = slice[i].ppp.perihal
        let sifat = slice[i].ppp.sifat
        let pekerjaan = slice[i].ppp.pekerjaan
        let status = slice[i].ppp.status

        let newRow = document.createElement("tr")
        let idCol = document.createElement("td")
        let dateCol = document.createElement("td")
        let notaCol = document.createElement("td")
        let perihalCol = document.createElement("td")
        let sifatCol = document.createElement("td")
        let pekerjaanCol = document.createElement("td")
        let statusCol = document.createElement("td")
        let aksiCol = document.createElement("td")
        let btnDetail = document.createElement("button")
        let btnCreateRP = document.createElement("button")

        btnDetail.classList.add("btn","m-1")
        btnDetail.onclick = ()=>{
            PPPDetail(pppid)
        }
        btnDetail.textContent = "detail"

        btnCreateRP.classList.add("btn","btn-success","m-1")
        btnCreateRP.onclick = ()=>{
            createRP(id,pppid)
        }
        btnCreateRP.textContent = "buat rencana pekerjaan"

        aksiCol.appendChild(btnDetail)
        aksiCol.appendChild(btnCreateRP)

        idCol.textContent = id
        dateCol.textContent = dateCreated
        notaCol.textContent = nota
        perihalCol.textContent = perihal
        sifatCol.textContent = sifat
        pekerjaanCol.textContent = pekerjaan
        statusCol.textContent = status
        
        newRow.appendChild(idCol)
        newRow.appendChild(dateCol)
        newRow.appendChild(notaCol)
        newRow.appendChild(perihalCol)
        newRow.appendChild(sifatCol)
        newRow.appendChild(pekerjaanCol)
        newRow.appendChild(statusCol)
        newRow.appendChild(aksiCol)

        console.log(tablePPP)
        tablePPP.appendChild(newRow)
    }
}

let hitEndpoint = (endpoint,payload,method) => {
    return fetch(endpoint,{
        method:method,
        headers:{
            "Content-Type":"application/json"
        },
        body:payload
    }).then((res)=>{return res.json()})
}

let handleSuccess = (resJSON) => {
    if (!resJSON.ok){
        handleError(resJSON.msg)
        return
    }
    alert(resJSON.msg)
    window.location.reload()
}

let handleError = (msg) => {
    if (msg){
        alert(`terjadi galat. ${msg}`)
        return
    }
    alert('terjadi galat.')
}

getKELBPPP()