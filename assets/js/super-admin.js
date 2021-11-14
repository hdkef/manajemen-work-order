const API_HOST = "localhost:8080"

let getEntity = ()=>{
    hitEndpoint(`http://${API_HOST}/api/v1/entity`,null,"get").then((resJSON)=>{
        populateTableEntity(resJSON)
    }).catch((err)=>{
        handleError()
    })
}

let deleteEntity = (id)=>{
    hitEndpoint(`http://${API_HOST}/api/v1/entity/${id}`,null,"delete").then((resJSON)=>{
        handleSuccess(resJSON)
    }).catch((err)=>{
        handleError()
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

let populateTableEntity = (slice)=>{
    let tableEntity = document.getElementById("table-entity")
    for (let i=0;i<slice.length;i++){
        let id = +slice[i].id
        let fullname = slice[i].fullname
        let username = slice[i].username
        let role = slice[i].role
        let email = slice[i].email
        let signature = slice[i].signature
        
        let newRow = document.createElement("tr")
        let idCol = document.createElement("td")
        let fullnameCol = document.createElement("td")
        let usernameCol = document.createElement("td")
        let roleCol = document.createElement("td")
        let emailCol = document.createElement("td")
        let signatureCol = document.createElement("td")
        let actionCol = document.createElement("td")
        let btnDelete = document.createElement("button")

        btnDelete.textContent = "delete"
        btnDelete.classList.add("btn","btn-danger")
        btnDelete.onclick = ()=>{
            deleteEntity(id)
        }

        idCol.textContent = id
        fullnameCol.textContent = fullname
        usernameCol.textContent = username
        roleCol.textContent = role
        emailCol.textContent = email
        signatureCol.textContent = signature
        actionCol.appendChild(btnDelete)

        newRow.appendChild(idCol)
        newRow.appendChild(fullnameCol)
        newRow.appendChild(usernameCol)
        newRow.appendChild(roleCol)
        newRow.appendChild(emailCol)
        newRow.appendChild(signatureCol)
        newRow.appendChild(actionCol)

        tableEntity.appendChild(newRow)
    }
}

let handleSuccess = (resJSON) => {
    if (!resJSON.ok){
        handleError(resJSON.msg)
        return
    }
    alert(resJSON.msg)
    window.location.reload()
}

let handleError = (msg) => {
    if (msg){
        alert(`terjadi galat. ${msg}`)
        return
    }
    alert('terjadi galat.')
}

//execute
getEntity()