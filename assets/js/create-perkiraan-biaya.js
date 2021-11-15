const API_HOST = "localhost:8080"

document.getElementById("btn-create-perkiraan-biaya").onclick = ()=>{
    createPerkiraanBiaya()
}

document.getElementById("btn-exit").onclick = ()=>{
    exit()
}

let createPerkiraanBiaya = ()=>{
    let inputID = document.getElementById("input-id")
    let inputRPID = document.getElementById("input-rp-id")
    let inputToWhom = document.getElementById("input-towhom")
    let inputEstCost = document.getElementById("input-est-cost")
    let inputDOC = document.getElementById("input-doc")
    
    let id = +inputID.value
    let rpID = +inputRPID.value
    let toWhom = inputToWhom.value
    let estCost = inputEstCost.value

    let formData = new FormData()
    formData.append("est_cost",estCost)

    if (!inputDOC.files[0]){
        alert("NO FILE")
        return
    }
    formData.append("doc",inputDOC.files[0])
    
    switch(toWhom){
        case "ULP":
            sendXML(`http://${API_HOST}/api/v1/perkiraan-biaya/ulp/${rpID}/${id}`,formData,"post")
            return
        case "PPE":
            sendXML(`http://${API_HOST}/api/v1/perkiraan-biaya/ppe/${rpID}/${id}`,formData,"post")
            return
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