function Login(){
    let usernameInput  = document.getElementById("Username")
    let passwordInput = document.getElementById("Password")
    let username = usernameInput.value
    let password = passwordInput.value

    fetch("/login", {
        method:"POST",
        headers:{
            "Content-Type":"application/json"
        },
        body:JSON.stringify({user_username:username,user_password:password})
    }).then(res=>{
        return res.json()
    }).then(resJSON=>{
        if (resJSON.ok == false){
            alert(resJSON.msg)
        }else{
            //TOBE
            window.location.href = resJSON.msg
        }
    })
}