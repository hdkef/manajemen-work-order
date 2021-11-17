const API_HOST = window.location.hostname + ":8080"

document.getElementById("btn-tolak").onclick = ()=>{
    tolak()
}

document.getElementById("btn-exit").onclick = ()=>{
    exit()
}

let tolak = ()=>{
    let inputrpID = document.getElementById("input-rp-id")
    let inputInboxID = document.getElementById("input-inbox-id")
    let inputMsg = document.getElementById("input-msg")
    let rpid = +inputrpID.value
    let inboxid = +inputInboxID.value
    let msg = inputMsg.value
    let payload = JSON.stringify({reason:msg})
    hitEndpoint(`http://${API_HOST}/api/v1/rp/${rpid}/no/${inboxid}`,payload,"post").then((resJSON)=>{
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