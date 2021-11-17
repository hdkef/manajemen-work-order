const API_HOST = window.location.hostname + ":8080"

document.getElementById("btn-create-ppp").onclick = ()=>{
    CreatePPP()
}

let CreatePPP = ()=>{
    let inputNota = document.getElementById("input-nota")
    let inputPerihal = document.getElementById("input-perihal")
    let inputSifat = document.getElementById("input-sifat")
    let inputPekerjaan = document.getElementById("input-pekerjaan")
    
    let payload = JSON.stringify({
        nota:inputNota.value,
        pekerjaan:inputPekerjaan.value,
        sifat:inputSifat.value,
        perihal:inputPerihal.value
    })
    
    hitEndpoint(`http://${API_HOST}/api/v1/ppp`,payload,"post").then((resJSON)=>{
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