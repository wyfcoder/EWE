function checkTitle() {
    let name=document.getElementById("name");

    if(name.value.length===0){
        alert("Input the title.");
        return false;
    }
    let content=document.getElementById("content");
    if(content.value.length===0){
        alert("Input the content.");
        return false;
    }
    return true;
}