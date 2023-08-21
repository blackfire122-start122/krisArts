function InBasket(e){
    let xhr = new XMLHttpRequest();
    let data = "artId=" + encodeURIComponent(e.id);

    xhr.open("POST", "/api/addToBasket", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");

    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                console.log("Art added to basket");
            } else {
                console.log("Error adding art to basket:", xhr.statusText);
            }
        }
    };

    xhr.send(data);
}