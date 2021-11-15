const API_HOST = "localhost:8080"

document.getElementById("btn-create-rp").onclick = ()=>{
    createRP()
}

document.getElementById("btn-exit").onclick = ()=>{
    exit()
}

let createRP = ()=>{
    let inputID = document.getElementById("input-id")
    let inputPPPID = document.getElementById("input-ppp-id")
    let inputDOC = document.getElementById("input-rp")
    
    let id = +inputID.value
    let pppid = +inputPPPID.value

    let formData = new FormData()

    if (!inputDOC.files[0]){
        alert("NO FILE")
        return
    }
    formData.append("doc",inputDOC.files[0])

    sendXML(`http://${API_HOST}/api/v1/rp/new/${pppid}/${id}`,formData,"post")
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