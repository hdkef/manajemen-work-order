const API_HOST = window.location.hostname + ":8080"

let getRiwayat = ()=>{
    alert("function is not implemented yet")
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