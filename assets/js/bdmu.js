const API_HOST = "localhost:8080"

let getBDMUPPP = ()=>{
    hitEndpoint(`http://${API_HOST}/api/v1/bdmu/ppp`,null,"get").then((resJSON)=>{
        populateTablePPP(resJSON)
    }).catch((err)=>{
        handleError(err)
    })
}

let getBDMURP = ()=>{
    hitEndpoint(`http://${API_HOST}/api/v1/bdmu/rp`,null,"get").then((resJSON)=>{
        populateTableRP(resJSON)
    }).catch((err)=>{
        handleError(err)
    })
}

let PPPDetail = (id)=>{
    window.open(`/ppp/${id}`)
}

let PPPAcc = (id,pppID)=>{
    hitEndpoint(`http://${API_HOST}/api/v1/ppp/${pppID}/ok/bdmu/${id}`,null,"post").then((resJSON)=>{
        handleSuccess(resJSON)
    }).catch((err)=>{
        handleError(err)
    })
}

let PPPDec = (id)=>{

}

let populateTablePPP = (slice)=>{
    let tablePPP = document.getElementById("table-ppp")
    for (let i=0;i<slice.length;i++){
        let id = +slice[i].id
        let pppId = +slice[i].ppp_id
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
        let btnTerima = document.createElement("button")
        let btnTolak = document.createElement("button")

        btnDetail.classList.add("btn","m-1")
        btnDetail.onclick = ()=>{
            PPPDetail(pppId)
        }
        btnDetail.textContent = "detail"

        btnTerima.classList.add("btn","btn-success","m-1")
        btnTerima.onclick = ()=>{
            PPPAcc(id,pppId)
        }
        btnTerima.textContent = "terima"

        btnTolak.classList.add("btn","btn-danger","m-1")
        btnTolak.onclick = ()=>{
            PPPDec(id)
        }
        btnTolak.textContent = "tolak"

        aksiCol.appendChild(btnDetail)
        aksiCol.appendChild(btnTerima)
        aksiCol.appendChild(btnTolak)

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

        tablePPP.appendChild(newRow)
    }
}

let RPLihat = (id)=>{
    window.open(`/rp/${id}`)
}

let RPAcc = (id,rpID)=>{
    hitEndpoint(`http://${API_HOST}/api/v1/rp/${rpID}/ok/bdmu/${id}`,null,"post").then((resJSON)=>{
        handleSuccess(resJSON)
    }).catch((err)=>{
        handleError(err)
    })
}

let RPDec = (id)=>{

}

let populateTableRP = (slice)=>{
    let tableRP = document.getElementById("table-rp")
    for (let i=0;i < slice.length;i++){
        let id = +slice[i].id
        let rpID = +slice[i].rp_id
        let dateCreate = slice[i].date_created
        let status = slice[i].rp.status

        let newRow = document.createElement("tr")
        let idCol = document.createElement("td")
        let rpIDCol = document.createElement("td")
        let dateCol = document.createElement("td")
        let statusCol = document.createElement("td")
        let aksiCol = document.createElement("td")
        let btnLihat = document.createElement("button")
        let btnTerima = document.createElement("button")
        let btnTolak = document.createElement("button")

        btnLihat.classList.add("btn")
        btnLihat.onclick = ()=>{
            RPLihat(rpID)
        }
        btnLihat.textContent = "lihat"

        btnTerima.classList.add("btn","btn-success")
        btnTerima.onclick = ()=>{
            RPAcc(id,rpID)
        }
        btnTerima.textContent = "terima"

        btnTolak.classList.add("btn","btn-danger")
        btnTolak.onclick = ()=>{
            RPDec(id)
        }
        btnTolak.textContent = "tolak"

        idCol.textContent = id
        rpIDCol.textContent = rpID
        dateCol.textContent = dateCreate
        statusCol.textContent = status

        aksiCol.appendChild(btnLihat)
        aksiCol.appendChild(btnTerima)
        aksiCol.appendChild(btnTolak)

        newRow.appendChild(idCol)
        newRow.appendChild(rpIDCol)
        newRow.appendChild(dateCol)
        newRow.appendChild(statusCol)
        newRow.appendChild(aksiCol)

        tableRP.appendChild(newRow)
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

getBDMUPPP()
getBDMURP()