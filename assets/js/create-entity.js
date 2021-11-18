const API_HOST = window.location.hostname + ":8080"

document.getElementById("btn-create-entity").onclick = ()=>{
    createEntity()
}

let createEntity = ()=>{
    let inputFullname = document.getElementById("input-fullname")
    let inputUsername = document.getElementById("input-username")
    let inputPassword = document.getElementById("input-password")
    let inputEmail = document.getElementById("input-email")
    let inputRole = document.getElementById("input-role")
    let inputSignature = document.getElementById("input-signature")
    let fullname = inputFullname.value
    let username = inputUsername.value
    let password = inputPassword.value
    let email = inputEmail.value
    let role = inputRole.value

    let formData = new FormData()
    formData.append("fullname",fullname)
    formData.append("username",username)
    formData.append("password",password)
    formData.append("email",email)
    formData.append("role",role)
    if (!inputSignature.files[0]){
        alert("NO FILE")
        return
    }
    formData.append("signature",inputSignature.files[0])

    sendXML(`http://${API_HOST}/api/v1/entity`,formData,"post")
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