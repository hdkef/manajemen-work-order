const API_HOST = window.location.hostname + ":8080"

document.getElementById("btn-paging").onclick = ()=>{
    paging()
}

let LASTID = 0
let findLastID = ()=>{
    let tableBody = document.getElementById("table-perkiraan-biaya")
    let tableRow = tableBody.getElementsByTagName("tr")
    for (let i=0;i < tableRow.length;i++){
        let idCol = tableRow[i].firstElementChild
        let id = +idCol.textContent
        if (id > LASTID){
            LASTID = id
        }
    }
    return LASTID
}

let paging = ()=>{
    let lastid = +findLastID()
    hitEndpoint(`http://${API_HOST}/api/v1/perkiraan-biaya?last-id=${lastid}`,null,"get").then((resJSON)=>{
        populateTablePerkiraanBiaya(resJSON)
    }).catch((err)=>{
        handleError(err)
    })
}

let populateTablePerkiraanBiaya = (slice)=>{
    let tableRP = document.getElementById("table-perkiraan-biaya")
    for (let i=0;i<slice.length;i++){
        let id = +slice[i].id
        let dateCreated = slice[i].date_created
        let estCost = slice[i].est_cost
        let doc = slice[i].doc
        let creatorID = slice[i].creator_id

        let newRow = document.createElement("tr")
        let idCol = document.createElement("td")
        let dateCol = document.createElement("td")
        let estCostCol = document.createElement("td")
        let aksiCol = document.createElement("td")
        let anchorLihatDoc = document.createElement("a")
        let btnLihatDoc = document.createElement("button")
        let anchorLihatDetail = document.createElement("a")
        let btnLihatDetail = document.createElement("button")
        let anchorDetailPembuat = document.createElement("a")
        let btnDetailPembuat = document.createElement("button")

        anchorLihatDoc.href = `/${doc}`
        anchorLihatDoc.target = "_blank"
        anchorLihatDetail.href = `/perkiraan-biaya/${id}`
        anchorLihatDetail.target = "_blank"
        anchorDetailPembuat.href = `/entity/${creatorID}`
        anchorDetailPembuat.target = "_href"

        anchorLihatDoc.classList.add("m-2")
        anchorLihatDetail.classList.add("m-2")
        anchorDetailPembuat.classList.add("m-2")
        btnLihatDoc.classList.add("btn","btn-primary")
        btnLihatDetail.classList.add("btn","btn-primary")
        btnDetailPembuat.classList.add("btn","btn-primary")

        btnLihatDoc.textContent = "lihat dokumen"
        btnLihatDetail.textContent = "lihat detail"
        btnDetailPembuat.textContent = "detail pembuat"

        anchorLihatDoc.appendChild(btnLihatDoc)
        anchorLihatDetail.appendChild(btnLihatDetail)
        anchorDetailPembuat.appendChild(btnDetailPembuat)
        
        aksiCol.appendChild(anchorLihatDoc)
        aksiCol.appendChild(anchorLihatDetail)
        aksiCol.appendChild(anchorDetailPembuat)

        idCol.textContent = id
        dateCol.textContent = dateCreated
        estCostCol.textContent = estCost

        newRow.appendChild(idCol)
        newRow.appendChild(dateCol)
        newRow.appendChild(estCostCol)
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

let handleError = (msg) => {
    if (msg){
        alert(`terjadi galat. ${msg}`)
        return
    }
    alert('terjadi galat.')
}