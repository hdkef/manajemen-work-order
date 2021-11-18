const API_HOST = window.location.hostname + ":8080"

document.getElementById("btn-create-pengadaan").onclick = ()=>{
    createPengadaan()
}

document.getElementById("btn-exit").onclick = ()=>{
    exit()
}

let createPengadaan = ()=>{
    let inputID = document.getElementById("input-id")
    let inputPerkiraanBiayaID = document.getElementById("input-perkiraan-biaya-id")
    let inputDOC = document.getElementById("input-doc")
    let inputRole = document.getElementById("input-role")
    
    let id = +inputID.value
    let perkiraanBiayaID = +inputPerkiraanBiayaID.value
    let role = inputRole.value

    let formData = new FormData()

    if (!inputDOC.files[0]){
        alert("NO FILE")
        return
    }
    formData.append("doc",inputDOC.files[0])
    
    switch(role){
        case "PPE":
            sendXML(`http://${API_HOST}/api/v1/pengadaan/${perkiraanBiayaID}/ppe/${id}`,formData,"post")
            return
        case "ULP":
            sendXML(`http://${API_HOST}/api/v1/pengadaan/${perkiraanBiayaID}/ulp/${id}`,formData,"post")
            return
        default:
            alert("NOT ULP / PPE")
    }
}

let sendXML = (endpoint,payload,method) => {
    let request = new XMLHttpRequest();
    request.open(method, endpoint);
    request.send(payload);
    request.onreadystatechange = ()=>{
        if(request.readyState === XMLHttpRequest.DONE){
            let status = request.status
            let resJSON = JSON.parse(request.responseText)
            if (status == 200) {
                handleSuccess(resJSON)
            }else{
                handleError(resJSON.msg)
            }
        }
    }
}

let handleSuccess = (resJSON) => {
    if (!resJSON.ok){
        handleError(resJSON.msg)
        return
    }
    alert(resJSON.msg)
    exit()
}

let handleError = (msg) => {
    if (msg){
        alert(`terjadi galat. ${msg}`)
        return
    }
    alert('terjadi galat.')
}

let exit = ()=>{
    window.close()
}