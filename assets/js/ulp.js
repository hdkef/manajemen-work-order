const API_HOST = "localhost:8080"

let getULPPerkiraanBiaya = ()=>{
    hitEndpoint(`http://${API_HOST}/api/v1/ulp/perkiraan-biaya`,null,"get").then((resJSON)=>{
        populateTablePerkiraanBiaya(resJSON)
    }).catch((err)=>{
        handleError(err)
    })
}

let perkiraanBiayaLihat = (id)=>{
    window.open(`/perkiraan-biaya/${id}`)
}

let createPengadaan = (id,perkiraanBiayaID)=>{
    window.open(`http://${API_HOST}/create-pengadaan/${id}/${perkiraanBiayaID}`)
}

let perkiraanBiayaDec = (id)=>{
}

let populateTablePerkiraanBiaya = (slice)=>{
    let tablePerkiraanBiaya = document.getElementById("table-perkiraan-biaya")
    for (let i=0;i < slice.length;i++){
        let id = +slice[i].id
        let perkiraanBiayaID = +slice[i].perkiraan_biaya_id
        let dateCreate = slice[i].date_created
        let estCost = slice[i].perkiraan_biaya.est_cost

        let newRow = document.createElement("tr")
        let idCol = document.createElement("td")
        let perkiraanBiayaIDCol = document.createElement("td")
        let dateCol = document.createElement("td")
        let estCostCol = document.createElement("td")
        let aksiCol = document.createElement("td")
        let btnLihat = document.createElement("button")
        let btnCreatePengadaan = document.createElement("button")

        btnLihat.classList.add("btn")
        btnLihat.onclick = ()=>{
            perkiraanBiayaLihat(perkiraanBiayaID)
        }
        btnLihat.textContent = "lihat"

        btnCreatePengadaan.classList.add("btn","btn-success")
        btnCreatePengadaan.onclick = ()=>{
            createPengadaan(id,perkiraanBiayaID)
        }
        btnCreatePengadaan.textContent = "buat pengadaan"

        idCol.textContent = id
        perkiraanBiayaIDCol.textContent = perkiraanBiayaID
        dateCol.textContent = dateCreate
        estCostCol.textContent = estCost

        aksiCol.appendChild(btnLihat)
        aksiCol.appendChild(btnCreatePengadaan)

        newRow.appendChild(idCol)
        newRow.appendChild(perkiraanBiayaIDCol)
        newRow.appendChild(dateCol)
        newRow.appendChild(estCostCol)
        newRow.appendChild(aksiCol)

        tablePerkiraanBiaya.appendChild(newRow)
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

getULPPerkiraanBiaya()