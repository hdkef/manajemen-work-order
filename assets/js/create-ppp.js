const API_HOST = window.location.hostname + ":8080"

document.getElementById("btn-create-ppp").onclick = ()=>{
    CreatePPP()
}

let CreatePPP = ()=>{
    let inputNota = document.getElementById("input-nota")
    let inputPerihal = document.getElementById("input-perihal")
    let inputSifat = document.getElementById("input-sifat")
    let inputPekerjaan = document.getElementById("input-pekerjaan")
    let inputPhoto = document.getElementById("input-photo")
    
    let formData = new FormData()

    if (inputPhoto.files[0]){
        formData.append("photo",inputPhoto.files[0])
    }

    formData.append("nota",inputNota.value)
    formData.append("perihal",inputPerihal.value)
    formData.append("sifat",inputSifat.value)
    formData.append("pekerjaan",inputPekerjaan.value)
    
    sendXML(`http://${API_HOST}/api/v1/ppp`,formData,"post")
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
}

let handleError = (msg) => {
    if (msg){
        alert(`terjadi galat. ${msg}`)
        return
    }
    alert('terjadi galat.')
}