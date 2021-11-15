const API_HOST = "localhost:8080"

document.getElementById("btn-changepwd").onclick = ()=>{
    changePWD()
}

let changePWD = ()=>{
    let inputPassword = document.getElementById("input-password")
    let payload = JSON.stringify({password:inputPassword.value})
    hitEndpoint(`http://${API_HOST}/api/v1/changepwd`,payload,"put").then((resJSON)=>{
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