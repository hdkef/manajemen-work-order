const API_HOST = "localhost:8080"

document.getElementById("btn-submit-status").onclick = ()=>{
    UpdateStatus()
}

document.getElementById("btn-report-done").onclick = ()=>{
    laporSelesai()
}

let laporSelesai = ()=>{
    let inputID = document.getElementById("input-id")
    let inputPin = document.getElementById("input-pin")
    let id = +inputID.value
    let pin = +inputPin.value
    if (!pin){
        alert("PIN CANNOT EMPTY")
        return
    }
    let payload = JSON.stringify({pin:pin})
    hitEndpoint(`http://${API_HOST}/api/v1/spk/${id}/lapor`,payload,"post").then((resJSON)=>{
        handleSuccess(resJSON)
    }).catch((err)=>{
        handleError(err)
    })
}

let UpdateStatus = ()=>{
    let inputID = document.getElementById("input-id")
    let inputPin = document.getElementById("input-pin")
    let inputStatus = document.getElementById("input-status")
    let id = +inputID.value
    let status = inputStatus.value
    let pin = +inputPin.value
    if (!pin || !status){
        alert("CANNOT EMPTY")
        return
    }
    let payload = JSON.stringify({status:status,pin:pin})
    hitEndpoint(`http://${API_HOST}/api/v1/spk/${id}`,payload,"put").then((resJSON)=>{
        handleSuccess(resJSON)
    }).catch((err)=>{
        handleError(err)
    })
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
}

let handleError = (msg) => {
    if (msg){
        alert(`terjadi galat. ${msg}`)
        return
    }
    alert('terjadi galat.')
}