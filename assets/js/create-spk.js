const API_HOST = window.location.hostname + ":8080"

document.getElementById("btn-create-spk").onclick = ()=>{
    createSPK()
}

document.getElementById("btn-exit").onclick = ()=>{
    exit()
}

let createSPK = ()=>{
    let inputID = document.getElementById("input-id")
    let inputPengadaanID = document.getElementById("input-pengadaan-id")
    let inputDOC = document.getElementById("input-doc")
    let inputEmail = document.getElementById("input-worker-email")
    
    let id = +inputID.value
    let pengadaanID = +inputPengadaanID.value
    let email = inputEmail.value

    let formData = new FormData()
    formData.append("worker_email",email)

    if (!inputDOC.files[0]){
        alert("NO FILE")
        return
    }
    formData.append("doc",inputDOC.files[0])

    sendXML(`http://${API_HOST}/api/v1/spk/new/${pengadaanID}/${id}`,formData,"post")
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