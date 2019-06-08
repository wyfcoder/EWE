
//检查用户名，不超过20个字符,唯一
function checkName() {
    let name=document.getElementById("name");
    if(name.value.toString().length>20){
        alert("The length of name is more than 20.Set name shorter.");
        return false;
    }
    if(name.value.length===0){
        alert("Input your name.");
        return false;
    }
    return true;
}

//检查账号，账号长度为6～11，且都为数字组成，唯一
function checkAccount() {
    let account=document.getElementById("account");
    let message=account.value.toString();
    if(message.length>11 || message.length<6){
        alert("Account needs 6～11 charsets from 0-9 .");
        return false;
    }else{
        for(i=0;i<11;i++){
            if(!(message[i]>='0'&&message[i]<='9')){
                alert("The Account's charset is from 0-9");
                return false;
            }
        }
    }
    return true;
}

//检查密码的设置，密码由数字和字母组成，最长为12个字符
function checkPassword() {
    let password=document.getElementById("password");

    if(password.value.length<6){
        alert("The Length of password is less than 6. Set your password longer.");
        return false;
    }
    else if(password.value.length>12){
        alert("The Length of password is more than 12. Set your password shorter.");
        return false;
    }
    else if(!isNumberOr_Letter(password.value.toString()))
{
    alert("Using Illegal character.The password is only build by numbers or letters.");
    return false;
}
    return true;
}

//检查找回密码的字符串，可由任意的字符构成，长度小于50即可
function checkFindPassword() {
    let findPassword=document.getElementById("findPassword");
    if(findPassword.value.length===0){
        alert("You need to edit the string.");
        return false;
    }

    if(findPassword.value.length>50){
        alert("The Length of string is more than 50. Set your string shorter.");
        return false;
    }
    return true;
}

//检查确定的字符是否相等
function checkConform() {
    let name=document.getElementById("name");
    let password=document.getElementById("password");
    let findPassword=document.getElementById("findPassword");
    let cPassword=document.getElementById("conformPassword");
    let cFindPassword=document.getElementById("conformFindPassword");

    if(name.value.toString().length>20){
        alert("The length of name is more than 20.Set name shorter.");
        return false;
    }
    if(name.value.length===0){
        alert("Input your name.");
        return false;
    }

    let account=document.getElementById("account");
    let message=account.value.toString();

    if(message.length>11 || message.length < 6){
        alert("Account needs 6~11 charsets from 0-9 .");
        return false;
    }else{
        for(i=0;i<message.length;i++){
            if(!(message[i]>='0'&&message[i]<='9')){
                alert("Account charset is from 0-9");
                return false;
            }
        }
    }

    if(password.value.length<6){
        alert("The Length of password is less than 6. Set your password longer.");
        return false;
    }
     if(password.value.length>12){
        alert("The Length of password is more than 12. Set your password shorter.");
        return false;
    }
     if(!isNumberOr_Letter(password.value.toString()))
    {
        alert("Using Illegal character.The password is only build by numbers or letters.");
        return false;
    }

    if(findPassword.value.length===0){
        alert("You need to edit the string.");
        return false;
    }

    if(findPassword.value.length>12){
        alert("The Length of string is more than 12. Set your string shorter.");
        return false;
    }

    if(!(password.value.toString()===cPassword.value.toString())){
        alert("Conform your password.");
        return false;
    }
    if(!(findPassword.value.toString()===cFindPassword.value.toString())){
        alert("Conform your string.");
        return false;
    }

    return true;
}

//检查密码是否相等
function checkConformPassword() {
    let password=document.getElementById("password");
    let cPassword=document.getElementById("conformPassword");
    if(password.value!=cPassword.value){
        alert("Conform your password.");
        return false;
    }
    return true;
}
//提交函数
function submit() {

    if(!checkName()) return false;

    if(!checkAccount()) return false;

    if(!checkPassword()) return false;

    if(!checkFindPassword()) return false;

    return true;
}

//判断是否是数字或字母组成，是则返回true,长度为6-12
function isNumberOr_Letter(s){
    var regu = "^[0-9a-zA-Z]{6,12}$";
    var re = new RegExp(regu);
    if (re.test(s)) {
        return true;
    }else{
        return false;
    }
}

function checkResetPassword() {
    if(checkPassword()&&checkConformPassword()){
        return true
    }
    return false;
}

function check_file() {
    let file = document.getElementById('file').files[0];
    if (file) {

        var fileSize = 0;

        if (file.size > 1024 * 1024) {
            fileSize = (Math.round(file.size * 100 / (1024 * 1024)) / 100);
            if (fileSize > 10) {
                alert("The file is exceed 10M !");
                return false;
            }
        }
    }
    else {
        alert("Please select file.");
        return false;
    }
    return true;
}