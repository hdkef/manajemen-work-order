const API_HOST = "localhost:8080"

let getPPKRP = ()=>{
    hitEndpoint(`http://${API_HOST}/api/v1/ppk/rp`,null,"get").then((resJSON)=>{
        populateTableRP(resJSON)
    }).catch((err)=>{
        handleError(err)
    })
}

let getPPKPengadaan = ()=>{
    hitEndpoint(`http://${API_HOST}/api/v1/ppk/pengadaan`,null,"get").then((resJSON)=>{
        populateTablePengadaan(resJSON)
    }).catch((err)=>{
        handleError(err)
    })
}

let getPPKSPK = ()=>{
    hitEndpoint(`http://${API_HOST}/api/v1/ppk/spk`,null,"get").then((resJSON)=>{
        populateTableWork(resJSON)
    }).catch((err)=>{
        handleError(err)
    })
}

let RPLihat = (id)=>{
    window.open(`/rp/${id}`)
}

let createPerkiraanBiaya = (id,rpID)=>{
    window.open(`http://${API_HOST}/create-perkiraan-biaya/${id}/${rpID}`)
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
        let btnCreatePerkiraanBiaya = document.createElement("button")
        let btnTolak = document.createElement("button")

        btnLihat.classList.add("btn")
        btnLihat.onclick = ()=>{
            RPLihat(rpID)
        }
        btnLihat.textContent = "lihat"

        btnCreatePerkiraanBiaya.classList.add("btn","btn-success")
        btnCreatePerkiraanBiaya.onclick = ()=>{
            createPerkiraanBiaya(id,rpID)
        }
        btnCreatePerkiraanBiaya.textContent = "buat perkiraan biaya"

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
        aksiCol.appendChild(btnCreatePerkiraanBiaya)
        aksiCol.appendChild(btnTolak)

        newRow.appendChild(idCol)
        newRow.appendChild(rpIDCol)
        newRow.appendChild(dateCol)
        newRow.appendChild(statusCol)
        newRow.appendChild(aksiCol)

        tableRP.appendChild(newRow)
    }
}

let PengadaanLihat = (id)=>{
    window.open(`/pengadaan/${id}`)
}

let createSPK = (id,pengadaanID)=>{
    window.open(`http://${API_HOST}/create-spk/${id}/${pengadaanID}`)
}

let populateTablePengadaan = (slice)=>{
    let tablePengadaan = document.getElementById("table-pengadaan")
    for (let i=0;i < slice.length;i++){
        let id = +slice[i].id
        let pengadaanID = +slice[i].pengadaan_id
        let dateCreate = slice[i].date_created

        let newRow = document.createElement("tr")
        let idCol = document.createElement("td")
        let pengadaanIDCol = document.createElement("td")
        let dateCol = document.createElement("td")
        let aksiCol = document.createElement("td")
        let btnLihat = document.createElement("button")
        let btnCreateSPK = document.createElement("button")

        btnLihat.classList.add("btn")
        btnLihat.onclick = ()=>{
            PengadaanLihat(pengadaanID)
        }
        btnLihat.textContent = "lihat"

        btnCreateSPK.classList.add("btn","btn-success")
        btnCreateSPK.onclick = ()=>{
            createSPK(id,pengadaanID)
        }
        btnCreateSPK.textContent = "buat surat perintah kerja"

        idCol.textContent = id
        pengadaanIDCol.textContent = pengadaanID
        dateCol.textContent = dateCreate

        aksiCol.appendChild(btnLihat)
        aksiCol.appendChild(btnCreateSPK)

        newRow.appendChild(idCol)
        newRow.appendChild(pengadaanIDCol)
        newRow.appendChild(dateCol)
        newRow.appendChild(aksiCol)

        tablePengadaan.appendChild(newRow)
    }
}

let SPKLihat = (id)=>{
    window.open(`/spk/${id}`)
}

let SPKOK = (spkID,id)=>{
    hitEndpoint(`http://${API_HOST}/api/v1/spk/${spkID}/ok/${id}`,null,"post").then((resJSON)=>{
        handleSuccess(resJSON)
    }).catch((err)=>{
        handleError(err)
    })
}

let SPKRevisi = (spkID,id)=>{
    window.open(`/revision/${id}/${spkID}`)
}

let populateTableWork = (slice)=>{
    let tableWorker = document.getElementById("table-worker")
    for (let i=0;i < slice.length;i++){
        console.log(slice)
        let id = +slice[i].id
        let spkID = +slice[i].spk_id
        let dateCreated = slice[i].date_created
        let status = slice[i].spk.status
        let workerEmail = slice[i].spk.worker_email

        let newRow = document.createElement("tr")
        let idCol = document.createElement("td")
        let spkIDCol = document.createElement("td")
        let dateCol = document.createElement("td")
        let statusCol = document.createElement("td")
        let workerEmailCol = document.createElement("td")
        let aksiCol = document.createElement("td")
        let btnLihat = document.createElement("button")
        let btnTerima = document.createElement("button")
        let btnRevisi = document.createElement("button")

        btnLihat.onclick = ()=>{
            SPKLihat(spkID)
        }
        btnLihat.textContent = "lihat"
        btnLihat.classList.add("btn")

        btnTerima.onclick = ()=>{
            SPKOK(spkID,id)
        }
        btnTerima.textContent = "terima"
        btnTerima.classList.add("btn","btn-success")

        btnRevisi.onclick = ()=>{
            SPKRevisi(spkID,id)
        }
        btnRevisi.textContent = "revisi"
        btnRevisi.classList.add("btn","btn-danger")

        aksiCol.append(btnLihat)
        aksiCol.append(btnTerima)
        aksiCol.append(btnRevisi)

        idCol.textContent = id
        spkIDCol.textContent = spkID
        dateCol.textContent = dateCreated
        statusCol.textContent = status
        workerEmailCol.textContent = workerEmail

        newRow.appendChild(idCol)
        newRow.appendChild(spkIDCol)
        newRow.appendChild(dateCol)
        newRow.appendChild(statusCol)
        newRow.appendChild(workerEmailCol)
        newRow.appendChild(aksiCol)

        tableWorker.appendChild(newRow)
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

getPPKRP()
getPPKPengadaan()
getPPKSPK()