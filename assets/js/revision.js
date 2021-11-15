const API_HOST = "localhost:8080"

document.getElementById("btn-revisi").onclick = ()=>{
    revision()
}

document.getElementById("btn-exit").onclick = ()=>{
    exit()
}

let revision = ()=>{
    let inputID = document.getElementById("input-id")
    let inputSPKID = document.getElementById("input-spk-id")
    let inputMsg = document.getElementById("input-msg")
    let id = +inputID.value
    let spkID = +inputSPKID.value
    let msg = inputMsg.value
    let payload = JSON.stringify({msg:msg})
    hitEndpoint(`http://${API_HOST}/api/v1/spk/${spkID}/no/${id}`,payload,"post").then((resJSON)=>{
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