const API_HOST = window.location.hostname + ":8080"

document.getElementById("btn-login").onclick = ()=>{
    login()
}

let login = ()=>{

    let usernameInput = document.getElementById(("input-username"))
    let passwordInput = document.getElementById(("input-password"))
    let username = usernameInput.value
    let password = passwordInput.value
    let payload = JSON.stringify({username:username,password:password})
    hitEndpoint(`http://${API_HOST}/api/v1/login`,payload,"post").then((resJSON)=>{
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
    location.href = `http://${API_HOST}/login`
}

let handleError = (msg) => {
    if (msg){
        alert(`terjadi galat. ${msg}`)
        return
    }
    alert('terjadi galat.')
}